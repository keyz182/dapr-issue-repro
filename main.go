package main

import (
	"context"
	"reflect"

	dapr "github.com/dapr-sandbox/components-go-sdk"
	skdBinding "github.com/dapr-sandbox/components-go-sdk/bindings/v1"
	"github.com/dapr/components-contrib/bindings"
	contribMetadata "github.com/dapr/components-contrib/metadata"
	"github.com/dapr/kit/logger"
)

var log = logger.NewLogger("kafka-bindings-pluggable")

type LoggingBinding struct {
}

type LoggingMetadata struct {
	Version string `mapstructure:"version"`
}

func (b *LoggingBinding) Init(ctx context.Context, metadata bindings.Metadata) error {
	log.Infof("Init!")

	return nil
}

func (b *LoggingBinding) Operations() []bindings.OperationKind {
	return []bindings.OperationKind{
		"exec",
	}
}

func (b *LoggingBinding) Invoke(ctx context.Context, req *bindings.InvokeRequest) (*bindings.InvokeResponse, error) {
	log.Info(req.Operation)
	log.Info(string(req.Data[:]))
	return nil, nil
}

func (b *LoggingBinding) GetComponentMetadata() map[string]string {
	metadataStruct := LoggingMetadata{}
	metadataInfo := map[string]string{}
	_ = contribMetadata.GetMetadataInfoFromStructType(reflect.TypeOf(metadataStruct), &metadataInfo, contribMetadata.BindingType)
	return metadataInfo
}

func main() {
	dapr.Register("mylog",
		dapr.WithOutputBinding(func() skdBinding.OutputBinding {
			return &LoggingBinding{}
		}),
	)
	log.Info("Starting binding")
	dapr.MustRun()
}
