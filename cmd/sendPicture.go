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
	"fmt"
	"github.com/DarkReduX/gRPC_client/protocol"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// sendPictureCmd represents the sendPicture command
var sendPictureCmd = &cobra.Command{
	Use:   "sendPicture",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sendPicture called")
		conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure(), grpc.WithBlock())

		if err != nil {
			logrus.Fatalf("did not connect: %v", err)
		}

		defer conn.Close()

		client := protocol.NewHelloServiceClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*80)
		defer cancel()
		pic, err := os.Open("./pic.png")

		if err != nil {
			logrus.Fatalf("error: %v ", err)
		}
		//for {
		//
		//}
		bytes := make([]byte, 32*1024)
		stream, err := client.SendPicture(ctx)
		if err != nil {
			logrus.Fatalf("couldn't gree: %v", err)
		}
		for {
			_, err := pic.Read(bytes)
			if err == io.EOF {
				log.Println("read  pic")
				break
			}
			if err != nil {
				log.Fatalf("error read message")
			}
			if err := stream.Send(&protocol.Picture{Content: bytes}); err != nil {
				log.Fatal("error while sending pic")
			}
		}
		if err := stream.CloseSend(); err != nil {
			log.Fatal("failed closed string")
		}
		if err := pic.Close(); err != nil {
			log.Fatalf("error closing pic")
		}
		log.Printf("Received message: %v")
	},
}

func init() {
	rootCmd.AddCommand(sendPictureCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sendPictureCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sendPictureCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	sendPictureCmd.Flags().StringP("path", "p", "pic.img", "Path to picture")
}
