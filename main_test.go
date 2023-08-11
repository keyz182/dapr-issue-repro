package main

import (
	"context"
	"fmt"
	logger "log"
	"net/http"
	"os"
	"testing"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/gorilla/mux"
)

/*
dapr run --app-id batch-sdk --app-port 6004 -G 6005 --resources-path components --log-level debug -- go test
*/

var (
	cronBindingName, bindingName string = "cron", "mylog"
)

func receiveData(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Cron Trigger...")

	client, err := dapr.NewClient()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	ctx := context.Background()

	// Insert order using Dapr output binding via Dapr SDK
	in := &dapr.InvokeBindingRequest{
		Name:      bindingName,
		Operation: "exec",
		Data:      []byte(""),
		Metadata:  map[string]string{},
	}
	err = client.InvokeOutputBinding(ctx, in)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func TestBinding(t *testing.T) {
	log.Info("Port")
	log.Info(os.Getenv("DAPR_GRPC_PORT"))
	var appPort string
	var okHost bool
	if appPort, okHost = os.LookupEnv("APP_PORT"); !okHost {
		appPort = "6004"
	}

	r := mux.NewRouter()

	// Triggered by Dapr input binding
	r.HandleFunc("/"+cronBindingName, receiveData).Methods("POST")

	if err := http.ListenAndServe(":"+appPort, r); err != nil {
		logger.Panic(err)
	}
}
