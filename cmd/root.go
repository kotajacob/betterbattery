/*
Copyright © 2020 Dakota Walsh

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
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const now string = "/sys/class/power_supply/BAT0/energy_now"
const max string = "/sys/class/power_supply/BAT0/energy_full"
const status string = "/sys/class/power_supply/BAT0/status"

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "betterbattery",
	Short: "A battery printing utility.",
	Long:  `betterbattery prints the battery percentage, status, and can run a command if the percentage fell below a specified value since it was last ran.`,
	Run: func(cmd *cobra.Command, args []string) {
		bb()
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
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.betterbattery.yaml)")
}

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
		viper.AddConfigPath(home)
		viper.SetConfigName(".betterbattery")
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func bb() {
	n := read(now)
	m := read(max)
	s := read(status)
	ni, err := strconv.Atoi(n)
	if err != nil {
		fmt.Fprintf(os.Stderr, "betterbattery: %v\n", err)
	}
	mi, err := strconv.Atoi(m)
	if err != nil {
		fmt.Fprintf(os.Stderr, "betterbattery: %v\n", err)
	}
	percent := int(float32(ni) / (float32(mi) / 100))
	fmt.Printf("%v", percent)
	fmt.Println(s)
}

func read(p string) string {
	v, err := ioutil.ReadFile(p)
	if err != nil {
		fmt.Fprintf(os.Stderr, "betterbattery: reading %s: %v\n", p, err)
		os.Exit(1)
	}
	return strings.TrimSuffix(string(v), "\n")
}
