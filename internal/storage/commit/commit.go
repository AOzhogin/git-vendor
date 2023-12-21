package commit

import (
	"github.com/AOzhogin/git-vendor/internal/views"
	"strings"
)

type Git interface {
	Log() ([]byte, error)
}

type commitStorage struct {
	git Git
}

func NewStorage(g Git) *commitStorage {

	s := &commitStorage{
		git: g,
	}

	return s

}

func (s *commitStorage) NeedCommit() bool {
	return false
}

func (s *commitStorage) Load(list *views.VendorsList) error {

	output, err := s.git.Log()
	if err != nil {
		return err
	}

	lines := strings.Split(string(output), "\n")

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			continue
		}
		if strings.Contains(line, "START ") {
			var entry = views.Entry{
				Commit: strings.TrimSpace(strings.TrimPrefix(line, "START ")),
			}

			for j := i + 1; j < len(lines); j++ {
				if lines[j] == "" {
					continue
				}
				if strings.Contains(lines[j], "git-vendor-dir:") {
					entry.Dir = strings.TrimSpace(strings.TrimPrefix(lines[j], "git-vendor-dir:"))
				} else if strings.Contains(lines[j], "git-vendor-repository:") {
					entry.Repository = strings.TrimSpace(strings.TrimPrefix(lines[j], "git-vendor-repository:"))
				} else if strings.Contains(lines[j], "git-vendor-ref: ") {
					entry.Ref = strings.TrimPrefix(lines[j], "git-vendor-ref: ")
				} else if strings.Contains(lines[j], "git-vendor-name: ") {
					entry.Name = strings.TrimPrefix(lines[j], "git-vendor-name: ")
				}
			}
			if entry.IsComplete() {
				if _, exists := (*list)[entry.Name]; !exists {
					(*list)[entry.Name] = entry
				}
			}

			for i++; i < len(lines); i++ {
				if lines[i] == "END" {
					break
				}
			}
		}
	}

	return nil

}

func (s *commitStorage) Save(_ *views.VendorsList) error {

	return nil

}
