package util

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/fyllekanin/gnsagent/internal/schema"
)

func GetIpFromEndPoints(endpoints []schema.ConfigEndPoint) (string, error) {
	for _, item := range endpoints {
		result, err := getPublicIp(item)
		if err == nil {
			return result, nil
		}
	}
	return "", errors.New("failed to get any public ip")
}

func getPublicIp(endpoint schema.ConfigEndPoint) (string, error) {
	req, err := http.Get(endpoint.Url)
	if err != nil {
		return "", errors.New("failed requesting for URL: " + endpoint.Url)
	}
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return "", errors.New("failed reading body for URL: " + endpoint.Url)
	}

	var result map[string]string
	json.Unmarshal(body, &result)

	return result[endpoint.Property], nil
}
