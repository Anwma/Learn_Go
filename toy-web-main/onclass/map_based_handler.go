package main

import "net/http"

type HandlerBasedOnMap struct {
	handlers map[string]func(ctx *Context)
}

func (h *HandlerBasedOnMap) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	key := h.key(request)
	if handler, ok := h.handlers[key]; ok {
		handler(NewContext(writer, request))
	} else {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("Not Found"))
	}
	//h.handlers[key] =
}

func (h *HandlerBasedOnMap) key(method string, pattern string) string {
	//TODO implement me
	return method + "#" + pattern
}
