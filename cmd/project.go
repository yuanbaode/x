package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yuanbaode/x/service/project"
)

var projectCmd = &cobra.Command{
	Use:     "project",
	Short:   "create project",
	Example: "",

	Run: createProject,
}

var (
	projectName string
)

func init() {
	projectCmd.Flags().StringVarP(&projectName, "name", "", "demo", "specify the name of the project")
	rootCmd.AddCommand(projectCmd)
}

func createProject(cmd *cobra.Command, args []string) {
	p := project.NewProject(projectName)
	_ = p.CreateProject()
}
