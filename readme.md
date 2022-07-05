## A File Hosting server written in Go 

A lightweight file hosting server that can host your file.

```shell
$ curl -F "data=@yourfilename" 
```
this will respond with the unique url.

```shell
$ curl -JO "url"
```

## Install

- if you are on linux machine just run `setup.sh`.
- setup will create a database directory at `$HOME/.config/filehost/database.db`.
- and add aliad to your default shell.
- or you can build using below commands 
```shell
$ go build -o filehost main.go

```

run `./filehost` as service or executable.

## TODOS
- [ ]  URL Shortner.
- [ ]  Once file retrived it should be delete.
- [ ]  Set Download count limit.
- [ ]  Making it more structured.
