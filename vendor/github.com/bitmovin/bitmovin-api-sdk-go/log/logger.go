package log

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestLogInfo struct {
	URL     string
	Method  string
	RawBody []byte
	Headers http.Header
}

type ResponseLogInfo struct {
	URL        string
	Method     string
	RawBody    []byte
	Headers    http.Header
	StatusCode int
}

type Logger interface {
	LogRequest(reqInfo RequestLogInfo)
	LogResponse(respInfo ResponseLogInfo)
}

type StdOutLogger struct {
}

func (c *StdOutLogger) LogRequest(reqInfo RequestLogInfo) {
	headerStr, _ := json.Marshal(reqInfo.Headers)
	logStr := fmt.Sprintf("\n#### REQUEST ####\n"+
		"Url: %s - %s\n"+
		"Headers: %s\n",
		reqInfo.Method, reqInfo.URL, headerStr)

	if len(reqInfo.RawBody) > 0 {
		logStr += fmt.Sprintf("Payload: %s\n", reqInfo.RawBody)
	}

	logStr += "#### REQUEST END ####\n"
	fmt.Print(logStr)

}

func (c *StdOutLogger) LogResponse(respInfo ResponseLogInfo) {
	headerStr, _ := json.Marshal(respInfo.Headers)
	logStr := fmt.Sprintf("\n#### RESPONSE ####\n"+
		"Url: %s - %s\n"+
		"Status: %d\n"+
		"Headers: %s\n",
		respInfo.Method, respInfo.URL, respInfo.StatusCode, headerStr)

	if len(respInfo.RawBody) > 0 {
		logStr += fmt.Sprintf("Payload: %s\n", respInfo.RawBody)
	}

	logStr += "#### RESPONSE END ####\n"
	fmt.Print(logStr)
}
