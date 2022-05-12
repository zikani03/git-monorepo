package monorepo

import (
	"path"
	"sort"
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
	SourceRepositoryURLs []string
}

func NewMonorepoFromSources(repos []string) *Monorepo {
	return &Monorepo{
		SourceRepositoryURLs: repos,
	}
}

func (m *Monorepo) Init(targetDir string) error {
	// ... retrieves the commit history
	since := time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC)
	until := time.Now()
	return m.initWithCommitsBetween(targetDir, since, until)
}

func (m *Monorepo) initWithCommitsBetween(targetDir string, fromTime, toTime time.Time) error {
	// commitSL := skiplist.New()
	// This commit list could potentially be huuuge we need a good way to store it.
	commitList := make([]*object.Commit, 0)

	for _, sourceRepositoryUrl := range m.SourceRepositoryURLs {
		repo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
			URL: sourceRepositoryUrl,
		})
		if err != nil {
			return err
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

		err = cIter.ForEach(func(c *object.Commit) error {
			//commitSL.Insert(skiplist.Element(c))
			commitList = append(commitList, c)
			return nil
		})

		if err != nil {
			return err
		}
	}

	sort.Slice(commitList, func(i, j int) bool {
		return commitList[i].Committer.When.Before(commitList[j].Committer.When)
	})

	worktreeDir := billyfs.New(targetDir)
	dotGitDir := billyfs.New(path.Join(targetDir, ".git"))
	monorepo, err := git.Init(filesystem.NewStorage(dotGitDir, cache.NewObjectLRUDefault()), worktreeDir)

	if err != nil {
		return err
	}

	w, err := monorepo.Worktree()

	if err != nil {
		return err
	}

	for _, c := range commitList {
		files, err := c.Files()
		if err != nil {
			return err
		}

		files.ForEach(func(f *object.File) error {
			w.Add(f.Name)
			return nil
		})

		_, err = w.Commit(c.Message, &git.CommitOptions{
			Author:    &c.Author,
			Committer: &c.Committer,
		})
		if err != nil {
			return err
		}
		// fmt.Println(c)
	}

	return nil
}
