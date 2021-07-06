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
	dir     string
	port    string
	sslKey  string
	sslCert string
)

var (
	flagDir     string
	flagPort    string
	flagKey     string
	flagCert    string
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

	if flagPort != "" {
		port = flagPort
	} else {
		port = "8080"
	}

	if flagHTTPS {
		if flagPort == "" {
			port = "443"
		}
		if flagKey != "" {
			sslKey = flagKey
		} else {
			sslKey = "server.key"
		}

		if flagCert != "" {
			sslCert = flagCert
		} else {
			sslCert = "server.crt"
		}
	}

	if flagHTTPS {
		colour.Blue("Starting Web Server over HTTPS...")

	} else {
		colour.Blue("Starting Web Server...")
	}

	_, present := os.LookupEnv("PORT")
	if present {
			port = os.Getenv("PORT")
	}
	
	colour.Blue("Starting on port " + port)

	http.Handle("/", http.FileServer(http.Dir(dir)))
	if flagHTTPS {
		log.Fatal(http.ListenAndServeTLS(":"+port, sslCert, sslKey, nil))
	} else {
		log.Fatal(http.ListenAndServe(":"+port, nil))
	}
}

func init() {
	flag.StringVarP(&flagDir, "dir", "d", "", "directory to serve")
	flag.StringVarP(&flagPort, "port", "p", "", "set port to serve")
	flag.StringVarP(&flagKey, "key", "k", "", "openssl key location for HTTPS")
	flag.StringVarP(&flagCert, "cert", "C", "", "openssl cert location for HTTPS")
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
	version := "1.1.0"
	fmt.Println(os.Args[0] + " version " + version)
	os.Exit(0)
}
