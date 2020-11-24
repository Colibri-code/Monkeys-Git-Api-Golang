package api_git

import (
	"fmt"
	"os/exec"
)

//Variable to point to Base Repository
var baseRepoDir = "/var/www/git/"

//CreateMerge Method to create merge between to branches
func CreateMerge(mergeRequest MergeRequest) (string, error) {
	cmdCkech := exec.Command("git", "checkout "+mergeRequest.TargetBranch)
	cmdCkech.Dir = baseRepoDir + mergeRequest.Url
	resultCmd, err := cmdCkech.Output()
	var result string = ""
	if resultCmd != nil {
		fmt.Printf("%s", resultCmd)
	}
	if err == nil {
		cmd := exec.Command("git", "merge ", mergeRequest.Branch)
		cmd.Dir = baseRepoDir + mergeRequest.Url
		resultCmd, err := cmd.Output()

		if err != nil {
			fmt.Printf("%s", err.Error())
			return "", err
		}
		fmt.Printf("%s", resultCmd)
	} else {
		fmt.Printf("%s", err.Error())
	}

	return result, nil
}
