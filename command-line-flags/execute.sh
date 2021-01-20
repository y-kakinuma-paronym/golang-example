#!/bin/sh
go build command-line-flags.go

./command-line-flags -word=opt -numb=7 -fork -svar=flag
./command-line-flags -word=opt
./command-line-flags -word=opt a1 a2 a3
./command-line-flags -word=opt a1 a2 a3 -numb=7
./command-line-flags -h
./command-line-flags -wat