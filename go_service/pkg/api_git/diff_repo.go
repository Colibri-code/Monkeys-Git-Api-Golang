package api_git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func getCommit(hash string, repo git.Repository) (*object.Commit, error) {
	hashObject := plumbing.NewHash(hash)
	commit, err := repo.CommitObject(hashObject)
	if err != nil {
		return nil, err
	}

	return commit, nil
}

func diffToHead(path string, hash string) (*object.Patch, error) {
	repo, err := git.PlainOpen(path)
	var patch *object.Patch = nil

	if repo != nil {
		commit, err := getCommit(hash, *repo)
		if err == nil && commit != nil {

			ref, _ := repo.Head()
			commitHead, _ := repo.CommitObject(ref.Hash())
			//Compare with the head of the repo
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
