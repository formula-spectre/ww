#!/bin/bash

# A shell script designed to be executed by qBittorrent's "Run external program on torrent completion"
# This script will send an ntfy notification with information from the qBittorrent when a download is complete
#
# An example how to fill in qBittorrent's "Run external program on torrent completion" to execute this script
# /bin/bash -c "chmod +x /path/to/qbittorrent-notify.sh; /path/to/qbittorrent-slack-notify.sh '%N' '%Z' 'https://ntfy.sh/qbittorrent"
#
# Supported parameters (case sensitive):
# - %N: Torrent name
# - %L: Category
# - %G: Tags (separated by comma)
# - %F: Content path (same as root path for multifile torrent)
# - %R: Root path (first torrent subdirectory path)
# - %D: Save path
# - %C: Number of files
# - %Z: Torrent size (bytes)
# - %T: Current tracker
# - %I: Info hash

# https://unix.stackexchange.com/a/259254
bytesToHuman() {
    b=${1:-0}; d=''; s=0; S=(Bytes {K,M,G,T,P,E,Y,Z}iB)
    while ((b > 1024)); do
        d="$(printf ".%02d" $((b % 1024 * 100 / 1024)))"
        b=$((b / 1024))
        let s++
    done
    echo "$b$d ${S[$s]}"
}

name="$1"
if [ -z "$name" ]; then 
    echo "ERROR: Expected <name> as the 1st argument but none given, <name> should be the Torrent name (\"%N\") from qBittorrent"
    exit 1
fi

sizeBytes="$2"
if [ -z "$sizeBytes" ]; then 
    echo "ERROR: Expected <size> as the 2nd argument but none given, <size> should be the Torrent size (bytes) (\"%Z\") from qBittorrent"
    exit 1
fi
size=`bytesToHuman $sizeBytes`

qbittorrent_webhook="$3"
if [ -z "$qbittorrent_webhook" ]; then 
    echo "ERROR: Expected <qbittorrent_webhook> as the 3rd argument but none given"
    exit 1
fi

ts=`date "+%s"`

/usr/bin/curl -H "Title: Download Completed" -H "Tags: white_check_mark, qBittorrent" -d "Name: $name
Size: $size" $qbittorrent_webhook
