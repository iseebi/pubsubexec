package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"cloud.google.com/go/pubsub"
)

func runCommand(command string, payload []byte) error {
	cmd := exec.Command(command)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	if _, err = stdin.Write(payload); err != nil {
		return err
	}
	if err = stdin.Close(); err != nil {
		return err
	}
	_, err = cmd.Output()
	return err
}

func watch(projectId string, subscriptionId string, command string) error {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectId)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %w", err)
	}
	defer client.Close()

	subscription := client.Subscription(subscriptionId)

	err = subscription.Receive(ctx, func(ctx context.Context, message *pubsub.Message) {
		log.Printf("Got message: %q\n", string(message.Data))
		message.Ack()
		err := runCommand(command, message.Data)
		if err != nil {
			_ = fmt.Errorf("execute error: %w", err)
			return
		}
	})
	if err != nil {
		return fmt.Errorf("receive error: %w", err)
	}
	return nil
}

func main() {
	project := flag.String("project", "", "Google Cloud Project ID")
	subscription := flag.String("subscription", "", "Google Cloud Pub/Sub Subscription ID")
	command := flag.String("command", "", "execute command")
	flag.Parse()

	if *project == "" {
		log.Fatalf("Must provide Google Cloud Project ID by --project option")
	}
	if *subscription == "" {
		log.Fatalf("Must provide Google Cloud Pub/Sub Subscription ID by --subscription option")
	}
	if *command == "" {
		log.Fatalf("Must provide command by --command option")
	}
	err := watch(*project, *subscription, *command)
	if err != nil {
		os.Exit(1)
	}
}