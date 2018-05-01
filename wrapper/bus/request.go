package bus

type Request struct {
	Type string `json:"type"`
}

type StartListenRequest struct {
	Time string `json:"time"`
}

type StopListenRequest struct {
	Command    string `json:"command"`
	ReturnCode int    `json:"returnCode"`
	Time       string `json:"time"`
}

type GetRequest struct {
	Number int `json:"number"`
}
