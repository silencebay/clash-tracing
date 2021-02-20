package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func envOrDefault(env string, def string) string {
	value, exist := os.LookupEnv(env)
	if !exist {
		return def
	}
	return value
}

func main() {
	serverurl := envOrDefault("INFLUXDB_SERVERURL", "http://localhost:8086")
	token := envOrDefault("INFLUXDB_TOKEN", "")
	org := envOrDefault("INFLUXDB_ORG", "clash")
	bucket := envOrDefault("INFLUXDB_BUCKET", "events")

	clashHost := envOrDefault("CLASH_HOST", "localhost:9090")
	clashToken := envOrDefault("CLASH_TOKEN", "")
	client := influxdb2.NewClient(serverurl, token)
	defer client.Close()

	startQueue(client.WriteAPI(org, bucket))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go handleTraffic(ctx, clashHost, clashToken)
	go handleTracing(ctx, clashHost, clashToken)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
}
