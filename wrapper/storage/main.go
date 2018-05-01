package storage

import "sync"

type Command struct {
	Command string
	Output string
	ReturnCode int
	StartTime string
	EndTime string
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

func (s *Storage) StartListen(time string) {
	s.mutex.Lock()
	s.currentCommand = &Command{"", "", -1, time, ""}
	s.mutex.Unlock()
}

func (s *Storage) StopListen(command string, returnCode int, time string) {
	s.mutex.Lock()
	s.currentCommand.Command = command
	s.currentCommand.ReturnCode = returnCode
	s.currentCommand.EndTime = time
	s.previousCommands = append(s.previousCommands, s.currentCommand)
	s.currentCommand = nil
	s.mutex.Unlock()
}

func (s *Storage) Get(number int) *Command {
	return s.previousCommands[number]
}
