package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/kardianos/service"
)

var logger service.Logger

type program struct{}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Go - Hello World</h1>")
}

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}
func (p *program) run() {
	// Do work here

	// Run command: 
	// C:\Users\Jiean\Desktop\work\supervisord-research\supervisord.exe --configuration=\"C:\Users\Jiean\Desktop\work\supervisord-research\supervisord-windows.conf\"

	fmt.Println("http.exe test")
	cmd := exec.Command("http.exe")
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
		panic(err)
	}

	// * fmt.Println("simple test")
	// http.HandleFunc("/", rootHandler)

	// err := http.ListenAndServe(":9001", nil)
	// if err != nil {
	// 	panic(err)
	// }
}
func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        "GoServiceTest",
		DisplayName: "Go Service Test",
		Description: "This is a test Go service.",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}