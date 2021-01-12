/*
Copyright © 2020 Sam Pointer <sam@outsidethe.net>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"

	"github.com/sampointer/digaws/command"
	"github.com/sampointer/digaws/manifest"
)

var cfgFile string
var format string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "digaws ipv4_address|ipv6_address ...",
	Short: "look up AWS IP address details",
	Long: `Intelligently parses the current AWS ip-ranges.json to enable you to
look up details of any specific IP address.

See https://docs.aws.amazon.com/general/latest/gr/aws-ip-ranges.html for more
information.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			res, err := command.Lookup(arg, manifest.GetManifest())
			if err != nil {
				fmt.Println(err)
				return
			}

			for _, p := range res {
				switch format {
				case "text":
					fmt.Println(p)
				case "json":
					out, err := p.JSON()
					if err != nil {
						fmt.Println(err)
						os.Exit(3)
					}
					fmt.Println(out)
				default:
					fmt.Println("invalid format")
					os.Exit(2)
				}
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.digaws.yaml)")
	rootCmd.Flags().StringVarP(&format, "format", "f", "text", "one of: text|json")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".digaws" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".digaws")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
