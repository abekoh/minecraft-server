#!/bin/bash -e

# install dependencies
wget https://golang.org/dl/go1.16.6.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.16.6.linux-amd64.tar.gz

export PATH=$PATH:/usr/local/go/bin

go get github.com/abekoh/minecraft-server/bot

screen -d -m -S discord-bot bot -project minecraft-abekoh -zone asia-northeast1-b -name minecraft-abekoh