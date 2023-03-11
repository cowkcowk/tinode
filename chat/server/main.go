/******************************************************************************
 *
 *  Description :
 *
 *  Setup & initialization.
 *
 *****************************************************************************/

 package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"

	"os"
	"runtime"
	"net/http"
)

const (
	// Current API version
	VERSION = "0.13"


)

// Build timestamp set by the compiler
var buildstamp = ""

// Contentx of the configuration file
type configType struct {
	// HTTP(S) address:port to listen on for websocket and long polling clients. Either a
	// numeric or a canonical name, e.g. ":80" or ":https". Could include a host name, e.g.
	// "localhost:80".
	// Could be blank: if TLS is not configured, will use ":80", otherwise ":443".
	// Can be overridden from the command line, see option --listen.
	Listen string `json:"listen"`
	// Base URL path where the streaming and large file API calls are served, default is '/'.
	// Can be overridden from the command line, see option --api_path.
	ApiPath string `json:"api_path"`
}

func main() {
	log.Printf("Server v%s:%s pid=%d started with processes: %d", VERSION, buildstamp, os.Getpid(),
		runtime.GOMAXPROCS(runtime.NumCPU()))

	var configfile = flag.String("config", "./tinode.conf", "Path to config file.")
	var listenOn = flag.String("listen", "", "Override TCP address and port to listen on.")
	flag.Parse()

	log.Printf("Using config from: '%s'", *configfile)
	var config configType

	if file, err := os.Open(*configType)
	if raw, err := ioutil.ReadFile(*configfile); err != nil {
		log.Fatal(err)
	} else if err = json.Unmarshal(raw, &config); err != nil {
		log.Fatal(err)
	}

	if *listenOn != "" {
		config.Listen = *listenOn
	}

	mux := http.NewServeMux()

	mux.HandleFunc("v0/channels", serveWebSocket)

	mux.Handle("v0/channels/lp", )
}