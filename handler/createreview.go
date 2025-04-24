package handler

import (
	"fmt"
	"net/http"

	"github.com/Bevs-n-Devs/trusthouse/logs"
	"github.com/Bevs-n-Devs/trusthouse/trusthousemap"
	"github.com/Bevs-n-Devs/trusthouse/utils"
)

func CreateReview(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		logs.Logs(logErr, "Invalid request method")
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		logs.Logs(logErr, "Failed to parse form: "+err.Error())
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	propertyNumber := r.FormValue("propertyNumber")
	streetName := r.FormValue("streetName")
	townOrCity := r.FormValue("townOrCity")
	postcode := r.FormValue("postcode")
	propertyRating := r.FormValue("propertyRating")
	propertyReview := r.FormValue("propertyReview")
	reviewedBy := r.FormValue("reviewedBy")

	cleanPostcode := utils.CleanPostcode(postcode)

	longitude, latitude, err := trusthousemap.PostcodeCoordinates(cleanPostcode)
	if err != nil {
		logs.Logs(logErr, "Failed to get coordinates: "+err.Error())
		http.Error(w, "Failed to get coordinates", http.StatusBadRequest)
		return
	}

	logs.Logs(logInfo, "Creating review for property: "+propertyNumber+" "+streetName+", "+townOrCity+", "+postcode+" by "+reviewedBy)
	logs.Logs(logInfo, "Property rating: "+propertyRating)
	logs.Logs(logInfo, "Property review: "+propertyReview)
	logs.Logs(logInfo, fmt.Sprintf("Coordinates: %f, %f", latitude, longitude))
}
