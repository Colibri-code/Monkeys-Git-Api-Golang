package api_git

import (
	"bytes"
	"fmt"
	"go_service/tools"
	"os/exec"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

//Variable to point to Base Repository
const baseRepoDir = "/var/www/git/"
const BranchPrefix = "refs/heads/"

//CreateMerge Method to create merge between to branches
func CreateMerge(mergeRequest MergeRequest) (string, error) {
	repo, err := git.PlainOpen(baseRepoDir + mergeRequest.Url)
	var result string = ""
	if repo != nil {
		reference, _ := getBranch(*repo, mergeRequest.TargetBranch)
		if reference != nil {
			isReference, err := checkoutBranch(*reference, *repo)

			if isReference {

				cmd := exec.Command("git", "merge", mergeRequest.Branch)
				cmd.Dir = baseRepoDir + mergeRequest.Url

				var stderr bytes.Buffer

				cmd.Stderr = &stderr
				resultCmd, err := cmd.Output()

				if err != nil {
					fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
					return "", err
				}
				fmt.Printf("%s", resultCmd)
				result = tools.BytesToString(resultCmd)
			} else if err != nil {
				return result, err
			}
		}

		return result, nil
	} else {
		return "", err
	}

}

//Checkout a branch
func checkoutBranch(reference plumbing.Reference, repo git.Repository) (bool, error) {
	w, err := repo.Worktree()
	err = w.Checkout(&git.CheckoutOptions{
		Branch: reference.Name(),
		Force:  true,
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

//Get specific branch reference
func getBranch(repo git.Repository, branchName string) (*plumbing.Reference, error) {

	reference, err := repo.Reference(plumbing.ReferenceName(BranchPrefix+branchName), true)
	if err != nil {
		return nil, err
	} else {
		return reference, nil
	}
}
