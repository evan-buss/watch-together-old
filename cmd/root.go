/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

//var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "watch-together",
	Short: "A self-hosted way to watch movies with friends",
	Long: `Watch Together is a way to stream movies from your 
  computer in real time with your friends. Simply run the server on your computer,
  go to the website and send your friends the IP to connect to. 
  Watch videos and chat in real-time.`,
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
	// Define future flags and configuration settings here
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Search config in home directory with name ".watch-together.toml"
	// viper.AddConfigPath()
	// viper.SetConfigName("watch-together")
	// viper.SetConfigType("toml")
	viper.SetConfigFile(filepath.Join(home, ".config/watch-together/watch-together.toml"))

	// Default values
	viper.SetDefault("video-dir", filepath.Join(home, "Videos"))
	viper.SetDefault("database", "library.db")
	viper.SetDefault("port", "8080")

	//viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Creating Config File Directory: ~/.config/watch-together/")
		path := filepath.Join(home, ".config/watch-together")
		// fmt.Println("Creating Config file:", path)
		err := os.MkdirAll(path, 0700)
		if err != nil {
			log.Println(err)
		}

		// Write the config file if it doesn't exist
		err = viper.WriteConfigAs(filepath.Join(path, "watch-together.toml"))
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Configuration Loaded")
	}
}
