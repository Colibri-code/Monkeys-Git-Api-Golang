package api_git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func getCommit(hash plumbing.Hash, repo git.Repository) (*object.Commit, error) {
	commit, err := repo.CommitObject(hash)
	if err != nil {
		return nil, err
	}

	return commit, nil
}

func difference(path string, hash string) (*object.Patch, error) {
	repo, err := git.PlainOpen(path)
	var patch *object.Patch = nil

	if repo != nil {
		hashObject := plumbing.NewHash(hash)
		commit, err := getCommit(hashObject, *repo)
		if err == nil && commit != nil {
			ref, _ := repo.Head()
			commitHead, _ := repo.CommitObject(ref.Hash())
			patch, err = commit.Patch(commitHead)
			if err == nil && patch != nil {
				return patch, nil
			}
		}
	}

	if err != nil {
		return nil, err
	}
	return patch, err

}
