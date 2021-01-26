package api_git

func pullRequest(pullRequest PRCreate) (*PRCreate, error) {
	var diffRequest = DiffRequest{
		pullRequest.UrlRepoReceivePR,
		"",
		pullRequest.UrlRepoCreatePR,
		pullRequest.CommitHash,
	}

	result, err := diffTreeRepos(diffRequest)
	if err == nil {
		pullRequest.Patch = result.String()
		return &pullRequest, nil
	}
	if err != nil {
		return nil, err
	}

	return &pullRequest, err
}
