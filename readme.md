### A File Hoster written in Go 

A lightweight file hoster server that can host your file.

```shell
$ curl -F "data=@yourfilename" 
```
this will response with unique url

```shell
$ curl -O -J "url"
```

## Install

Build using below commands 
```shell
$ go build -o filehost main.go

```

run `./filehost` as service or executable.

## TODOS
[] - URL Shortner.
[] - Once file retrived it should be delete.
[] - Set Download count limit.
[] - Making it more structured.