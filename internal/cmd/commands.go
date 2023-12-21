package cmd

import (
	"fmt"
	"github.com/AOzhogin/git-vendor/internal/git"
	"github.com/AOzhogin/git-vendor/internal/storage/file"
	"github.com/AOzhogin/git-vendor/internal/vendors"
	"github.com/spf13/cobra"
)

var (
	g *git.Git
	s vendors.Storage
	v vendors.Vendors

	err     error
	workDir string = "git-vendor"
	name    string
	ref     string
	repo    string

	RootCmd = &cobra.Command{
		Use:   "git-vendor",
		Short: "git subtree wrapper",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	listCmd = &cobra.Command{
		Use:   "list [name]",
		Short: "List vendors",
		Args: func(cmd *cobra.Command, args []string) error {

			if len(args) < 1 {
				return nil
			}

			name = args[0]

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			s = file.NewStorage(workDir)

			a := vendors.New(workDir, g, s)
			for _, vendor := range a.Vendors {
				if (vendor.Name == name) || (name == "") {
					vendor.PettyPrint()
				}
			}
			s.Save(&a.Vendors)
		},
	}

	updateCmd = &cobra.Command{
		Use:   "update <name> [ref]",
		Short: "Update vendor",
		Args: func(cmd *cobra.Command, args []string) error {

			if len(args) < 1 {
				return fmt.Errorf("Incorrect options provided: git vendor update <name> [<ref>]")
			}

			name = args[0]

			ref = "master"
			if len(args) == 2 {
				ref = args[1]
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			s = file.NewStorage(workDir)
			a := vendors.New(workDir, g, s)
			if err = a.Update(name, ref); err != nil {
				println(err.Error())
			}
		},
	}

	upStreamCmd = &cobra.Command{
		Use:   "upstream <name> [ref] [repository]",
		Short: "UpStream vendor",
		Args: func(cmd *cobra.Command, args []string) error {

			if len(args) < 1 {
				return fmt.Errorf("Incorrect options provided: git vendor upstream <name> [ref] [repository]")
			}

			name = args[0]

			ref = "master"

			if len(args) >= 2 {
				ref = args[1]
			}

			if len(args) > 2 {
				repo = args[2]
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			s = file.NewStorage(workDir)
			a := vendors.New(workDir, g, s)

			if err = a.UpStream(name, repo, ref); err != nil {
				println(err.Error())
			}
		},
	}

	addCmd = &cobra.Command{
		Use:   "add <name> <repository> [<ref>]",
		Short: "add new vendor",
		Args: func(cmd *cobra.Command, args []string) error {

			if len(args) < 2 {
				return fmt.Errorf("Incorrect options provided: git vendor add <name> <repository> [ref]")
			}

			name = args[0]

			ref = "master"
			if len(args) > 2 {
				ref = args[2]
			}

			repo = ""

			if len(args) >= 2 {
				repo = args[1]
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			s = file.NewStorage(workDir)
			a := vendors.New(workDir, g, s)

			if err = a.Add(name, repo, ref); err != nil {
				println(err.Error())
			}
		},
	}

	removeCmd = &cobra.Command{
		Use:   "remove <name>",
		Short: "Remove vendor",
		Args: func(cmd *cobra.Command, args []string) error {

			if len(args) < 1 {
				return fmt.Errorf("Incorrect options provided: git vendor remove <name>")
			}

			name = args[0]

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			s = file.NewStorage(workDir)
			a := vendors.New(workDir, g, s)
			if err = a.Remove(name); err != nil {
				println(err.Error())
			}
		},
	}
)

func init() {
	RootCmd.AddCommand(listCmd)
	RootCmd.AddCommand(updateCmd)
	RootCmd.AddCommand(upStreamCmd)
	RootCmd.AddCommand(addCmd)
	RootCmd.AddCommand(removeCmd)
	g = git.New()
	if !g.IsInstalled() {
		fmt.Println("Error: git command not found in PATH.")
		return
	}

}
