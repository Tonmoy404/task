package cmd

import (
	"log"

	"github.com/Tonmoy404/hello-world/repo"
	"github.com/Tonmoy404/hello-world/rest"
	"github.com/Tonmoy404/hello-world/service"
)

func serveRest() {
	stdRepo := repo.NewStudentRepo()
	svc := service.NewService(stdRepo)

	server, err := rest.NewServer(svc)
	if err != nil {
		log.Panic("cannot start the server", err)
	}

	server.Start()
}
