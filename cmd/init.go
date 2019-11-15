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
	"log"

	"github.com/evan-buss/watch-together/video/metadata"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCmd represents the serve command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize your library",
	Long: `Initialize takes care of several different library management tasks. 

	The video directory provided within your .watch-together.toml file will be scanned for movies. 
	Watch Together will save video metadata and file locations to a local cache.`,
	Run: func(cmd *cobra.Command, args []string) {
		metadata.ParseDir(viper.GetString("video-dir"))
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringP("video-dir", "v", viper.GetString("video-dir"), "Set the video directory to scan for movies.")
	err := viper.BindPFlag("port", serveCmd.Flags().Lookup("port"))
	if err != nil {
		log.Println(err)
	}
}
