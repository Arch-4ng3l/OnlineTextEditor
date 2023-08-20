package server

import (
	"encoding/json"
	"net/http"

	"example.com/test/frontend"
	"example.com/test/lang"
	"example.com/test/objects"
)

type Server struct {
	listningAddr string
}

func NewServer(addr string) *Server {
	return &Server{
		addr,
	}
}

func (s *Server) Init() {
	frontend.InitFrontEnd()
	http.HandleFunc("/api/code", s.RecvCode)

	http.ListenAndServe(s.listningAddr, nil)
}

func (s *Server) RecvCode(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	encoder := json.NewEncoder(w)
	obj := &objects.Code{}
	decoder.Decode(obj)
	s.executeCode(obj.Text, encoder)
}

func (s *Server) executeCode(code string, encoder *json.Encoder) {
	output := lang.ExecCode(code)
	encoder.Encode(map[string]string{"output": output})
}
