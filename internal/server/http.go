package server

import (
	"context"
	"log"
	"net/http"
	"noun-guesser/internal/tree"
)

type Server struct {
	s *http.ServeMux

	t    *tree.Word
	addr string
}

func NewServer(addr string, tree *tree.Word) *Server {
	return &Server{
		addr: addr,
		t:    tree,
		s:    http.NewServeMux(),
	}
}

func (s *Server) Run(ctx context.Context) {
	s.s.HandleFunc("/find", s.findHandler)

	server := &http.Server{
		Addr:    s.addr,
		Handler: s.s,
	}

	go func() {
		log.Println("server is starting")
		err := server.ListenAndServe()
		if err != nil {
			log.Println("server listen: ", err.Error())
		}
	}()

	<-ctx.Done()
	log.Println("server is stopping")
	server.Shutdown(context.Background())
}
