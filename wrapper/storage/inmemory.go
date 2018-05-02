package storage

import "sync"

type inMemoryStorage struct {
	currentCommand   *Command
	previousCommands []*Command
	mutex            *sync.Mutex
}

func handleBuffer(storage *inMemoryStorage, buffer <-chan []byte) {
	for {
		line := <-buffer
		storage.mutex.Lock()
		if storage.currentCommand != nil {
			storage.currentCommand.Output += string(line)
		}
		storage.mutex.Unlock()
	}
}

func (s *inMemoryStorage) StartListening(startTime int) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.currentCommand = &Command{"", "", -1, startTime, -1}
}

func (s *inMemoryStorage) StopListening(command string, returnCode int, endTime int) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.currentCommand == nil {
		return;
	}

	s.currentCommand.Command = command
	s.currentCommand.ReturnCode = returnCode
	s.currentCommand.EndTime = endTime
	s.previousCommands = append(s.previousCommands, s.currentCommand)
	s.currentCommand = nil
}

func (s *inMemoryStorage) List(count int) []*Command {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if count > len(s.previousCommands) {
		count = len(s.previousCommands)
	}

	return s.previousCommands[:count]
}
