package monorepo

import (
	"fmt"
	"os"
	"strings"

	billyfs "github.com/go-git/go-billy/v5/osfs"
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/cache"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/filesystem"
	"github.com/go-git/go-git/v5/storage/memory"
)

type Monorepo struct {
	parallelClone bool            // whether to clone repos in parallel or sequentially
	cloneCh       <-chan struct{} // channel for parallel cloning

	Dir                  string        // Target directory of the final monorepo
	SourceRepositoryURLs []*sourceRepo // URLs to source repositories
	UseInMemoryStorage   bool          // Whether to use in-memory storage or not, (not recommended for big repos)
	CloneDepth           int           // Used to configure deep or shallow clones, defaults to 0 which means all commits
}

type RepositoryCommit struct {
	*object.Commit
	RepositoryName string
}

type sourceRepo struct {
	url string
}

type TempDirs []string

func (tempDirs TempDirs) Clean() error {
	// Clean up the tmp directories
	if len(tempDirs) > 0 {
		for _, dirToPurge := range tempDirs {
			if err := os.RemoveAll(dirToPurge); err != nil {
				fmt.Printf("failed to remove tmp directory at %s got %v\n", dirToPurge, err)
				return err
			}
		}
	}
	return nil
}

func (s *sourceRepo) Name() string {
	repositoryName := s.url[strings.LastIndex(s.url, "/")+1:]
	repositoryName = strings.ReplaceAll(repositoryName, ".git", "")
	return repositoryName
}

func (s *sourceRepo) CloneToMemory(cloneOpts *git.CloneOptions) (*git.Repository, error) {
	storage := memory.NewStorage()
	repo, err := git.Clone(storage, nil, cloneOpts)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to git clone %s : %v", s.url, err))
		return nil, err
	}
	return repo, nil
}

func (s *sourceRepo) CloneToDisk(cloneOpts *git.CloneOptions) (string, *git.Repository, error) {
	tmpDirName, err := os.MkdirTemp(os.TempDir(), "monorepo_")
	if err != nil {
		return "", nil, err
	}
	storage := filesystem.NewStorage(billyfs.New(tmpDirName), cache.NewObjectLRUDefault())
	repo, err := git.Clone(storage, nil, cloneOpts)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to git clone %s to %s: %v", s.url, tmpDirName, err))
		return "", nil, err
	}
	return tmpDirName, repo, nil
}

func NewMonorepoFromSources(repos []string) *Monorepo {
	return newMonoRepo(repos, false)
}

func NewMonorepoFromSourcesInMemory(repos []string) *Monorepo {
	return newMonoRepo(repos, true)
}

func newMonoRepo(repos []string, inMemory bool) *Monorepo {
	sourceRepos := make([]*sourceRepo, 0)
	for _, repoUrlOrShorthand := range repos {
		sourceRepos = append(sourceRepos, &sourceRepo{
			url: makeFullUrl(repoUrlOrShorthand),
		})
	}
	return &Monorepo{
		SourceRepositoryURLs: sourceRepos,
		UseInMemoryStorage:   inMemory,
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
