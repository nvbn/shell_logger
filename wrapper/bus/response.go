package bus

type Response struct {
	Status string `json:"status"`
}

type GetResponse struct {
	Status     string `json:"status"`
	Command    string `json:"command"`
	ReturnCode int    `json:"returnCode"`
	StartTime  string `json:"time"`
	EndTime    string `json:"time"`
}
