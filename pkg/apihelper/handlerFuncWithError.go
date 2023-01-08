package apihelper

import "net/http"

type HandlerFuncWithError func(http.ResponseWriter, *http.Request) error
