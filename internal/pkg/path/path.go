package path

import "regexp"

type Path struct {
	User     string
	Host     string
	FilePath string
}

var pattern = regexp.MustCompile(`^([^@]+)@([^:]+):(.+)$`)

func Parse(path string) Path {
	matchs := pattern.FindStringSubmatch(path)
	return Path{
		User:     matchs[1],
		Host:     matchs[2],
		FilePath: matchs[3],
	}
}
