package file

import (
	"errors"
	"fmt"
	"github.com/AOzhogin/git-vendor/internal/views"
	"gopkg.in/yaml.v3"
	"os"
)

var (
	defaultWorkDir = "git-vendor"
	vendorList     = "modules.yaml"
)

type fileStorage struct {
	workDir string
}

func NewStorage(d string) *fileStorage {

	s := &fileStorage{
		workDir: d,
	}

	if d == "" {
		s.workDir = defaultWorkDir
	}

	return s

}

func (s *fileStorage) NeedCommit() bool {
	return true
}

func (s *fileStorage) Load(list *views.VendorsList) error {

	data, err := os.ReadFile(fmt.Sprintf("%s/%s", s.workDir, vendorList))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return fmt.Errorf("read modules file: %w", err)
	}

	err = yaml.Unmarshal(data, list)
	if err != nil {
		return err
	}

	return nil

}

func (s *fileStorage) Save(list *views.VendorsList) error {

	data, err := yaml.Marshal(list)
	if err != nil {
		return err
	}

	err = os.WriteFile(fmt.Sprintf("%s/%s", s.workDir, vendorList), data, 0600)
	if err != nil {
		return fmt.Errorf("write modules file: %w", err)
	}

	return nil
}
