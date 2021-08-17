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

	helloCmd.Flags().StringP("name", "n", "world", "Help message for toggle")
}
