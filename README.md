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

```bash
wget https://hub.spigotmc.org/jenkins/job/BuildTools/lastSuccessfulBuild/artifact/target/BuildTools.jar
java -jar BuildTools.jar --rev 1.17.1
```

### Ops エージェント設定

Terraform の設定は少しめんどそうだったのでとりあえず手動で。

See: https://cloud.google.com/logging/docs/agent/ops-agent/installation

```bash
curl -sSO https://dl.google.com/cloudagents/add-google-cloud-ops-agent-repo.sh
sudo bash add-google-cloud-ops-agent-repo.sh --also-install
rm add-google-cloud-ops-agent-repo.sh
```

### Secret 設定

```
gcloud secrets create discord-bot-token --replication-policy="automatic"
echo -n "BOT_TOKEN" | gcloud secrets versions add discord-bot-token --data-file=-
```
