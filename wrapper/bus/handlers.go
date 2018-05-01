package bus

import "github.com/nvbn/shell_logger/wrapper/storage"

func startListening(store storage.Storage, bytes []byte) []byte {
	request, err := newStartListeningRequest(bytes)

	if err != nil {
		return errorResponse(err)
	}

	store.StartListening(request.Time)

	return okResponse
}

func stopListening(store storage.Storage, bytes []byte) []byte {
	request, err := newStopListeningRequest(bytes)

	if err != nil {
		return errorResponse(err)
	}

	store.StopListening(request.Command, request.ReturnCode, request.Time)

	return okResponse
}

func list(store storage.Storage, bytes []byte) []byte {
	request, err := newListRequest(bytes)

	if err != nil {
		return errorResponse(err)
	}

	commands := store.List(request.Count)

	return listResponse(commands)
}
