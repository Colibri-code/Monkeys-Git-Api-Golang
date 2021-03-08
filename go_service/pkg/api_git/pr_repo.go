package api_git

/*
*Funcion que toma como parametros los datos que vienen
del la ruta para tomar y retornar los cambios que se han realizado en el
Pull request
*/
func pullRequest(pullRequest PRCreate) (*PRCreate, error) {
	var diffRequest = DiffRequest{
		pullRequest.UrlRepoReceivePR,
		"",
		pullRequest.UrlRepoCreatePR,
		pullRequest.CommitHash,
	}
	//Funcion que comprueba las diferencias que hay la rama main
	// con la rama que esta actualizando
	result, err := diffTreeRepos(diffRequest)
	if err == nil {
		//Guarda en el atributo Path, la accion que se esta
		//realizando dentro del repositorio principal
		pullRequest.Patch = result.String()
		return &pullRequest, err
	}
	if err != nil {
		return nil, err
	}

	return &pullRequest, err
}
