//go:build windows

package main

import (
	"flag"
	"hyper-v-rest-wmi/rest"
	"hyper-v-rest-wmi/utilities"
	"os"

	"github.com/kardianos/service"
)

var log = utilities.Log
var logger service.Logger

type program struct {
	exit chan struct{}
}

func (p *program) Start(s service.Service) error {
	if service.Interactive() {
		logger.Info("Running in interactive terminal.")
	} else {
		logger.Info("Running under service manager.")
	}
	p.exit = make(chan struct{})

	go p.run()

	return nil
}

func (p *program) run() {
	c := utilities.ParseConfig()
	rest.StartServer(c.Port, "1.0.0")
}

func (p *program) Stop(s service.Service) error {
	if service.Interactive() {
		os.Exit(0)
	} else {
		logger.Info("Stopping service...")
		close(p.exit)
	}
	return nil
}

func main() {
	svcFlag := flag.String("service", "", "Control the system service.")
	flag.Parse()

	options := make(service.KeyValue)
	options["Restart"] = "on-success"
	options["SuccessExitStatus"] = "1 2 8 SIGKILL"
	svcConfig := &service.Config{
		Name:        "hyper-v-rest-wmi",
		DisplayName: "Hyper-V REST WMI Service",
		Description: "Simple REST service for querying VMs in Hyper-V via WMI",
		Option:      options,
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	errs := make(chan error, 5)
	logger, err = s.Logger(errs)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			err := <-errs
			if err != nil {
				log.Print(err)
			}
		}
	}()

	if len(*svcFlag) != 0 {
		err := service.Control(s, *svcFlag)
		if err != nil {
			log.Printf("Valid actions: %q\n", service.ControlAction)
			log.Fatal(err)
		}
		return
	}
	err = s.Run()
	if err != nil {
		logger.Error(err)
	}

	select {}
}
