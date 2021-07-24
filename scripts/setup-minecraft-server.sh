#!/bin/bash

# install dependencies
apt-get update
apt-get install -y openjdk-16-jre-headless

# setup permenent disk
mkdir -p /mnt/disks/game-data
mount -o discard,defaults /dev/sdb /mnt/disks/game-data
chmod a+w /mnt/disks/game-data

# make symbolic link
ln -sf /mnt/disks/game-data/minecraft-java ~/minecraft-java

# run minecraft
RUN_COMMAND="java -Xmx2048M -Xms2048M -jar ~/minecraft-java/server.jar nogui"
tmux new-session -d -s my_session $RUN_COMMAND