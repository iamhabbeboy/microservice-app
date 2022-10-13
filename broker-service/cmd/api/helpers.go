package main

func statusResponse(status bool, err string) JsonResponse {
	return JsonResponse{
		Error:   status,
		Message: err,
	}
}
