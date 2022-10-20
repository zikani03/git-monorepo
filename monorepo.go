package monorepo

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"sort"
	"strings"
	"time"

	billyfs "github.com/go-git/go-billy/v5/osfs"
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/cache"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/filesystem"
	"github.com/go-git/go-git/v5/storage/memory"
)

type Monorepo struct {
	// done chan bool
	Dir                  string
	SourceRepositoryURLs []string
	UseInMemoryStorage   bool
}

type RepositoryCommit struct {
	*object.Commit
	RepositoryName string
}

func NewMonorepoFromSources(repos []string) *Monorepo {
	return &Monorepo{
		SourceRepositoryURLs: repos,
		UseInMemoryStorage:   false,
	}
}

func NewMonorepoFromSourcesInMemory(repos []string) *Monorepo {
	return &Monorepo{
		SourceRepositoryURLs: repos,
		UseInMemoryStorage:   true,
	}
}

func makeFullUrl(spec string) string {
	if strings.HasPrefix(spec, "gh:") {
		return strings.ReplaceAll(spec, "gh:", "https://github.com/")
	}
	if strings.HasPrefix(spec, "bb:") {
		return strings.ReplaceAll(spec, "bb:", "https://bitbucket.org/")
	}
	if strings.HasPrefix(spec, "gl:") {
		return strings.ReplaceAll(spec, "gl:", "https://gitlab.com/")
	}

	return spec
}

func (m *Monorepo) Init(targetDir string) error {
	// ... retrieves the commit history
	m.Dir = targetDir
	since := time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC)
	until := time.Now()
	return m.initWithCommitsBetween(targetDir, since, until)
}

func (m *Monorepo) initWithCommitsBetween(targetDir string, fromTime, toTime time.Time) error {
	// commitSL := skiplist.New()
	// This commit list could potentially be huuuge we need a good way to store it.
	commitList := make([]*RepositoryCommit, 0)

	worktreeDir := billyfs.New(targetDir)

	dirsToClean := make([]string, 0)
	for _, urlSpec := range m.SourceRepositoryURLs {
		sourceRepositoryUrl := makeFullUrl(urlSpec)
		var repo *git.Repository
		var err error
		if m.UseInMemoryStorage {
			storage := memory.NewStorage()
			repo, err = git.Clone(storage, nil, &git.CloneOptions{
				URL: sourceRepositoryUrl,
			})
			if err != nil {
				fmt.Println(fmt.Errorf("failed to git clone %s : %v", sourceRepositoryUrl, err))
				return err
			}
		} else {
			tmpDirName, err := os.MkdirTemp(os.TempDir(), "monorepo_")
			if err != nil {
				return err
			}
			storage := filesystem.NewStorage(billyfs.New(tmpDirName), cache.NewObjectLRUDefault())
			repo, err = git.Clone(storage, nil, &git.CloneOptions{URL: sourceRepositoryUrl})
			if err != nil {
				fmt.Println(fmt.Errorf("failed to git clone %s to %s: %v", sourceRepositoryUrl, tmpDirName, err))
				return err
			}
			dirsToClean = append(dirsToClean, tmpDirName)
		}

		// retrieves the branch pointed by HEAD
		repoRef, err := repo.Head()
		if err != nil {
			return err
		}
		cIter, err := repo.Log(&git.LogOptions{From: repoRef.Hash(), Since: &fromTime, Until: &toTime})
		if err != nil {
			return err
		}

		repositoryName := sourceRepositoryUrl[strings.LastIndex(sourceRepositoryUrl, "/"):]
		repositoryName = strings.ReplaceAll(repositoryName, ".git", "")

		err = cIter.ForEach(func(c *object.Commit) error {
			repoCommit := &RepositoryCommit{RepositoryName: repositoryName, Commit: c}
			//fmt.Println("Adding commit", repoCommit)
			commitList = append(commitList, repoCommit)
			return nil
		})

		if err != nil {
			return fmt.Errorf("error during iterating commits for repo:%s : %v", sourceRepositoryUrl, err)
		}
	}

	sort.Slice(commitList, func(i, j int) bool {
		return commitList[i].Committer.When.Before(commitList[j].Committer.When)
	})

	dotGitDir := billyfs.New(path.Join(targetDir, ".git"))
	monorepo, err := git.Init(filesystem.NewStorage(dotGitDir, cache.NewObjectLRUDefault()), worktreeDir)

	if err != nil {
		return fmt.Errorf("failed to git init new monorepo: %v", err)
	}

	w, err := monorepo.Worktree()
	if err != nil {
		return fmt.Errorf("failed to get working tree of new monorepo: %v", err)
	}

	for _, c := range commitList {
		files, err := c.Files()
		if err != nil {
			return err
		}
		if err = worktreeDir.MkdirAll(c.RepositoryName, fs.ModeDir); err != nil {
			return err
		}
		err = files.ForEach(func(f *object.File) error {
			// absFilePath := path.Join(m.Dir, c.RepositoryName, f.Name)
			localFilePath := path.Join(c.RepositoryName, f.Name)
			workingFile, err := worktreeDir.Create(localFilePath)
			if err != nil {
				return fmt.Errorf("failed to create file in worktree %s got: %s", localFilePath, err)
			}
			defer workingFile.Close()
			reader, err := f.Reader()
			if err != nil {
				return fmt.Errorf("failed to get io.ReadCloser for file %s got: %s", localFilePath, err)
			}
			defer reader.Close()
			b := make([]byte, f.Size)

			if _, err = reader.Read(b); err != nil {
				if err != io.EOF { // empty files get EOF
					return fmt.Errorf("failed to read file for writing %s got: %s", localFilePath, err)
				}
			} else {
				if err != io.EOF { // empty files get EOF
					if _, err = workingFile.Write(b); err != nil {
						return fmt.Errorf("failed to write file %s got: %s", localFilePath, err)
					}
				}
			}

			err = w.AddWithOptions(&git.AddOptions{
				All:  true,
				Path: c.RepositoryName,
			})
			return err
		})
		if err != nil {
			return fmt.Errorf("failed to write files: %v", err)
		}
		_, err = w.Commit(c.Message, &git.CommitOptions{
			Author:    &c.Author,
			Committer: &c.Committer,
			All:       true,
		})
		if err != nil {
			return fmt.Errorf("failed to commit files: %v", err)
		}
	}

	// Clean up the tmp directories
	if len(dirsToClean) > 0 {
		for _, dirToPurge := range dirsToClean {
			if err := os.RemoveAll(dirToPurge); err != nil {
				fmt.Printf("failed to remove tmp directory at %s got %v\n", dirToPurge, err)
			}
		}
	}

	return nil
}
