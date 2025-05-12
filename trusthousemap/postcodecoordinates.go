package trusthousemap

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Bevs-n-Devs/trusthouse/logs"
)

func PostcodeCoordinates(postcode string) (float64, float64, error) {
	url := fmt.Sprintf("https://api.postcodes.io/postcodes/%s", postcode)

	response, err := http.Get(url)
	if err != nil {
		logs.Logs(logErr, "failed to get postcode coordinates: "+err.Error())
		return 0, 0, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		logs.Logs(logErr, "API to get postcode coordinates: "+response.Status)
		return 0, 0, fmt.Errorf("failed to get postcode coordinates: %s", response.Status)
	}

	var data PostcodeResponse
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		logs.Logs(logErr, "failed to decode postcode coordinates: "+err.Error())
		return 0, 0, err
	}

	return data.Result.Latitude, data.Result.Longitude, nil

}
