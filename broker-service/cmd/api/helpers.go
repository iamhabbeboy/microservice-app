package main

type Payload struct {
	Name string `json:"name"`
	Data any    `json:"data"`
}

func statusResponse(status bool, err string) JsonResponse {
	return JsonResponse{
		Error:   status,
		Message: err,
	}
}
