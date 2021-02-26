package examples

// change package to main and put it inside the main.go file if u need to see
// the running example

import (
	"fmt"
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func main() {

	repo, _ := git.PlainOpen("../monkeytest.git")

	ref, _ := repo.Head()

	commit, _ := repo.CommitObject(ref.Hash())

	var author = commit.Author

	var committer = commit.Committer

	var hash = commit.Hash

	var message = commit.Message

	var time = commit.Author.When

	fmt.Println("Author of the commit:", author)
	fmt.Println("Commiter of this commit: ", committer)
	fmt.Println("Message of this commit: ", message)
	fmt.Println("hash of this commit: ", hash)
	fmt.Println("Time stamp of the commit: ", time)
}

func ListFilesDirectories(Path string) ([]string, error) {
	repo, err := git.PlainOpen(Path)
	var files []string
	if repo != nil {
		treeIter, errW := repo.TreeObjects()

		if treeIter != nil {
			treeIter.ForEach(func(t *object.Tree) error {
				t.Files().ForEach(func(f *object.File) error {
					files = append(files, f.Name)
					return nil
				})
				return nil
			})
		}
		if errW != nil {
			log.Printf("There is not a tree %s", errW)
		}
	}
	return files, err
}
