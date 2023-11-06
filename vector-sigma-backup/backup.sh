#!/bin/sh
[ -d ~/ww/vector-sigma-backup ] || mkdir -p ~/ww/vector-sigma-backup
cp  -- /etc/fstab ~/ww/vector-sigma-backup/fstab
cp -- ~/backup.sh ~/ww/vector-sigma-backup/backup.sh
eval $(ssh-agent)
ssh-add ~/.ssh/id_ed25519-gh

cd ~/ww
git add *
git commit -m "autosync from vector-sigma @ $(date +%c)"
git push
cd ~/
killall ssh-agent
