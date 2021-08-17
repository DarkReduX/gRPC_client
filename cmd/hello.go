/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"context"
	"github.com/DarkReduX/gRPC_client/protocol"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
	"time"
)

// helloCmd represents the hello command
var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Hello",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure(), grpc.WithBlock())

		if err != nil {
			logrus.Fatalf("did not connect: %v", err)
		}

		defer conn.Close()

		client := protocol.NewHelloServiceClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*80)
		defer cancel()
		name := cmd.Flag("name").Value.String()
		res, err := client.SayHello(ctx, &protocol.UserNameMessage{Name: name})
		if err != nil {
			logrus.Fatalf("couldn't gree: %v", err)
		}
		log.Printf("Received message: %v", res.GetMessage())
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helloCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	helloCmd.Flags().StringP("name", "n", "world", "Help message for toggle")
}
