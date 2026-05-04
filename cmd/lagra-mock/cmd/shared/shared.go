package shared

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/d-darac/lagra/internal/database"
	"github.com/d-darac/lagra/pkg/id"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Key string

const (
	KeyAccountID  Key = "account_id"
	KeyDBHost     Key = "db_host"
	KeyDBName     Key = "db_name"
	KeyDBPassword Key = "db_password"
	KeyDBPort     Key = "db_port"
	KeyDBUser     Key = "db_user"
)

var (
	CfgFile   string
	AccountID id.ID
)

func readViperConfig(paths []string) error {
	for _, p := range paths {
		if _, err := os.Stat(p); err == nil {
			viper.SetConfigFile(p)
			break
		}
	}
	return viper.ReadInConfig()
}

func InitConfig() {
	viper.SetDefault(string(KeyDBHost), "")
	viper.SetDefault(string(KeyDBName), "")
	viper.SetDefault(string(KeyDBPassword), "")
	viper.SetDefault(string(KeyDBPort), "")
	viper.SetDefault(string(KeyDBUser), "")

	if CfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(filepath.Clean(CfgFile))
		cobra.CheckErr(viper.ReadInConfig())
	} else {
		// find home dir
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// collect paths where existing config files may be located
		var configPaths []string

		// first check XDG_CONFIG_HOME if set
		xdgConfigHome := os.Getenv("XDG_CONFIG_HOME")
		var xdgEnvPath string
		if xdgConfigHome != "" {
			xdgEnvPath = filepath.Join(xdgConfigHome, "lagra-mock", "config.yaml")
			configPaths = append(configPaths, xdgEnvPath)
		}

		// then check legacy hard-coded "XDG" path, then home dotfile
		xdgLegacyPath := filepath.Join(home, ".config", "lagra-mock", "config.yaml")
		homeDotfilePath := filepath.Join(home, ".lagra-mock.yaml")

		configPaths = append(configPaths, xdgLegacyPath)
		configPaths = append(configPaths, homeDotfilePath)

		if err := readViperConfig(configPaths); err != nil {
			// no existing config found; try to create a new one
			// respect XDG_CONFIG_HOME if set, otherwise use dotfile in home dir
			var newConfigPath string
			if xdgEnvPath != "" {
				newConfigPath = xdgEnvPath
				cobra.CheckErr(os.MkdirAll(filepath.Dir(newConfigPath), 0o755))
			} else {
				newConfigPath = homeDotfilePath
			}

			cobra.CheckErr(viper.SafeWriteConfigAs(newConfigPath))
			viper.SetConfigFile(newConfigPath)
			cobra.CheckErr(viper.ReadInConfig())
		}
	}
}

func RequireConfig(cmd *cobra.Command, args []string) {
	promptConfigAndExit := func(condition bool) {
		if condition {
			fmt.Fprintln(os.Stderr, "You must configure lagra-mock before using.")
			fmt.Fprintln(os.Stderr, "Run 'lagra-mock config' first.")
			os.Exit(1)
		}
	}

	config := map[string]string{
		string(KeyDBHost):     viper.GetString(string(KeyDBHost)),
		string(KeyDBName):     viper.GetString(string(KeyDBName)),
		string(KeyDBPassword): viper.GetString(string(KeyDBPassword)),
		string(KeyDBPort):     viper.GetString(string(KeyDBPort)),
		string(KeyDBUser):     viper.GetString(string(KeyDBUser)),
	}

	for _, v := range config {
		promptConfigAndExit(v == "")
	}
}

func InitDatabase(ctx context.Context) (database.Database, error) {
	db, err := database.New(ctx, database.Config{
		Host:     viper.GetString(string(KeyDBHost)),
		Port:     5432,
		User:     "lagra-test",
		Password: "mysecretpassword",
		DBName:   "lagra-test",
		SSLMode:  "disable",
	})
	if err != nil {
		return database.Database{}, fmt.Errorf("[main] Error connecting to database: %v", err)
	}
	return db, nil
}

func Compose(commands ...func(cmd *cobra.Command, args []string)) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		for _, command := range commands {
			command(cmd, args)
		}
	}
}

func PrintResult(ids ...id.ID) {
	fmt.Println("--------------------------------")
	for _, id := range ids {
		fmt.Println(id)
	}
	fmt.Println("--------------------------------")
}

func SetAccID(cmd *cobra.Command, args []string) error {
	accID, err := id.Parse(viper.GetString(string(KeyAccountID)))
	if err != nil {
		return fmt.Errorf("cannot parse %s from config: %v", string(KeyAccountID), err)
	}
	AccountID = accID
	return nil
}
