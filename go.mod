module core-data

go 1.13

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/advwacloud/datahub-edge-domain-models v0.0.0-20201015072722-0e67c490edc9 // indirect
	github.com/edgexfoundry/core-data-go v0.0.0-20180105155421-58619b9a0201
	github.com/edgexfoundry/core-domain-go v0.0.0-20180130224812-7acdb6490aba // indirect
	github.com/golang/protobuf v1.4.2
	github.com/micro/go-micro/v2 v2.9.1
	google.golang.org/protobuf v1.25.0
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
)
