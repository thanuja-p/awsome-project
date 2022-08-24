package main

import (
	"bitbucket.org/ntuclink/devex-sdk-go/logging/v2"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

const (
	amplitudeAPIUrl     = "https://api2.amplitude.com/2/httpapi"
	amplitudeBulkAPIUrl = "https://api2.amplitude.com/batch"
)

func SendEvent(ctx context.Context, request *AmplitudeRequest) {
	headers := map[string][]string{
		"Content-Type": []string{"application/json"},
		"Accept":       []string{"*/*"},
	}
	requestJson, marshalError := json.Marshal(*request)
	if marshalError != nil {
		logging.WithCtx(ctx).Errorf("Unable to marshal the Amplitude request : %v", marshalError)
	} else {
		data := bytes.NewBuffer(requestJson)
		//req, err := http.NewRequest(http.MethodPost, os.Getenv("AMPLITUDE_URL"), data)
		req, err := http.NewRequest(http.MethodPost, amplitudeAPIUrl, data)
		req.Header = headers

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			logging.WithCtx(ctx).Errorf("Failed to send request to Amplitude: %v", err)
		} else {
			logging.WithCtx(ctx).Debugf("Successfully published the amplitude request : %v", resp)
		}
	}
}

func SendBatchEvent(ctx context.Context, request *AmplitudeRequest) {
	headers := map[string][]string{
		"Content-Type": []string{"application/json"},
		"Accept":       []string{"*/*"},
	}
	requestJson, marshalError := json.Marshal(*request)
	if marshalError != nil {
		logging.WithCtx(ctx).Errorf("Unable to marshal the Amplitude batch request : %v", marshalError)
	} else {
		data := bytes.NewBuffer(requestJson)
		//req, err := http.NewRequest(http.MethodPost, os.Getenv("AMPLITUDE_BATCH_URL"), data)
		req, err := http.NewRequest(http.MethodPost, amplitudeBulkAPIUrl, data)
		req.Header = headers

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			logging.WithCtx(ctx).Errorf("Failed to send request to Amplitude Batch Request: %v", err)
		} else {
			logging.WithCtx(ctx).Debugf("Successfully published sent amplitude batch request : %v", resp)
		}
	}
}
