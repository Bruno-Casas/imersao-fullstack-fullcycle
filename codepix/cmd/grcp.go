/*
Copyright © 2021 Bruno Casas <brunocasas04@gmail.com>

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
	"os"

	"github.com/codeedu/imersao/codepix-go/application/grpc"
	"github.com/codeedu/imersao/codepix-go/infrastructure/db"
	"github.com/spf13/cobra"
)

var portNumber int

// grcpCmd represents the grcp command
var grcpCmd = &cobra.Command{
	Use:   "grcp",
	Short: "Start gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		database := db.ConnectDB(os.Getenv("env"))
		grpc.StartGrpcServer(database, portNumber)
	},
}

func init() {
	rootCmd.AddCommand(grcpCmd)
	grcpCmd.Flags().IntVarP(&portNumber, "port", "p", 50051, "gRPC server port")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// grcpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// grcpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
