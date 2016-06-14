package commands

import (
	"github.com/spf13/cobra"

	"github.com/nanobox-io/nanobox/processor"
)

var (
	alias string

	// LinkCmd ...
	LinkCmd = &cobra.Command{
		Use:   "link",
		Short: "Manages links between dev & production apps.",
		Long:  ``,
	}

	// LinkAddCmd ...
	LinkAddCmd = &cobra.Command{
		Use:   "add",
		Short: "Adds a new link between a dev & production app.",
		Long: `
Adds a new link between a dev and production app. A dev app can
be linked to multiple production apps. Each link needs an alias.
If no alias is provided, 'default' is assumed.
		`,

		Run: func(ccmd *cobra.Command, args []string) {
			processor.DefaultConfig.Meta["name"] = app
			processor.DefaultConfig.Meta["alias"] = alias
			handleError(processor.Run("link_add", processor.DefaultConfig))
		},
	}

	// LinkListCmd ...
	LinkListCmd = &cobra.Command{
		Use:   "list",
		Short: "Lists all links for the current dev app.",
		Long:  ``,

		Run: func(ccmd *cobra.Command, args []string) {
			handleError(processor.Run("link_list", processor.DefaultConfig))
		},
	}

	// LinkRemoveCmd ...
	LinkRemoveCmd = &cobra.Command{
		Use:   "remove",
		Short: "Removes a link between a dev & production app.",
		Long:  ``,

		Run: func(ccmd *cobra.Command, args []string) {
			processor.DefaultConfig.Meta["alias"] = alias
			handleError(processor.Run("link_remove", processor.DefaultConfig))
		},
	}
)

//
func init() {
	LinkAddCmd.Flags().StringVarP(&app, "app_name", "n", "", "app name")
	LinkCmd.PersistentFlags().StringVarP(&alias, "alias", "a", "", "alias")

	LinkCmd.AddCommand(LinkAddCmd)
	LinkCmd.AddCommand(LinkListCmd)
	LinkCmd.AddCommand(LinkRemoveCmd)
}
