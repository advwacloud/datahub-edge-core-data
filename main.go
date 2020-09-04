package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	//"strings"
	"time"

	"datahub-edge-core/clients"
	//"datahub-edge-core/pkg/metadataclients" // temporary

	logger "github.com/advwacloud/datahub-edge-common/log"
)

// Global variables
var dbc clients.DBClient
var loggingClient logger.LoggingClient

//var mdc metadataclients.DeviceClient
//var msc metadataclients.ServiceClient

// Heartbeat for the data microservice - send a message to logging service
func heartbeat() {
	// Loop forever
	for true {
		loggingClient.Info(configuration.Heartbeatmsg, "")
		time.Sleep(time.Millisecond * time.Duration(configuration.Heartbeattime)) // Sleep based on configuration
	}
}

// Read the configuration file and update configuration struct
func readConfigurationFile(path string) error {
	// Read the configuration file
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading configuration file: " + err.Error())
		return err
	}

	// Decode the configuration as JSON
	err = json.Unmarshal(contents, &configuration)
	if err != nil {
		fmt.Println("Error reading configuration file: " + err.Error())
		return err
	}

	return nil
}

func main() {
	start := time.Now()

	// Load configuration data
	err := readConfigurationFile("./res/configuration.json")
	if err != nil {
		fmt.Printf("Could not read configuration file(%s): %#v\n", "./res/configuration.json", err)
		os.Exit(1)
	}

	// Create Logger (Default Parameters)
	loggingClient = logger.NewClient(configuration.Servicename, configuration.Loggingremoteurl)
	loggingClient.LogFilePath = configuration.Loggingfile

	// Create a database client
	dbc, err = clients.NewDBClient(clients.DBConfiguration{
		DbType:       clients.MONGO,
		Host:         configuration.MongoDBHost,
		Port:         configuration.MongoDBPort,
		Timeout:      configuration.MongoDBConnectTimeout,
		DatabaseName: configuration.MongoDatabaseName,
		Username:     configuration.MongoDBUserName,
		Password:     configuration.MongoDBPassword,
	})
	if err != nil {
		loggingClient.Error("Couldn't connect to database: "+err.Error(), "")
		return
	}

	fmt.Println(dbc)

	// Create metadata clients
	//mdc = metadataclients.NewDeviceClient(configuration.Metadbdeviceurl)
	//msc = metadataclients.NewServiceClient(configuration.Metadbdeviceserviceurl)

	// Start heartbeat
	go heartbeat()

	r := loadRestRoutes()
	http.TimeoutHandler(nil, time.Millisecond*time.Duration(5000), "Request timed out")
	loggingClient.Info(configuration.Appopenmsg, "")

	// Time it took to start service
	loggingClient.Info("Service started in: "+time.Since(start).String(), "")
	loggingClient.Info("Listening on port: " + strconv.Itoa(configuration.Serverport))

	loggingClient.Error(http.ListenAndServe(":"+strconv.Itoa(configuration.Serverport), r).Error())
}
