package ZoeiDNS

import (
	"github.com/miekg/dns"
	"math/rand"
)

type Server struct {
	Instance *dns.Server
	Serial   uint32
	Handler  func(w dns.ResponseWriter, r *dns.Msg)
}

func NewServer() (s *Server) {
	s.Serial = rand.Uint32()
	s.Handler = func(w dns.ResponseWriter, r *dns.Msg) {
		// TODO: write the handler function's body.
	}

	return
}

func (s *Server) Start() {
	s.Instance = &dns.Server{
		Addr: ":53",
		Net:  "udp",
	}

	go func() {
		err := s.Instance.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()
}
