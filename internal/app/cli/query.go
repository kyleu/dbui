package cli

import (
	"github.com/kyleu/dbui/internal/app/conn"
	"github.com/spf13/cobra"
)

func NewQueryCommand() *cobra.Command {
	var connNameArg string
	var inputArg string
	var outputArg string

	cmd := &cobra.Command{
		Use:     "query",
		Aliases: []string{"q"},
		Short:   "Runs the provided sql, displaying or saving the result",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if connNameArg == "" {
				connNameArg = "default"
			}
			if inputArg == "" {
				inputArg = "show tables"
			}
			if outputArg == "" {
				outputArg = "table"
			}
			return conn.Sandbox(connNameArg, inputArg, outputArg)
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&connNameArg, "conn", "c", "", "Connection name")
	flags.StringVarP(&inputArg, "input", "i", "", "???")
	flags.StringVarP(&outputArg, "output", "o", "", "???")

	return cmd
}