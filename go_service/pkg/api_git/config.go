package api_git

import (
	format "github.com/go-git/go-git/v5/plumbing/format/config"
)

type Configure struct {
	Core struct {
		// IsBare if true this repository is assumed to be bare and has no
		// working directory associated with it.
		IsBare bool
		// Worktree is the path to the root of the working tree.
		Worktree string
		// CommentChar is the character indicating the start of a
		// comment for commands like commit and tag
		CommentChar string
	}

	User struct {
		// Name is the personal name of the author and the commiter of a commit.
		Name string
		// Email is the email of the author and the commiter of a commit.
		Email string
	}

	Author struct {
		// Name is the personal name of the author of a commit.
		Name string
		// Email is the email of the author of a commit.
		Email string
	}

	Committer struct {
		// Name is the personal name of the commiter of a commit.
		Name string
		// Email is the email of the  the commiter of a commit.
		Email string
	}

	Pack struct {
		// Window controls the size of the sliding window for delta
		// compression.  The default is 10.  A value of 0 turns off
		// delta compression entirely.
		Window uint
	}

	Init struct {
		// DefaultBranch Allows overriding the default branch name
		// e.g. when initializing a new repository or when cloning
		// an empty repository.
		DefaultBranch string
	}

	// Remotes list of repository remotes, the key of the map is the name
	// of the remote, should equal to RemoteConfig.Name.
	Remotes map[string]*RemoteConfig

	// Branches list of branches, the key is the branch name and should
	// equal Branch.Name
	Branches map[string]*Branch
	// URLs list of url rewrite rules, if repo url starts with URL.InsteadOf value, it will be replaced with the
	// key instead.
	URLs map[string]*URL
	// Raw contains the raw information of a config file. The main goal is
	// preserve the parsed information from the original format, to avoid
	// dropping unsupported fields.
	Raw *format.Config
}
