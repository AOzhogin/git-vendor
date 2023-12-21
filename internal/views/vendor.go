package views

import "fmt"

type VendorsList map[string]Entry //`yaml:"vendors"`

type Entry struct {
	Name       string `yaml:"name"`
	Dir        string `yaml:"dir"`
	Repository string `yaml:"repository"`
	Ref        string `yaml:"ref"`
	Commit     string `yaml:"-"`
}

func (e Entry) IsComplete() bool {
	return e.Name != "" && e.Dir != "" && e.Repository != "" && e.Ref != ""
}

func (e Entry) PettyPrint() {
	fmt.Printf("%s@%s:\n", e.Name, e.Ref)
	fmt.Printf("\tname:\t%s\n", e.Name)
	fmt.Printf("\tdir:\t%s\n", e.Dir)
	fmt.Printf("\trepo:\t%s\n", e.Repository)
	fmt.Printf("\tref:\t%s\n", e.Ref)
	if e.Commit != "" {
		fmt.Printf("\tcommit:\t%s\n", e.Commit)
	}
	fmt.Println()
}
