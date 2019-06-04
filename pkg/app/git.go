package app

import (
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

type GitHashProvider interface {
	getCurrentGitHash() string
}

type goGitHashProvider struct {
}

func (gitProvider *goGitHashProvider) getCurrentGitHash() string {

	path := "."
	revision := "HEAD"

	r, err := git.PlainOpen(path)
	if err != nil {
		panic(err)
	}
	h, err := r.ResolveRevision(plumbing.Revision(revision))
	if err != nil {
		panic(err)
	}
	return h.String()
}
