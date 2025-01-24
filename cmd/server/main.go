package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/HDuoDuo/metatube/cmd"
	"github.com/HDuoDuo/metatube/engine"
	V "github.com/HDuoDuo/metatube/internal/version"
)

func showVersionAndExit() {
	fmt.Println(V.BuildString())
	os.Exit(0)
}

func main() {
	if _, isSet := os.LookupEnv("VERSION"); cmd.Config.VersionFlag &&
		!isSet /* NOTE: ignore this flag if ENV contains VERSION variable. */ {
		showVersionAndExit()
	}

	var (
		addr = net.JoinHostPort(
			cmd.Config.Bind,
			cmd.Config.Port)
		router = cmd.Router(engine.DefaultEngineName)
	)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatal(err)
	}
}
