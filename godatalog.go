package godatalog

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type LogConfig struct {
	// LogConfig is used to set the datadog parameters used for sending log messages.
	URL string
	Port int
	UseSSL bool
	UseTCP bool
	APIKey string
	Tags string
}

type LogEntry struct {
	// LogEntry is used to set the log messages.
	Message string `json:"message"`
	ServiceName string `json:"service"`
	Source string `json:"source"`
	HostName string `json:"hostname"`
	Level string `json:"level"`
	Logger string `json:"logger"`
	AppName string `json:"appname"`
	Tags string `json:"ddtags"` // comma separated with no spaces
}

func CreateLogEntry(message string, serviceName string, source string, hostName string, level string, logger string, appName string, tags string) *LogEntry {
	return &LogEntry{Message: message, ServiceName: serviceName, Source: source, HostName: hostName, Level: level, Logger: logger, AppName: appName, Tags: tags}
}

func CreateLogConfig(URL string, port int, useSSL bool, useTCP bool, APIKey string, tags string) *LogConfig {
	return &LogConfig{URL: URL, Port: port, UseSSL: useSSL, UseTCP: useTCP, APIKey: APIKey, Tags: tags}
}

func DebugLog(err error, DDC *LogConfig){

}
func InfoLog(err error, DDC *LogConfig){

}
func WarnLog(err error, DDC *LogConfig){

}
func ErrLog(err error, DDC *LogConfig){

}

func PushLog(message *LogEntry, DDC *LogConfig)(*http.Response, error){
	// PushLog
	p, err := json.Marshal(message)
	if err != nil {fmt.Println(err)}

	client := &http.Client {}
	req, err := http.NewRequest("POST", DDC.URL+DDC.APIKey, bytes.NewReader(p))
	if err != nil {fmt.Println(err)}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	return res, err
}