package examples

// change package to main and put it inside the main.go file if u need to see
// the running example

import (
	"fmt"
	"io"
	"log"
	"path"

	//	"strings"

	"github.com/emirpasic/gods/trees/binaryheap"
	"github.com/go-git/go-billy/v5"

	//"github.com/go-git/go-billy/v5/osfs"
	"github.com/go-git/go-git/plumbing"
	"github.com/go-git/go-git/v5"
	commitgraph_fmt "github.com/go-git/go-git/v5/plumbing/format/commitgraph"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/object/commitgraph"
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

type commitAndPaths struct {
	commit commitgraph.CommitNode
	// Paths that are still on the branch represented by commit
	paths []string
	// Set of hashes for the paths
	hashes map[string]plumbing.Hash
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

/*Funcion que retorna el contenido de cada Blob File
con la ruta dada por parametro*/
func ListContenBlobFile(repoPath string, fileP string) ([]string, error) {

	repo, err := git.PlainOpen(repoPath)

	var filepath []string

	var datafile []string

	ref, err := repo.Head()

	commit, err := repo.CommitObject(ref.Hash())

	tree, err := commit.Tree()

	/*For que llena el array para que
	Se pueda obtener el main tree de repositorio
	*/
	tree.Files().ForEach(func(f *object.File) error {

		filepath = append(filepath, f.Name)

		fmt.Printf("File Hashe and Path: %s    %s\n", f.Hash, f.Name)

		return nil
	})

	//Retorna el file de la ruta que se le pasa
	//Ejemplo
	/*
		repoPath = /var/www/git/repoexample.git
		fileP = css/assets

		retorna -> ERROR porque no es un archivo esta path es de una carpeta

		fileP = css/styles.css

		retorna ->  "body{\n    background-color: aqua;\n}\n\nheader{\n    position: relative;\n}"
	*/

	treefile, err := tree.File(fileP)

	if err != nil {
		return nil, err
	}

	for i := 0; i < len(filepath); i++ {

		if filepath[i] == fileP {

			content, err := treefile.Contents()

			datafile = append(datafile, string(content))

			if err != nil {

			}
			fmt.Printf("FilePath and parameter: %s    %s\n", filepath[i], fileP)
			fmt.Println(content)
			break
		}

	}

	fmt.Println(treefile)

	return datafile, err
}

func getCommitNodeIndex(r *git.Repository, fs billy.Filesystem) (commitgraph.CommitNodeIndex, io.ReadCloser) {

	file, err := fs.Open(path.Join("objects", "info", "commit-graph"))

	if err == nil {
		index, err := commitgraph_fmt.OpenFileIndex(file)
		if err == nil {
			return commitgraph.NewGraphCommitNodeIndex(index, r.Storer), file
		}
		file.Close()
	}
	return commitgraph.NewObjectCommitNodeIndex(r.Storer), nil
}

func getcommitTree(c commitgraph.CommitNode, treePath string) (*object.Tree, error) {

	tree, err := c.Tree()

	if err != nil {
		return nil, err
	}

	if treePath != "" {
		tree, err = tree.Tree(treePath)

		if err != nil {
			return nil, err
		}
	}
	return tree, nil
}
func getFileHashes(c commitgraph.CommitNode, treePath string, paths []string) (map[string]plumbing.Hash, error) {

	tree, err := getcommitTree(c, treePath)

	if err == object.ErrDirectoryNotFound {
		return make(map[string]plumbing.Hash), nil
	}
	if err != nil {
		return nil, err
	}

	hashes := make(map[string]plumbing.Hash)

	for _, path := range paths {
		if path != "" {
			entry, err := tree.FindEntry(path)

			if err == nil {
				hashes[path] = plumbing.Hash(entry.Hash)
			}
		} else {
			hashes[path] = plumbing.Hash(tree.Hash)
		}
	}

	return hashes, nil
}

// GetLastCommitForPaths returns last commit information

func getLastCommitForPaths(c commitgraph.CommitNode, treePath string, paths []string) (map[string]*object.Commit, error) {

	// We do a tree traversal with nodes sorted by commit time
	heap := binaryheap.NewWith(func(a, b interface{}) int {
		if a.(*commitAndPaths).commit.CommitTime().Before(b.(*commitAndPaths).commit.CommitTime()) {
			return 1
		}
		return -1
	})

	resultNode := make(map[string]commitgraph.CommitNode)

	initialHashes, err := getFileHashes(c, treePath, paths)

	if err != nil {
		return nil, err
	}
	// Start search from the root commit and with full set of paths
	heap.Push(&commitAndPaths{c, paths, initialHashes})

	for {
		cIn, ok := heap.Pop()

		if !ok {
			break
		}
		current := cIn.(*commitAndPaths)

		// Load the parent commits for the one we are currently examining

		numParents := current.commit.NumParents()

		var parents []commitgraph.CommitNode

		for i := 0; i < numParents; i++ {

			parent, err := current.commit.ParentNode(i)

			if err != nil {
				break
			}
			parents = append(parents, parent)
		}

		// Examine the current commit and set of interesting paths

		pathUnchanged := make([]bool, len(current.paths))

		parentHashes := make([]map[string]plumbing.Hash, len(parents))

		for j, parent := range parents {

			parentHashes[j], err = getFileHashes(parent, treePath, current.paths)

			if err != nil {
				break
			}

			for i, path := range current.paths {
				if parentHashes[j][path] == current.hashes[path] {
					pathUnchanged[i] = true
				}
			}
		}

		var remainingPaths []string

		for i, path := range current.paths {

			if resultNode[path] == nil {

				if pathUnchanged[i] {
					remainingPaths = append(remainingPaths, path)
				} else {
					resultNode[path] = current.commit
				}
			}
		}

		if len(remainingPaths) > 0 {

			for j, parent := range parents {

				remainingPathsForParent := make([]string, 0, len(remainingPaths))

				newRemainingPaths := make([]string, 0, len(remainingPaths))

				for _, path := range remainingPaths {

					if parentHashes[j][path] == current.hashes[path] {

						remainingPathsForParent = append(remainingPathsForParent, path)

					} else {
						newRemainingPaths = append(newRemainingPaths, path)
					}
				}
				if remainingPathsForParent != nil {
					heap.Push(&commitAndPaths{parent, remainingPathsForParent, parentHashes[j]})
				}

				if len(newRemainingPaths) == 0 {
					break
				} else {
					remainingPaths = newRemainingPaths
				}
			}
		}
	}

	result := make(map[string]*object.Commit)

	for path, commitNode := range resultNode {

		var err error

		result[path], err = commitNode.Commit()

		if err != nil {
			return nil, err
		}
	}

	return result, nil
}
