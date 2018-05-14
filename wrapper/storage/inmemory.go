package storage

import "sync"

type inMemoryStorage struct {
	currentCommand   *Command
	previousCommands []*Command
	mutex            *sync.Mutex
}

// Creates new in-memory storage instance.
func NewInMemory(buffer <-chan []byte) *inMemoryStorage {
	storage := &inMemoryStorage{nil, []*Command{}, &sync.Mutex{}}
	go storage.handleBuffer(buffer)
	return storage
}

func (s *inMemoryStorage) handleBuffer(buffer <-chan []byte) {
	for {
		line := <-buffer
		s.mutex.Lock()
		if s.currentCommand != nil {
			s.currentCommand.Output += string(line)
		}
		s.mutex.Unlock()
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
		return
	}

	s.currentCommand.Command = command
	s.currentCommand.ReturnCode = returnCode
	s.currentCommand.EndTime = endTime
	s.previousCommands = append([]*Command{s.currentCommand}, s.previousCommands...)
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
