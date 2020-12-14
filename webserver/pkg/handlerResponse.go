package pkg

func getResponse(webserver Webserver, recordName string, amount int) (HandlerResponse, error) {
	entities, err := webserver.GenerateEntity(recordName, amount)

	return HandlerResponse{
		entities,
		map[string]string{
			"Content-Type": "application/json",
		},
	}, err
}

type HandlerResponse struct {
	Body    interface{}
	Headers map[string]string
}
