package server

import (
	"encoding/json"
	"io"
	"net/http"
)

type findRequest struct {
	Ignored string `json:"ignored"`
	Guessed string `json:"guessed"`
	Pattern string `json:"pattern"`
	Limit   uint32 `json:"limit"`
}

type findResponse struct {
	Found []string `json:"found"`
}

func (s *Server) findHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte{})

		return
	}

	if r.Body == http.NoBody {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no body"))

		return
	}

	req := &findRequest{}
	bb, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad body"))

		return
	}

	err = json.Unmarshal(bb, req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad body"))

		return
	}

	ignored := []rune{}

	for _, r := range req.Ignored {
		ignored = append(ignored, r)
	}

	guessed := []rune{}

	for _, r := range req.Guessed {
		guessed = append(guessed, r)
	}

	res := s.t.Find(req.Pattern, ignored, guessed)
	if req.Limit != 0 && len(res) > int(req.Limit) {
		res = res[:req.Limit]
	}

	bb, err = json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("cannot marshal response body"))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bb)
}
