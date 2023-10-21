package main

import (
	"net/http"

	"github.com/joshm998/printfarm/config"
	"github.com/joshm998/printfarm/server"
	log "github.com/sirupsen/logrus"
)

func main() {
	s := server.New()
	log.Info("Listening on port:", config.GetYamlValues().ServerConfig.Port)
	err := s.ListenAndServe()
	if err != http.ErrServerClosed {
		log.Fatalf("Listen: %s\n", err)
	}
	log.Info("service stopped")

}
