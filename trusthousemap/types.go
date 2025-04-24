package trusthousemap

const (
	logInfo = 1
	logWarn = 2
	logErr  = 3
)

type PostcodeResponse struct {
	Result struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"result"`
}
