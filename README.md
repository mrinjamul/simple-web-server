# simple web server

A simple (static) web server written in golang

## Build and Run

Install by go,

```sh
go get github.com/mrinjamul/simple-web-server

```

Running with port 8081 and current directory,

```sh
simple-web-server -p 8081 -d "./" # port and directory optional
```

Running over HTTPS requires openssl key and openssl certificate to run.

Remove demo key and certificate and generate with name `server.key` and `server.crt`.

To generate key and certificate,

```sh
$ openssl genrsa -out server.key 2048
$ openssl ecparam -genkey -name secp384r1 -out server.key
$ openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
```

Running over HTTPS,

```sh
simple-web-server -d "./" --https -p 8443 # port and directory optional
```

## Some defaults

- default port 8080 for http
- default port 443 for https (may require root to run)
- defualt directory "./"

## Usage

    Usage: simple-web-server [options]
    Options:
    -d, --dir string    directory to serve
    -h, --help          help message
    -S, --https         serve over HTTPS
    -p, --port string   set port to serve
    -v, --version       print version

## License

- licensed under MIT
