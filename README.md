# Overview

This is the kickstart Go module for a full-stack developer learning club

## Running the HTTP server

To run this app as a HTTP server, run the following in your Terminal/Command Line/text editor console:

```
go run main.go
```

or:

```
go run try-go

```

in which `try-go` is the module name as defined in the `go.mod` file.


## Calling the HTTP server

You can call the two provided endpoints with base URL `http://localhost:80` with a HTTP client of your choice

### Web browser

![/hello in web browser](browser-hello.png)

![/foobar in web browser](browser-json.png)

### curl

[curl](https://curl.se/) is a command line tool that includes a HTTP client

![/hello in curl](curl-hello-response.png)

![/foobar in curl](curl-json-response.png)

### Postman

[Postman](https://www.postman.com/) is an API platform for building and using APIs

![/hello response body in Postman](postman-hello-response-body.png)

![/hello response headers in Postman](postman-hello-response-headers.png)

![/foobar response body in Postman](postman-json-response-body.png)

![/foobar response headers in Postman](postman-json-response-headers.png)