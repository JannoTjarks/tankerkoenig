package api

import (
	"fmt"
	"io"
	"net/http"
)

var apiDomain string = "https://creativecommons.tankerkoenig.de"

func getApiResponse(fullRequestUrl string) string {
	resp, err := http.Get(fullRequestUrl)
	if err != nil {
		//handle error
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	return string(body)
}

func RequestAreaSearch(apiKey string, lat float32, lng float32, rad float32,
	sort string, fuel string) string {
	var apiPath string = "/json/list.php"
	var apiQuery string = fmt.Sprintf("lat=%f&lng=%f&rad=%f&sort=%s&type=%s"+
		"&apikey=%s", lat, lng, rad, sort, fuel, apiKey)
	var fullRequestUrl = fmt.Sprintf("%s%s?%s", apiDomain, apiPath, apiQuery)

	return getApiResponse(fullRequestUrl)
}

func RequestPrice(apiKey string, stationIds []string) string {
	var apiPath string = "/json/prices.php"
	var stationIdsList string

	for _, element := range stationIds {
		if stationIdsList == "" {
			stationIdsList += element
		} else {
			stationIdsList += "," + element
		}
	}

	var apiQuery string = fmt.Sprintf("apikey=%s&ids=%s", apiKey, stationIdsList)
	var fullRequestUrl = fmt.Sprintf("%s%s?%s", apiDomain, apiPath, apiQuery)

	return getApiResponse(fullRequestUrl)
}

func RequestStationDetails(apiKey string, stationId string) string {
	var apiPath string = "/json/detail.php"
	var apiQuery string = fmt.Sprintf("apikey=%s&id=%s", apiKey, stationId)
	var fullRequestUrl = fmt.Sprintf("%s%s?%s", apiDomain, apiPath, apiQuery)

	return getApiResponse(fullRequestUrl)
}
