package api_git

// change package to main and put it inside the main.go file if u need to see
// the running example

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

/*Toma todos los files que tiene el tree Head del repositorio y los
retorna en un []string*/
func ListPathFileRepository(repoPath string) ([]string, error) {

	if repoPath != "" {

		repo, err := git.PlainOpen(repoPath)

		var filepath []string

		ref, err := repo.Head()

		commit, err := repo.CommitObject(ref.Hash())

		tree, err := commit.Tree()

		for _, entry := range tree.Entries {

			filepath = append(filepath, entry.Name)
		}

		return filepath, err
	}

	return nil, git.ErrRepositoryNotExists
}

func ContentTreeData(repoPath string, filepath string) ([]string, error) {

	repo, err := git.PlainOpen(repoPath)

	ref, err := repo.Head()

	var EntryPaths []string

	var datafile []string

	var TreeEntries []string

	commit, err := repo.CommitObject(ref.Hash())

	tree, err := commit.Tree()

	/*Comprobacion de que la path no venga vacia, si viene vacia se lista el
	main tree del repositorio*/
	if filepath != "" {

		/*Comprobacion de que tipo de archivo voy a mostrar
		dependiendo de la ruta*/
		entry, err := tree.FindEntry(filepath)

		if err != nil {
			return nil, err
		}
		/*Comprobacion de que si lo que viene de la ruta es una carpeta
		  go-git reconoce el entry.mode 0040000 como carpeta(DIR)
		*/

		if entry.Mode.String() == "0040000" {
			Tree_entry, err := tree.Tree(filepath)

			if err != nil {

				return nil, object.ErrDirectoryNotFound

			} else {
				for _, entry := range Tree_entry.Entries {

					EntryPaths = append(EntryPaths, entry.Name)
				}
				fmt.Println(Tree_entry)

				if err != nil {
					return nil, object.ErrFileNotFound
				}
			}
			/*Comprueba y muestra el archivo si cumple la condicion*/
		} else if entry.Mode.String() == "0100644" {

			/*Llena el array de rutas de los archivos*/

			tree.Files().ForEach(func(f *object.File) error {

				TreeEntries = append(TreeEntries, f.Name)
				return nil
			})

			treefile, err := tree.File(filepath)

			if err != nil {

			}
			for i := 0; i < len(TreeEntries); i++ {

				if TreeEntries[i] == filepath {

					content, err := treefile.Contents()

					fileLines, err := treefile.Lines()

					fmt.Println(fileLines)

					datafile = append(datafile, content)

					if err != nil {

					}
					fmt.Printf("FilePath and parameter: %s    %s\n", filepath[i], filepath)
					fmt.Println(content)
					break
				}

			}

			return datafile, err

		}

	} else {

		for _, entry := range tree.Entries {

			EntryPaths = append(EntryPaths, entry.Name)
		}

	}

	return EntryPaths, err

}
