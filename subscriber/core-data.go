package subscriber

import (
	"context"
	"time"

	log "github.com/micro/go-micro/v2/logger"

	"core-data/clients"
	coredata "core-data/proto/core-data"

	models "github.com/advwacloud/datahub-edge-domain-models"
)

type CoreData struct{}

func (e *CoreData) Handle(ctx context.Context, msg *coredata.Message) error {
	log.Info("Handler Received message: ", msg.SourceId)

	ts, _ := time.Parse(time.RFC3339, msg.Time)
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
		clients.Dbc.AddData(data)
	}
	return nil
}

func Handler(ctx context.Context, msg *coredata.Message) error {
	log.Info("Function Received message: ", msg.SourceId)

	return nil
}
