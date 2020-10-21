package subscriber

import (
	"context"
	"fmt"

	//"time"

	"github.com/golang/protobuf/ptypes"

	"core-data/clients"

	core "github.com/advwacloud/datahub-edge-domain-models/protos/core-data"

	models "github.com/advwacloud/datahub-edge-domain-models/models"
	log "github.com/micro/go-micro/v2/logger"
)

type CoreData struct{}

func (e *CoreData) Handle(ctx context.Context, msg *core.Message) error {
	log.Info("Handler Received message: ", msg.SourceId)

	ts, _ := ptypes.Timestamp(msg.Time)
	sid := msg.SourceId

	for _, tag := range msg.Data {
		var data models.Data

		data.Created = ts
		data.SourceId = sid
		data.Name = tag.TagName

		number := tag.GetNumber()
		text := tag.GetText()

		if text != "" {
			data.Value = text
		} else {
			data.Value = number
		}

		// if data.Created.IsZero() {
		// 	fmt.Printf("No date has been set, %s\n", data.Created)
		// 	data.Created = time.Now()
		// }
		fmt.Println(data)
		clients.Dbc.AddData(data)
	}
	return nil
}

func Handler(ctx context.Context, msg *core.Message) error {
	log.Info("Function Received message: ", msg.SourceId)

	return nil
}
