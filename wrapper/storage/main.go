package storage

import "sync"

type Command struct {
	Command    string `json:"command"`
	Output     string `json:"output"`
	ReturnCode int    `json:"returnCode"`
	StartTime  int    `json:"time"`
	EndTime    int    `json:"time"`
}

type Storage struct {
	currentCommand   *Command
	previousCommands []*Command
	mutex            *sync.Mutex
}

func handleBuffer(storage *Storage, buffer <-chan []byte) {
	for {
		line := <-buffer
		storage.mutex.Lock()
		if storage.currentCommand != nil {
			storage.currentCommand.Output += string(line)
		}
		storage.mutex.Unlock()
	}
}

func New(buffer <-chan []byte) *Storage {
	storage := &Storage{nil, nil, &sync.Mutex{}}
	go handleBuffer(storage, buffer)
	return storage
}

func (s *Storage) StartListening(time int) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.currentCommand = &Command{"", "", -1, time, -1}
}

func (s *Storage) StopListening(command string, returnCode int, time int) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.currentCommand.Command = command
	s.currentCommand.ReturnCode = returnCode
	s.currentCommand.EndTime = time
	s.previousCommands = append(s.previousCommands, s.currentCommand)
	s.currentCommand = nil
}

func (s *Storage) List(count int) []*Command {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if count > len(s.previousCommands) {
		count = len(s.previousCommands)
	}

	return s.previousCommands[:count]
}
