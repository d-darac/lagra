package cmd

import (
	"context"
	"fmt"

	configcmd "github.com/d-darac/lagra/cmd/lagra-mock/cmd/config"
	groupscmd "github.com/d-darac/lagra/cmd/lagra-mock/cmd/groups"
	itemscmd "github.com/d-darac/lagra/cmd/lagra-mock/cmd/items"
	"github.com/d-darac/lagra/cmd/lagra-mock/cmd/shared"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "lagra-mock",
	Long: `lgra-mock is used to quickly insert mock resources into a configured database.

Start by running 'lagra-mock config' to configure database connection parameters,
then run 'lagra-mock help' to see the list of available commands.
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run:          func(cmd *cobra.Command, args []string) { fmt.Println(cmd.Long) },
	SilenceUsage: true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return rootCmd.ExecuteContext(context.Background())
}

func init() {
	cobra.OnInitialize(shared.InitConfig)
	rootCmd.PersistentFlags().StringVar(&shared.CfgFile, "config", "", "config file (default is $HOME/.lagra-mock.yaml)")
	rootCmd.AddCommand(groupscmd.GroupsCmd, itemscmd.ItemsCmd, configcmd.ConfigCmd)
}
