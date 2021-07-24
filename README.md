# minecraft-server

## 手動作業

### 永続ディスクの初期化

See: https://cloud.google.com/compute/docs/disks/add-persistent-disk

```bash
sudo mkfs.ext4 -m 0 -E lazy_itable_init=0,lazy_journal_init=0,discard /dev/sdb
sudo mkdir -p /mnt/disks/game-data
sudo mount -o discard,defaults /dev/sdb /mnt/disks/game-data
sudo chmod a+w /mnt/disks/game-data
```

### Minecraft のセットアップ

永続ディスクの`/game-data/minecraft-java`内に`server.jar`など設置

### 起動スクリプト設置

```bash
gsutil cp scripts/setup-minecraft-server.sh gs://minecraft-abekoh-scripts/
```
