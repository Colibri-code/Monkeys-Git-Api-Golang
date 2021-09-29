package api_git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func TreeCommitHead(repo *git.Repository) (repotree *object.Tree) {

	headrepo, _ := repo.Head()

	commit, _ := repo.CommitObject(headrepo.Hash())

	tree, _ := commit.Tree()

	return tree

}
