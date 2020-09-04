package main

type ConfigurationStruct struct {
	Applicationname            string
	Consulprofilesactive       string
	Readmaxlimit               int
	Metadatacheck              bool
	Addtoeventqueue            bool
	Persistdata                bool
	Heartbeattime              int
	Heartbeatmsg               string
	Appopenmsg                 string
	Formatspecifier            string
	Msgpubtype                 string
	Serverport                 int
	Serviceaddress             string
	Servicename                string
	Deviceupdatelastconnected  bool
	Serviceupdatelastconnected bool
	MongoDBUserName            string
	MongoDBPassword            string
	MongoDatabaseName          string
	MongoDBHost                string
	MongoDBPort                int
	MongoDBConnectTimeout      int
	DatamongodbsocketKeepAlive bool
	Consulhost                 string
	Consulcheckaddress         string
	Consulport                 int
	Checkinterval              string
	Loggingfile                string
	Loggingremoteurl           string
	Metadbaddressableurl       string
	Metadbdeviceserviceurl     string
	Metadbdeviceprofileurl     string
	Metadbdeviceurl            string
	Metadbdevicereporturl      string
	Metadbcommandurl           string
	Metadbeventurl             string
	Metadbscheduleurl          string
	Metadbprovisionwatcherurl  string
	Metadbpingurl              string
	Activemqbroker             string
	Zeromqaddressport          string
	Amqbroker                  string
}

var configuration ConfigurationStruct = ConfigurationStruct{} //  Needs to be initialized before used
