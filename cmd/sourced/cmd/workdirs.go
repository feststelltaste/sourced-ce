package cmd

import (
	"fmt"

	"github.com/src-d/sourced-ce/cmd/sourced/compose/workdir"
)

type workdirsCmd struct {
	Command `name:"workdirs" short-description:"List working directories" long-description:"List previously initialized working directories."`
}

func (c *workdirsCmd) Execute(args []string) error {
	dirs, err := workdir.ListRepoDirs()
	if err != nil {
		return err
	}

	// ignore errors if active dir doesn't exist or unavailable
	active, _ := workdir.ActiveRepoDir()
	for _, dir := range dirs {
		if dir == active {
			fmt.Printf("* %s\n", dir)
		} else {
			fmt.Printf("  %s\n", dir)
		}
	}

	return nil
}

func init() {
	rootCmd.AddCommand(&workdirsCmd{})
}