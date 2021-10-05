package api_git

// change package to main and put it inside the main.go file if u need to see
// the running example

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

/*Toma todos los files que tiene el tree Head del repositorio y los
retorna en un []string*/

type EntryInfo struct {
	EntryName string   `json:"EntryName"`
	ModeFile  string   `json:"FileMode"`
	Data      []string `json:"Data"`
}

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

func ContentTreeData(repoPath string, filepath string) (*EntryInfo, error) {

	ConcatRepoPath := baseRepoDir + repoPath + ".git"

	BlobData := EntryInfo{}
	if repoPath != "" {
		repo, err := git.PlainOpen(ConcatRepoPath)

		if err != nil {
			return nil, git.ErrRepositoryNotExists
		}

		/*Obtengo el Tree Head del repositorio*/
		TreeHead := TreeCommitHead(repo)

		/*Comprobacion de que la path no venga vacia, si viene vacia se lista el
		main tree del repositorio*/
		if filepath != "" {

			/*Comprobacion de que tipo de archivo voy a mostrar
			dependiendo de la ruta*/
			entry, err := TreeHead.FindEntry(filepath)

			if err != nil {
				return nil, err
			}
			/*Comprobacion de que si lo que viene de la ruta es una carpeta
			  go-git reconoce el entry.mode 0040000 como carpeta(DIR)
			*/
			entryfilemode, err := TreeEntryType(entry, TreeHead, filepath)

			if err != nil {
				return nil, err
			}

			return entryfilemode, err

			//	fmt.Println(entryfilemode)

		} else {

			for _, entry := range TreeHead.Entries {

				BlobData.Data = append(BlobData.Data, entry.Name)
			}

		}

		return &BlobData, err

	} else {
		return nil, git.ErrRepositoryNotExists
	}

}

func TreeEntryType(entry *object.TreeEntry, masterTree *object.Tree, path string) (*EntryInfo, error) {

	entrymode := entry.Mode.String()
	entryName := entry.Name

	var TreeEntries []string

	BlobData := EntryInfo{}
	switch entrymode {
	/*Comprobacion de que si lo que viene de la ruta es una carpeta
	  go-git reconoce el entry.mode 0040000 como carpeta(DIR)
	*/
	case "0040000":

		Tree_entry, err := masterTree.Tree(path)

		if err != nil {

			return nil, object.ErrDirectoryNotFound

		} else {
			for _, entry := range Tree_entry.Entries {

				BlobData.Data = append(BlobData.Data, entry.Name)
				//	EntryPaths = append(EntryPaths, entry.Name)
			}

			if err != nil {
				return nil, object.ErrFileNotFound
			}
			BlobData.ModeFile = "Dir"
			BlobData.EntryName = entryName
		}
		return &BlobData, nil

	case "0100644":

		/*Llena el array de rutas de los archivos*/

		masterTree.Files().ForEach(func(f *object.File) error {

			TreeEntries = append(TreeEntries, f.Name)
			return nil
		})

		treefile, err := masterTree.File(path)

		if err != nil {
			return nil, err
		}
		//De toda la lista de rutas de archivos que existen
		//Busco la que viene por parametro para enviar su
		//Contenido por medio de string[] var
		for i := 0; i < len(TreeEntries); i++ {

			if TreeEntries[i] == path {

				BlobData.Data, err = treefile.Lines()

				if err != nil {
					return nil, err
				}
				BlobData.ModeFile = "Regular"
				BlobData.EntryName = entryName

				return &BlobData, err

			}

		}

	}

	return &BlobData, nil
}
