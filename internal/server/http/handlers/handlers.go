package handlers

type Handlers struct {
	HubClientHandler *HubClientHandler
}

func NewHandlers(hubClientHandler *HubClientHandler) *Handlers {
	return &Handlers{
		HubClientHandler: hubClientHandler,
	}
}
