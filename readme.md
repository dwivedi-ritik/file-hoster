### A File Hoster written in Go 

A lightweight file hoster server that can host your file.

```shell
$ curl -F "data=@yourfilename" 
```
this will response with unique url

```shell
$ curl -O -J "url"
```