package configcmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/d-darac/lagra/cmd/lagra-mock/cmd/shared"
	"github.com/d-darac/lagra/pkg/id"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ConfigCmd represents the config command
var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Use to create a config for lagra-mock",
	Run:   config,
}

func init() {
}

func config(cmd *cobra.Command, args []string) {
	vals := []string{
		string(shared.KeyDBHost),
		string(shared.KeyDBPort),
		string(shared.KeyDBName),
		string(shared.KeyDBUser),
		string(shared.KeyDBPassword),
		string(shared.KeyAccountID),
	}
	reader := bufio.NewReader(os.Stdin)

	for i, prompt := range vals {
		prompt = strings.ReplaceAll(prompt, "_", " ")
		fmt.Print(prompt + ": ")
		s, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if prompt == strings.ReplaceAll(string(shared.KeyDBPort), "_", " ") {
			for true {
				_, e := strconv.Atoi(strings.TrimSuffix(s, "\n"))
				if e == nil {
					break
				}
				fmt.Printf("%s must be an integer\n", prompt)
				fmt.Print(prompt + ": ")
				s, err = reader.ReadString('\n')
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					os.Exit(1)
				}
			}
		}

		if prompt == strings.ReplaceAll(string(shared.KeyAccountID), "_", " ") {
			for true {
				_, e := id.Parse(strings.TrimSuffix(s, "\n"))
				if e == nil {
					break
				}
				fmt.Printf("%s must be a valid ID\n", prompt)
				fmt.Print(prompt + ": ")
				s, err = reader.ReadString('\n')
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					os.Exit(1)
				}
			}
		}

		vals[i] = strings.TrimSuffix(s, "\n")
	}

	viper.Set(string(shared.KeyDBHost), vals[0])
	viper.Set(string(shared.KeyDBPort), vals[1])
	viper.Set(string(shared.KeyDBName), vals[2])
	viper.Set(string(shared.KeyDBUser), vals[3])
	viper.Set(string(shared.KeyDBPassword), vals[4])
	viper.Set(string(shared.KeyAccountID), vals[5])

	if err := viper.WriteConfig(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
