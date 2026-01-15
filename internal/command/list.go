package command

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/basecamp/amar/internal/docker"
)

type ListCommand struct {
	root *RootCommand
	cmd  *cobra.Command
}

func NewListCommand(root *RootCommand) *ListCommand {
	l := &ListCommand{root: root}
	l.cmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List installed applications",
		RunE:    WithNamespace(l.run),
	}
	return l
}

func (l *ListCommand) Command() *cobra.Command {
	return l.cmd
}

// Private

func (l *ListCommand) run(ns *docker.Namespace, cmd *cobra.Command, args []string) error {
	for _, app := range ns.Applications() {
		fmt.Println(app.Settings.Name)
	}

	return nil
}
