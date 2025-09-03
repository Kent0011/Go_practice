#!/bin/bash
acc new $1 -c all
for question in a b c d e f g;
do
    touch $1/$question/main.go
done