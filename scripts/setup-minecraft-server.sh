#!/bin/bash -e

# install dependencies
apt-get update
apt-get install -y openjdk-16-jre-headless

# setup permenent disk
mkdir -p /mnt/disks/game-data
mount -o discard,defaults UUID=bc7783d6-280f-4214-930d-3d8a172fccbc /mnt/disks/game-data
chmod a+w /mnt/disks/game-data

# run minecraft
cd /mnt/disks/game-data/spigot
screen -d -m -S minecraft java -Xmx3072M -Xms3072M -jar spigot-1.17.1.jar nogui

# backup cron
backup_script=/etc/cron.weekly/backup-spigot.sh
gsutil cp gs://minecraft-abekoh-scripts/backup-spigot.sh $backup_script
chmod +x $backup_script
