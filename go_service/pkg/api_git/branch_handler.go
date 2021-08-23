package api_git

import "github.com/go-git/go-git/v5/plumbing"

type RemoteConfig struct {

	// Name of the remote
	Name string

	// URLs the URLs of a remote repository. It must be non-empty. Fetch will
	// always use the first URL, while push will use all of them.
	URLs []string

	// insteadOfRulesApplied have urls been modified
	insteadOfRulesApplied bool

	originalURLs []string
}

type URL struct {
	// Name new base url
	Name string
	// Any URL that starts with this value will be rewritten to start, instead, with <base>.
	// When more than one insteadOf strings match a given URL, the longest match is used.
	InsteadOf string

	// raw representation of the subsection, filled by marshal or unmarshal are
	// called.
	//raw *format.Subsection
}
type Config struct {
	Core struct {
		IsBare bool

		Worktree string

		CommentChar string
	}

	User struct {
		Name string

		Email string
	}

	Author struct {
		Name string

		Email string
	}

	Committer struct {
		Name string

		Email string
	}

	Pack struct {
		Window uint
	}

	Init struct {
		DefaultBranch string
	}

	Remotes  map[string]*RemoteConfig
	Branches map[string]*Branch
	URLs     map[string]*URL
}

type Branch struct {
	Name   string
	Remote string
	Merge  plumbing.ReferenceName
	Rebase string
}
