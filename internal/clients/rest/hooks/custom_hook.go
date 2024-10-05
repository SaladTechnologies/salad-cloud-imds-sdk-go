package hooks

type CustomHook struct{}

func NewCustomHook() Hook {
	return &CustomHook{}
}

func (h *CustomHook) BeforeRequest(req Request, params map[string]string) Request {
	req.SetHeader("Metadata", "true")
	return req
}

func (h *CustomHook) AfterResponse(req Request, resp Response, params map[string]string) Response {
	return resp
}

func (h *CustomHook) OnError(req Request, resp ErrorResponse, params map[string]string) ErrorResponse {
	return resp
}
