package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Bevs-n-Devs/trusthouse/logs"
	"github.com/Bevs-n-Devs/trusthouse/tmpl"
)

func StartServer() {
	logs.Logs(logInfo, "Starting HTTP Server")

	tmpl.InitTemplates()

	var staticFiles = http.FileServer((http.Dir("./static/")))
	http.Handle("/static/", http.StripPrefix("/static/", staticFiles))

	// define routes
	http.HandleFunc("/", Home)

	// initialise port for application
	httpPort := os.Getenv("PORT")

	// start server on machine if hosting platform port from hosting platform

	if httpPort == "" {

		logs.Logs(logWarn, fmt.Sprintf("Could not get PORT from hosting platform. Defaulting to http://localhost:%s...", localPort))
		httpPort = localPort
		err := http.ListenAndServe("localhost:"+localPort, nil)
		if err != nil {
			logs.Logs(logErr, fmt.Sprintf("Failed to start HTTP server: %s", err.Error()))
		}
	}

	// start server on hosting platform port
	logs.Logs(logInfo, fmt.Sprintf("HTTP server running on http://localhost:%s", httpPort))
	err := http.ListenAndServe(":"+httpPort, nil)
	if err != nil {
		logs.Logs(logErr, fmt.Sprintf("Error starting HTTP server: %s", err.Error()))
	}
}
