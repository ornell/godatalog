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
}

type LogEntry struct {
	// LogEntry is used to set the log messages.
	Message string `json:"message"`
	ServiceName string
	Source string `json:"source"`
	HostName string `json:"hostname"`
	Level string `json:"level"`
	Logger string `json:"service"`
	AppName string `json:"appname"`
	Tags string `json:"ddtags"` // comma separated with no spaces
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