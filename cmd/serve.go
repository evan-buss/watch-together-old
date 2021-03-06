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
	"net/http"
	"path/filepath"
	"time"

	"github.com/evan-buss/watch-together/server"
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the video server",
	Long: `Start the video server using the config specified in your ~/.watch-together.toml 
	Certain settings can be overridden using the appropriate flag variables`,
	Run: func(cmd *cobra.Command, args []string) {

		s := &server.Server{Router: chi.NewMux(), Hub: server.NewHub()}

		dbPath := filepath.Join(filepath.Dir(viper.ConfigFileUsed()), viper.GetString("database"))
		s.DB = sqlx.MustOpen("sqlite3", dbPath)

		// Start the local server

		go s.Hub.Run()
		s.Middlewares()
		s.Routes()

		// Make sure connections don't take too long
		server := &http.Server{
			Addr:         ":" + viper.GetString("port"),
			ReadTimeout:  time.Second * 10,
			WriteTimeout: time.Second * 10,
			Handler:      s.Router,
		}

		fmt.Println("Server listening on port:", viper.GetString("port"))
		log.Fatal(server.ListenAndServe())
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringP("port", "p", "8080", "Set listening port")
	err := viper.BindPFlag("port", serveCmd.Flags().Lookup("port"))
	if err != nil {
		log.Println(err)
	}
}
