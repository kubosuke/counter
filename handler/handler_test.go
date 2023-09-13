package handler_test

import (
	"counter/handler"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func Test_CounterHandler(t *testing.T) {
	for i := 0; i < 10; i++ {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()
		handler.CounterHandler(w, req)
		res := w.Result()
		defer res.Body.Close()
		data, err := io.ReadAll(res.Body)
		if err != nil {
			t.Errorf("expected error to be nil got %v", err)
		}
		if string(data) != fmt.Sprintln(strconv.Itoa(i)) {
			t.Errorf("expected %d got %v", i, string(data))
		}
	}
}
