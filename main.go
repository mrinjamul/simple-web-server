package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	colour "github.com/fatih/color"
	flag "github.com/spf13/pflag"
)

var (
	dir  string
	port string
)

var (
	flagDir     string
	flagPort    string
	flagHTTPS   bool
	flagHelp    bool
	flagVersion bool
)

func main() {
	// parse flags
	flag.Parse()

	if flagHelp {
		printUsage()
	}

	if flagVersion {
		printVersion()
	}

	if flagDir != "" {
		dir = flagDir
	} else {
		dir = "./"
	}

	if flagHTTPS {
		port = "443"
	}

	if flagPort != "" {
		port = flagPort
	} else {
		port = "8080"
	}

	if flagHTTPS {
		colour.Blue("Starting Web Server over HTTPS...")

	} else {
		colour.Blue("Starting Web Server...")
	}

	colour.Blue("Starting on port " + port)

	http.Handle("/", http.FileServer(http.Dir(dir)))
	if flagHTTPS {
		log.Fatal(http.ListenAndServeTLS(":"+port, "server.crt", "server.key", nil))
	} else {
		log.Fatal(http.ListenAndServe(":"+port, nil))
	}
}

func init() {
	flag.StringVarP(&flagDir, "dir", "d", "", "directory to serve")
	flag.StringVarP(&flagPort, "port", "p", "", "set port to serve")
	flag.BoolVarP(&flagHTTPS, "https", "S", false, "serve over HTTPS")
	flag.BoolVarP(&flagHelp, "help", "h", false, "help message")
	flag.BoolVarP(&flagVersion, "version", "v", false, "print version")

}

func printUsage() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
	os.Exit(0)
}
func printVersion() {
	version := "1.0.0"
	fmt.Println(os.Args[0] + " version " + version)
	os.Exit(0)
}
