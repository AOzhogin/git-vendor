package vendors

import (
	"fmt"
	"github.com/AOzhogin/git-vendor/internal/views"
	"strings"
)

var (
	defaultWorkDir = "git-vendor"
	vendorList     = "modules.yaml"
)

type Storage interface {
	NeedCommit() bool
	Load(list *views.VendorsList) error
	Save(list *views.VendorsList) error
}

type Git interface {
	IsInstalled() bool
	Status() error
	Add(workDir string, repository string, ref string, message string) error
	Commit(message string) error
	Update(workDir string, repository string, currRef string, message string, targetRef string) error
	UpStream(workDir string, repository string, ref string) error
	Remove(workDir string, message string) error
}

type Vendors struct {
	dir     string
	git     Git
	storage Storage
	Vendors views.VendorsList
}

func New(d string, g Git, s Storage) Vendors {

	v := Vendors{
		dir:     d,
		git:     g,
		storage: s,
		Vendors: make(views.VendorsList),
	}

	if d == "" {
		v.dir = defaultWorkDir
	}

	v.storage.Load(&v.Vendors)

	return v
}

func (v *Vendors) Add(name string, repository string, ref string) error {

	if ref == "" {
		ref = "master"
	}

	dir := v.getTargetDir(repository)

	message := fmt.Sprintf("Add \"%s\" from \"%s@%s\"\n\ngit-vendor-name: %s\ngit-vendor-dir: %s\ngit-vendor-repository: %s\ngit-vendor-ref: %s\n",
		name, repository, ref, name, dir, repository, ref)

	if err := v.git.Add(dir, repository, ref, message); err != nil {
		return err
	}

	v.Vendors[name] = views.Entry{
		Name:       name,
		Dir:        dir,
		Repository: repository,
		Ref:        ref,
	}

	if err := v.storage.Save(&v.Vendors); err != nil {
		return err
	}

	return v.git.Commit(message)

}

func (v *Vendors) UpStream(name string, repository string, ref string) error {

	vendor, exists := v.Vendors[name]
	if !exists {
		return fmt.Errorf("vendor [%s]: not exists", name)
	}

	var (
		curRef  = ref
		curRepo = repository
	)

	if curRef == "" {
		curRef = vendor.Ref
	}

	if curRepo == "" {
		curRepo = vendor.Repository
	}

	err := v.git.UpStream(vendor.Dir, curRepo, curRef)
	if err != nil {
		return err
	}

	return nil

}

func (v *Vendors) Update(name string, targetRef string) error {

	if err := v.git.Status(); err != nil {
		return err
	}

	if targetRef == "" {
		targetRef = "master"
	}
	vendor, exists := v.Vendors[name]
	if !exists {
		return fmt.Errorf("vendor [%s]: not exists", name)
	}

	message := fmt.Sprintf("Update \"%s\" from \"%s@%s\"\n\ngit-vendor-name: %s\ngit-vendor-dir: %s\ngit-vendor-repository: %s\ngit-vendor-ref: %s\n",
		name, vendor.Repository, targetRef, name, vendor.Dir, vendor.Repository, targetRef)

	if err := v.git.Update(vendor.Dir, vendor.Repository, vendor.Ref, message, targetRef); err != nil {
		return err
	}

	vendor.Ref = targetRef

	v.Vendors[name] = vendor

	if err := v.storage.Save(&v.Vendors); err != nil {
		return err
	}

	return v.git.Commit(message)

}

func (v *Vendors) Remove(name string) error {

	message := fmt.Sprintf("Removing \"%s\" from \"%s\"", name, v.dir)

	vendor, exists := v.Vendors[name]
	if !exists {
		return fmt.Errorf("vendor [%s]: not exists", name)
	}

	if err := v.git.Remove(vendor.Dir, message); err != nil {
		return err
	}

	delete(v.Vendors, name)

	if err := v.storage.Save(&v.Vendors); err != nil {
		return err
	}

	return v.git.Commit(message)

}

//func (v *Vendors) Load() error {
//
//	data, err := os.ReadFile(fmt.Sprintf("%s/%s", v.dir, vendorList))
//	if err != nil {
//		if errors.Is(err, os.ErrNotExist) {
//			return nil
//		}
//		return fmt.Errorf("read modules file: %w", err)
//	}
//
//	err = yaml.Unmarshal(data, &v)
//	if err != nil {
//		return err
//	}
//
//	return nil
//
//}
//
//func (v *Vendors) Save() error {
//
//	data, err := yaml.Marshal(&v)
//	if err != nil {
//		return err
//	}
//
//	err = os.WriteFile(fmt.Sprintf("%s/%s", v.dir, vendorList), data, 0600)
//	if err != nil {
//		return fmt.Errorf("write modules file: %w", err)
//	}
//
//	return nil
//
//}

func (v *Vendors) getTargetDir(repository string) string {

	switch strings.Contains(repository, "@") {
	case true:
		return fmt.Sprintf("%s/%s", v.dir, strings.ReplaceAll(strings.TrimPrefix(strings.TrimSuffix(repository, ".git"), "git@"), ":", "/"))
	case false:
		if strings.Contains(repository, "http://") {
			return fmt.Sprintf("%s/%s", v.dir, strings.TrimPrefix(strings.TrimSuffix(repository, ".git"), "http://"))
		}
		return fmt.Sprintf("%s/%s", v.dir, strings.TrimPrefix(strings.TrimSuffix(repository, ".git"), "https://"))
	default:
		return "unknown"
	}

}
