package main

import (
        "context"
        "log"
        "os"
        "os/signal"
        "syscall"
        "time"

        "cloud.google.com/go/pubsub"
        "gopkg.in/yaml.v2"
)

// Config struct for storing configuration parameters
type Config struct {
        ProjectID              string `yaml:"projectID"`
        TopicName              string `yaml:"topicName"`
        SubscriptionName       string `yaml:"subscriptionName"`
        SleepIntervalSecond    int    `yaml:"sleepIntervalSeconds"`
        MaxOutstandingMessages int    `yaml:"maxOutstandingMessages"`
}

func loadConfig(filename string) (Config, error) {
        configFile, err := os.Open(filename)
        if err != nil {
                return Config{}, err
        }
        defer configFile.Close()

        var config Config
        err = yaml.NewDecoder(configFile).Decode(&config)
        if err != nil {
                return Config{}, err
        }
        return config, nil
}

func main() {
        configFile := "config.yaml"

        config, err := loadConfig(configFile)
        if err != nil {
                log.Fatalf("Failed to load config: %v", err)
        }

        ctx, cancel := context.WithCancel(context.Background())
        defer cancel()

        client, err := pubsub.NewClient(ctx, config.ProjectID)
        if err != nil {
                log.Fatalf("Failed to create client: %v", err)
        }
        defer client.Close()

        sub := client.Subscription(config.SubscriptionName)
        exists, err := sub.Exists(ctx)
        if err != nil {
                log.Fatalf("Error checking subscription existence: %v", err)
        }
        if !exists {
                log.Fatalf("Subscription %s does not exist", config.SubscriptionName)
        }
        sub.ReceiveSettings.Synchronous = true
        sub.ReceiveSettings.MaxOutstandingMessages = config.MaxOutstandingMessages

        sigCh := make(chan os.Signal, 1)
        signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

        go func() {
                sig := <-sigCh
                log.Printf("Received signal: %v\n", sig)
                cancel()
        }()

        for {
                select {
                case <-ctx.Done():
                        log.Println("Subscription closed, exiting...")
                        return
                default:
                        receiveMessages(ctx, sub, config)
                }
        }
}

func receiveMessages(ctx context.Context, sub *pubsub.Subscription, config Config) {
        err := sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
                log.Printf("topic %s: %s\n", config.TopicName, string(msg.Data))
                msg.Ack()
                time.Sleep(time.Duration(config.SleepIntervalSecond) * time.Second)
        })
        if err != nil {
                log.Fatalf("Error receiving messages from topic %s: %v", config.TopicName, err)
        }
}
