#!/bin/sh

for i in $(ls *.go)
do
    printf "Compiling '$i'..."
    go build -compiler gccgo $i
    printf " Done\n"
done
