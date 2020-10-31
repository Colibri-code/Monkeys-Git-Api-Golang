package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
)

// - Clone a repository into memory
// - Get the HEAD reference
// - Using the HEAD reference, obtain the commit this reference is pointing to
// - Using the commit, obtain its history and print it
func main() {

	url := "https://github.com/go-git/go-git"
	dir := "/tmp/sample-controller-workshop"
	tag := "v0.1.0"

	r, err := cloneRepo(url, dir)

	if err != nil {
		log.Printf("clone repo error: %s", err)
		return
	}

	if tagExists(tag, r) {
		log.Printf("Tag %s already exists, nothing to do here", tag)
		return
	}

}

func cloneRepo(url, dir string) (*git.Repository, error) {
	// Clones the given repository, creating the remote, the local branches
	// and fetching the objects, everything in memory:
	Info(url)
	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: url,
	})
	// Gets the HEAD history from HEAD, just like this command:
	Info("git log")

	// ... retrieves the branch pointed by HEAD
	ref, err := r.Head()

	// ... retrieves the commit history
	since := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	until := time.Date(2020, 7, 30, 0, 0, 0, 0, time.UTC)
	cIter, err := r.Log(&git.LogOptions{From: ref.Hash(), Since: &since, Until: &until})
	CheckIfError(err)

	// ... Print head of the repo
	// ... c Variable contain the object
	err = cIter.ForEach(func(c *object.Commit) error {
		fmt.Println(c)

		return nil
	})

	return r, nil
}

func tagExists(tag string, r *git.Repository) bool {

	res := true

	return res
}

// CheckIfError should be used to naively panics if an error is not nil.
func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

// Info should be used to describe the example commands that are about to run.
func Info(format string, args ...interface{}) {
	fmt.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}
