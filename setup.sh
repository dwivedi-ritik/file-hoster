#! /usr/bin/bash

go build -o build/filehost main.go

mkdir $HOME/.config/filehost

shell=$(echo `(which $SHELL)` | awk -F "/" '{print $3}')

echo "alias filehost='$(pwd)/build/filehost'" >> ~/.${shell}rc   
