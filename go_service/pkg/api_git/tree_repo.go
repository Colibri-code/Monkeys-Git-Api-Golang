package api_git

// change package to main and put it inside the main.go file if u need to see
// the running example

import (
	"regexp"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

/*Toma todos los files que tiene el tree Head del repositorio y los
retorna en un []string*/

type EntryInfo struct {
	EntryName       string   `json:"EntryName"`
	ModeFile        string   `json:"FileMode"`
	Data            []string `json:"Data, omitempty"`
	isJson          bool     `json:"isJson, omitempty"`
	JsonDataContent string   `json:"JsonContent, omitempty"`
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
		//Abre el repositorio que viene en la request Name:
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
			  go-git reconoce el entry.mode 0040000 como carpeta(DIR) o si es un Archivo entry.mode 0100644 Regular
			*/
			entryfilemode, err := TreeEntryType(entry, TreeHead, filepath)

			if err != nil {
				return nil, err
			}
			/*Retorno la respuesta
			*Todo el contenido << estructura >>

			* EntryName
			* FileMode
			* Data
			* IsJson
			* JsonContent

			 */
			return entryfilemode, err

			//	fmt.Println(entryfilemode)

		} else {

			for _, entry := range TreeHead.Entries {

				BlobData.Data = append(BlobData.Data, entry.Name)
			}
			BlobData.ModeFile = "Dir"

		}

		return &BlobData, err

	} else {
		return nil, git.ErrRepositoryNotExists
	}

}

/*
* entry es el archivo al que quiero obtener
* master tree es el nodo (archivo) root o repo Raiz
* path es la ruta concatenda del archivo o carpeta (nodo actual) example: miproyecto/src
 */
func TreeEntryType(entry *object.TreeEntry, masterTree *object.Tree, path string) (*EntryInfo, error) {

	entrymode := entry.Mode.String()
	entryName := entry.Name

	//Variable bool que comprueba si el archivo es un Json
	isJson, _ := regexp.MatchString(".json", entryName)

	//toma todos lo que tiene una carpeta
	//	var TreeEntries []string

	BlobData := EntryInfo{}

	switch entrymode {
	/*Comprobacion de que si lo que viene de la ruta es una carpeta
	  go-git reconoce el entry.mode 0040000 como carpeta(DIR)
	*/
	case "0040000":

		// de la carpeta que estoy tomando por ruta tome lo que tiene adentro(contenido)
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
	/*Comprobacion de que si lo que viene de la ruta es una carpeta
	  go-git reconoce el entry.mode 0100644 como Regular
	*/
	case "0100644":

		/*Llena el array de rutas de los archivos*/

		/*	masterTree.Files().ForEach(func(f *object.File) error {

			TreeEntries = append(TreeEntries, f.Name)
			return nil
		})*/

		treefile, err := masterTree.File(path)

		if err != nil {
			return nil, err
		} else {
			if isJson {
				//	JsonContent, err := treefile.Lines()

				//fmt.Println(treefile.Lines())

				//fmt.Println("-----------")

				//fmt.Println(treefile.Contents())
				//BlobData.Data, err = treefile.Lines()
				BlobData.EntryName = entryName
				BlobData.isJson = true
				BlobData.ModeFile = "json"
				//BlobData.Data, err = treefile.Contents()
				BlobData.JsonDataContent, err = treefile.Contents()

				if err != nil {
					return nil, err
				}
				/* 		var content []*EntryInfo

				contenido := make(map[string][]*EntryInfo)

				for _, p := range content {
					for _, l := range p.Data {
						contenido[l] = append(contenido[l], p)
					}
				}
				fmt.Println(contenido) */

				///justString := strings.Join(JsonContent, "")

				///fmt.Println(justString)

				///BlobData.Data = justString

				return &BlobData, err

			} else {

				//De toda la lista de rutas de archivos que existen
				//Busco la que viene por parametro para enviar su
				//Contenido por medio de string[] var

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

//Contenido de un Archivo
/* func JsonContentData(RepoPath string, JsonPath string) (*JsonEntryInfo, error) {

	JsonEntry := JsonEntryInfo{}

	repo, err := OpenRepository(RepoPath)

	if err != nil {
		return nil, err
	}

	TreeHead := TreeCommitHead(repo)

	JsonFile, err := TreeHead.File(JsonPath)

	if err != nil {
		return nil, object.ErrEntryNotFound
	}

	JsonEntry.DataContent, err = JsonFile.Contents()

	JsonEntry.EntryName = JsonFile.Name

	return &JsonEntry, err
} */
