#! /bin/bash

clear
ls -ltrh archive
echo "enter file number: "
read -r fileno

git add --all
git commit -m 'before creating new source file'

mv src/app/main.go archive/main_"$fileno".go

now="$(date +"%T")"

printf '%s\n' "package main

import (
    \"fmt\"
)

func main() {

    fmt.Println(\"\n #($now): The End ...\")
}" >> src/app/main.go

git add --all
git commit -m 'before creating new source file'

