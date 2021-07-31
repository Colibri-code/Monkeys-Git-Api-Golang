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

func ListFile(Path string, Namefile string) (string, error) {
	repo, err := git.PlainOpen(Path)
	var file string
	if repo != nil {
		treeIter, errW := repo.TreeObjects()

		if treeIter != nil {
			treeIter.ForEach(func(t *object.Tree) error {

				// ... get the files iterator and print the file

				t.Files().ForEach(func(f *object.File) error {

					if f.Name == Namefile {

						fmt.Printf(f.Name, "entra aqui")

						file = f.Name
					} else {
						log.Printf("There is not exist a file ")
					}
					return nil
				})
				return nil
			})
		}
		if errW != nil {
			log.Printf("There is not a tree %s", errW)
		}
	}
	return file, err
}

func ListFilesDirectories(Path string) ([]string, error) {
	repo, err := git.PlainOpen(Path)
	var files []string
	if repo != nil {
		treeIter, errW := repo.TreeObjects()

		if treeIter != nil {
			treeIter.ForEach(func(t *object.Tree) error {

				// ... get the files iterator and print the file

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

func ListBlobFile(Path string) ([]string, error) {

	repo, err := git.PlainOpen(Path)
	var fileblob []string

	ref, err := repo.Head()

	commit, err := repo.CommitObject(ref.Hash())

	fmt.Println(commit)

	tree, err := commit.Tree()

	tree.Files().ForEach(func(f *object.File) error {

		fileblob = append(fileblob, f.Name)

		fmt.Printf("100644 blob %s    %s\n", f.Hash, f.Name)
		return nil
	})
	/*	if repo != nil {
		blobIter, errB := repo.BlobObjects()

		if blobIter != nil {

			blobIter.ForEach(func(b *object.Blob) error {
				fileblob = append(fileblob, b.ID().String())
				return nil
			})
		}
		if errB != nil {
			log.Printf("There is not a tree %s", errB)
		}

	}*/

	return fileblob, err
}
