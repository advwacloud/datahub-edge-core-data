package handler

import (
	"context"

	"github.com/golang/protobuf/ptypes"

	log "github.com/micro/go-micro/v2/logger"

	"core-data/clients"
	//coredata "core-data/proto/core-data"
	core "github.com/advwacloud/datahub-edge-domain-models/protos/core-data"

	models "github.com/advwacloud/datahub-edge-domain-models/models"
)

type CoreData struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *CoreData) SendRawData(ctx context.Context, req *core.SendDataRequest, rsp *core.SendDataReply) error {
	log.Info("Received CoreData.SendRawData request")
	//var datas []models.Data
	//ts, _ := time.Parse(time.RFC3339, req.GetTime())
	ts, _ := ptypes.Timestamp(req.GetTime())
	sid := req.GetSourceId()

	for _, tag := range req.GetData() {
		var data models.Data

		data.Created = ts
		data.SourceId = sid
		data.Name = tag.GetTagName()
		number := tag.GetNumber()
		text := tag.GetText()

		if text != "" {
			data.Value = text
		} else {
			data.Value = number
		}
		clients.Dbc.AddData(data) // write each data
		//datas = append(datas, data)
	}
	//clients.Dbc.AddData(datas)
	//id, err := clients.Dbc.AddData(d)
	return nil
}
