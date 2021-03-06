package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// appDelCmd deletes an application profile.
var (
	appDelCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete an application profile",
		Long:  `Delete an application`,
		Run:   runAppDel,
	}
)

func init() {
	appCmd.AddCommand(appDelCmd)
	appDelCmd.Flags().StringVar(&aciFlags.tenant, "tenant", "", "Name of tenant")
	appDelCmd.Flags().StringVar(&aciFlags.name, "name", "", "Name of app")
	appDelCmd.MarkFlagRequired("tenant")
	appDelCmd.MarkFlagRequired("name")
}

func runAppDel(cmd *cobra.Command, args []string) {
	if len(aciFlags.tenant) == 0 {
		cmd.Help()
		return
	}
	if len(aciFlags.name) == 0 {
		cmd.Help()
		return
	}
	if err := validateArgs(cmd, args); err != nil {
		return
	}

	client := mustClientFromCmd(cmd)

	err := client.Login()
	if err != nil {
		exitWithError(ExitError, err)
	}

	err = client.ApplicationProfileDel(aciFlags.tenant, aciFlags.name)
	if err != nil {
		exitWithError(ExitError, err)
	}

	fmt.Printf("Application profile %s deleted.\n", aciFlags.name)

	err = client.Logout()
	if err != nil {
		exitWithError(ExitError, err)
	}
}
