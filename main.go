package main

import "github.com/Bevs-n-Devs/trusthouse/logs"

const (
	logInfo  = 1
	logDbErr = 5
)

func main() {
	go logs.LogProcessor()
	logs.Logs(logInfo, "Starting Trust House application...")

	logs.Logs(logDbErr, "hello world, hello Yaw!")

	// select {} // Keep the main function running to allow log processing
}
