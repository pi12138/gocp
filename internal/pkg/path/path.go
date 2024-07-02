package path

import (
	"os/user"
	"strings"
)

type Path struct {
	User     string
	Host     string
	FilePath string
}

func Parse(path string) *Path {
	p := new(Path)
	idx1 := strings.IndexByte(path, '@')
	idx2 := strings.IndexByte(path, ':')
	if idx1 != -1 && idx2 != -1 { // user@host:path
		p.User = path[0:idx1]
		p.Host = path[idx1+1 : idx2]
		p.FilePath = path[idx2+1:]
	} else if idx1 != -1 && idx2 == -1 { // user@path
		p.User = path[0:idx1]
		p.FilePath = path[idx1+1:]
	} else if idx1 == -1 && idx2 != -1 { // host:path
		p.Host = path[0:idx2]
		p.FilePath = path[idx2+1:]
	} else { // path
		p.FilePath = path
	}
	if p.User == "" {
		user, _ := user.Current()
		p.User = user.Username
	}
	if p.Host == "" {
		p.Host = "localhost"
	}
	return p
}
