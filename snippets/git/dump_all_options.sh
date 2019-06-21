#!/usr/bin/env zsh

for x in $(cat sub.txt | awk '{print $1}'); do git ${x} -h > ${x}_option.txt 2>&1; done
