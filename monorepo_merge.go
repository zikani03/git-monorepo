package monorepo

import (
	"fmt"

	git "github.com/go-git/go-git/v5"
)

// MergeSourcesInto merges source repositories into an existing target Git repository
// each source repository will be added as a directory to the repo.
// History is not preserved, by default all commits are made as squash commit and added
// to the target repo in one commit.
//
// TODO: handle situation where an existing directory in the repo matches source repo name
// TODO: add option to preserve or drop organization name from repo e.g. zikani/repo maintains "zikani" as part of the name
// TODO:
func (m *Monorepo) MergeSourcesInto(targetDir string) error {
	if targetDir == "" {
		return fmt.Errorf("targetRepo cannot be empty or nil")
	}
	m.Dir = targetDir
	dirsToClean := make(TempDirs, 0)

	for _, aRepo := range m.SourceRepositoryURLs {
		var repo *git.Repository
		var tmpDirName string
		var err error

		cloneOpts := &git.CloneOptions{
			URL: aRepo.url,
		}

		if m.CloneDepth > 0 {
			cloneOpts.Depth = m.CloneDepth
		}

		if m.UseInMemoryStorage {
			repo, err = aRepo.CloneToMemory(cloneOpts)
			if err != nil {
				return err
			}
		} else {
			tmpDirName, repo, err = aRepo.CloneToDisk(cloneOpts)
			if err != nil {
				return err
			}
			dirsToClean = append(dirsToClean, tmpDirName)
		}

		// retrieves the branch pointed by HEAD
		repoRef, err := repo.Head()
		if err != nil {
			return err
		}
		// since := time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC)
		// until := time.Now()
		// cIter, err := repo.Log(&git.LogOptions{From: repoRef.Hash(), Since: &since, Until: &until})
		// if err != nil {
		// 	return err
		// }
		fmt.Println("Repository head ref", repoRef.String())
		// repositoryName := aRepo.Name()
	}

	dirsToClean.Clean()
	// TODO: merge the sources into the target repo here, preserve their history?
	return fmt.Errorf("failed to merge sources into target repo")
}
