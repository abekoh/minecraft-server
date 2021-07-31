#!/bin/bash -e

# install dependencies
wget https://golang.org/dl/go1.16.6.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.16.6.linux-amd64.tar.gz

mkdir -p $HOME/bin
export PATH=$PATH:$HOME/bin

gsutil cp gs://minecraft-abekoh-bot/bot $HOME/bin/bot
chmod +x $HOME/bin/bot

export DISCORD_TOKEN=$(gcloud secrets versions access latest --secret="discord-bot-token")

screen -d -m -S discord-bot bot -project minecraft-abekoh -zone asia-northeast1-b -name minecraft-abekoh
