[![Build Status](https://travis-ci.org/jad21/fetch.svg?branch=master)](https://travis-ci.org/jad21/fetch)
[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/jad21/fetch)
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/jad21/fetch/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/jad21/fetch)](https://goreportcard.com/report/github.com/jad21/fetch)
 
# Fetch HTTP Client

Búsqueda simple realizada en Go para simplificar la vida del programador.


## Install

> Default 
```shell
go get github.com/jad21/fetch
```

## Import

```go
import (
  "github.com/jad21/fetch"
)
```

## Test 
Para ejecutar la prueba del proyecto

```shell
go test -v --cover
```


## Example: 

#### Simple
    
```go
response, err := fetch.Get("https://httpbin.org/get/", nil)
``` 

#### Custom Headers

```go
header := http.Header{
    "Content-Type": []string{"application/json"},
    "User-Agent":   []string{"My-User-Agent"},
}
f := fetch.New(fetch.WithHeader(header))
rsp, err := f.Get("https://httpbin.org/headers", nil)
```

#### Simple JSON POST

```go
login := map[string]interface{}{
	"username": "jad21",
	"password": "loremIpsum",
}
response, err := fetch.
		IsJSON().
		Post("https://httpbin.org/post/", fetch.NewReader(login))
```

  
