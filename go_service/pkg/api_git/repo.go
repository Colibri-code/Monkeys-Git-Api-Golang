package api_git

import (
	"errors"
	"os/exec"

	"github.com/go-git/go-git/v5"
)

//Create repo from bash
func CreateRepoFromBash(name string) (bool, error) {
	out, err := exec.Command("/bin/sh", "../../../etc/git-create-repo.sh", name).Output()
	if err != nil {
		return false, err
	}
	if out != nil {
		return true, nil
	}
	return false, err
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

func CopyRepoFromTag(tagRequest TagRequest) (bool, error) {
	isCreate, err := CreateRepoFromBash(tagRequest.Name)
	if err != nil {
		return false, err
	}
	if isCreate {
		path := "/var/www/git/" + tagRequest.Name + ".git"
		repo, err := CloneRepository(path, tagRequest.Url)
		if err != nil {
			if err.Error() == "repository already exists" {
				repo, err = git.PlainOpen(path)
				if err != nil {
					return false, err
				}
			} else {
				return false, err
			}

		}
		if repo != nil {
			tagReference, err := TagExists(tagRequest.Tag, *repo)
			if err != nil {
				return false, err
			}
			if tagReference.Name().String() != "" {
				result, err := createBranchFromTag(&tagReference, repo)
				if err != nil {
					return false, err
				}
				if result {
					return true, nil
				}
			} else {
				return false, errors.New("Tag does not exist")
			}
		}
	} else {
		return false, err
	}
	return false, err
}

//Function to open repository
func OpenRepository(RepoPath string) (*git.Repository, error) {

	if RepoPath != "" {
		ConcatRepoPath := baseRepoDir + RepoPath + ".git"

		repository, err := git.PlainOpen(ConcatRepoPath)

		if err != nil {
			return nil, git.ErrRepositoryNotExists
		}
		return repository, err
	}
	return nil, git.ErrRepositoryNotExists
}
