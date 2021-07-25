# minecraft-server

## 手動作業

### 永続ディスクの初期化

See: https://cloud.google.com/compute/docs/disks/add-persistent-disk

```bash
sudo mkfs.ext4 -m 0 -E lazy_itable_init=0,lazy_journal_init=0,discard /dev/sdb
sudo mkdir -p /mnt/disks/game-data
sudo mount -o discard,defaults UUID=bc7783d6-280f-4214-930d-3d8a172fccbc /mnt/disks/game-data
sudo chmod a+w /mnt/disks/game-data
```

### Minecraft のセットアップ

Spigot を永続ディスク`/game-data/spigot`に構築

See: https://www.spigotmc.org/wiki/buildtools/

```
wget https://hub.spigotmc.org/jenkins/job/BuildTools/lastSuccessfulBuild/artifact/target/BuildTools.jar
java -jar BuildTools.jar --rev 1.17.1
touch start-server.sh
// scripts/start-server.shをコピー
chmod +x start-server.sh
```

### 起動スクリプト設置

```bash
gsutil cp scripts/setup-minecraft-server.sh gs://minecraft-abekoh-scripts/
```
c