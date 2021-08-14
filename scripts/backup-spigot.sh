#!/bin/bash -e
cd /mnt/disks/game-data
backup_file=spigot-$(date +%Y%m%d-%H%M%S).tgz
tar zcvf $backup_file spigot
gsutil cp $backup_file gs://minecraft-abekoh-backup/
rm $backup_file
