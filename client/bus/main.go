package bus

// Start listening to shell logger.
func StartListening(socketPath string, time int) error {
	request := startListeningRequest(time)

	body, err := communicate(socketPath, request)
	if err != nil {
		return err
	}

	if !isOk(body) {
		return getError(body)
	}

	return nil
}

// Stop listening to shell logger.
func StopListening(socketPath string, command string, returnCode int, time int) error {
	request := stopListeningRequest(command, returnCode, time)

	body, err := communicate(socketPath, request)
	if err != nil {
		return err
	}

	if !isOk(body) {
		return getError(body)
	}

	return nil
}
