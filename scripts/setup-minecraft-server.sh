#!/bin/bash -e

# install dependencies
apt-get update
apt-get install -y openjdk-16-jre-headless

# setup permenent disk
mkdir -p /mnt/disks/game-data
mount -o discard,defaults UUID=bc7783d6-280f-4214-930d-3d8a172fccbc /mnt/disks/game-data
chmod a+w /mnt/disks/game-data

# run minecraft
screen -d -m -S minecraft java -Xmx2048M -Xms2048M -jar /mnt/disks/game-data/minecraft-java/server.jar nogui