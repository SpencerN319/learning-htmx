package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *server) handleHealth() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.WriteHeader(int(s.status.Load()))
		w.Write([]byte(`{ "status": "UP" }`))
	}
}

func (s *server) handleIndex() httprouter.Handle {
	data := struct {
		Title string
	}{
		Title: "index",
	}
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("content-type", "text/html")
		w.WriteHeader(http.StatusOK)
		if err := s.tmpls.ExecuteTemplate(w, "index", data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
