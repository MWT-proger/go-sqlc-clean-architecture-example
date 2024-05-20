package handlers

type APIHandler struct {
	UserService UserServicer
}

func NewAPIHandler(
	userService UserServicer,

) (h *APIHandler, err error) {
	hh := &APIHandler{
		UserService: userService,
	}

	return hh, err
}
