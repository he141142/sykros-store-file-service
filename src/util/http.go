package util

import (
	"github.com/bytedance/sonic"
	"net/http"
)

func HttpResponse[T any](w *http.ResponseWriter, code int, v T) error {
	b, err := sonic.Marshal(v)
	if err != nil {
		return err
	}
	(*w).Header().Set("Content-Type", "application/json")
	(*w).WriteHeader(code)
	_, err = (*w).Write(b)
	if err != nil {
		return err
	}
	return nil
}
