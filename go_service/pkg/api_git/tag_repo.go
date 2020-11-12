package api_git

import (
	"errors"
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

//Check if the Tag exist and return it
func TagExists(tag string, r git.Repository) (plumbing.Reference, error) {
	tagFoundErr := "tag was found"
	tagadd := "refs/tags/" + tag
	var tagReference plumbing.Reference
	tags, err := r.Tags()
	if err != nil {
		fmt.Errorf("get tags error: %s", err)
		return tagReference, err
	}

	err = tags.ForEach(func(t *plumbing.Reference) error {
		if t.Name().String() == tagadd {
			tagReference = *t
			return nil
		}
		return errors.New(tagFoundErr)
	})

	if err != nil && err.Error() != tagFoundErr {
		fmt.Errorf("iterate tags error: %s", err)
		return tagReference, err
	}
	return tagReference, nil
}

// Create a Branch from specific tag
func createBranchFromTag(tag *plumbing.Reference, repository *git.Repository) (bool, error) {
	newBranch := plumbing.ReferenceName("refs/heads/" + tag.Name().String() + "-branch")
	worktree, err := repository.Worktree()
	if worktree != nil {
		opts := git.CheckoutOptions{
			Hash:   tag.Hash(),
			Create: true,
			Branch: newBranch,
			Force:  false,
		}
		err = worktree.Checkout(&opts)

		if err != nil {
			println(err.Error())
			return false, err
		}
	}
	if err != nil {
		return false, nil
	}
	return true, nil
}

//Clone a repository in a specific directory
func CloneRepository(dir string, url string) (*git.Repository, error) {
	repository, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL: url,
	})

	if err != nil {
		return nil, err
	}

	return repository, nil
}
