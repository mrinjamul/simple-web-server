package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/pkg/browser"
	flag "github.com/spf13/pflag"
)

var (
	AppName   = "sws"
	Version   = "dev"
	GitCommit = "none"
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
	flagWeb     bool
	flagHelp    bool
	flagVersion bool
)

var (
	config *Config
)

func init() {
	// Load Config from file
	config = GetConfig()
}

func main() {

	// Get config from file
	flagDir = config.Dir
	flagPort = config.Port
	flagKey = config.SslCert
	flagCert = config.SslCert
	flagHTTPS = config.HTTPS

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
		log.Println("Starting Web Server over HTTPS...")

	} else {
		log.Println("Starting Web Server...")
	}

	_, present := os.LookupEnv("PORT")
	if present {
		port = os.Getenv("PORT")
	}

	_, present = os.LookupEnv("DIR")
	if present {
		dir = os.Getenv("DIR")
	}

	log.Println("Starting on port " + port)

	url := "http://"
	if flagHTTPS {
		url = "https://"
	}
	url += "localhost" + ":" + port
	if flagWeb {
		err := browser.OpenURL(url)
		log.Println(err)
	}

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
	flag.BoolVarP(&flagWeb, "web", "w", false, "open url in web browser")
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
	fmt.Println(os.Args[0] + " version " + Version + "-" + GitCommit)
	os.Exit(0)
}
