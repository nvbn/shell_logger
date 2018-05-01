package storage

// Logged entry for previously executed shell command.
type Command struct {
	Command    string `json:"command"`
	Output     string `json:"output"`
	ReturnCode int    `json:"returnCode"`
	StartTime  int    `json:"time"`
	EndTime    int    `json:"time"`
}

// Storage for history of commands.
type Storage interface {
	// Start listening to shell logger.
	StartListening(startTime int)

	// Stop listening to shell logger.
	StopListening(command string, returnCode int, endTime int)

	// List logged commands:
	List(count int) []*Command
}
