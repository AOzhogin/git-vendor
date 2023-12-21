package main

import (
	"fmt"
	"github.com/AOzhogin/git-vendor/internal/cmd"
)

func main() {

	err := cmd.RootCmd.Execute()
	if err != nil && err.Error() != "" {
		println(fmt.Sprintf("Error: %s", err.Error()))
	}

}
