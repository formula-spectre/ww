package main 

import (
	"syscall"
)

func main () {
    const MS_NOSUID = 0x2
    syscall.Mount("/dev/disk/by-uuid/3dcaa245-20c7-43b1-ada9-9e65a8f290b8", "/mnt", "btrfs" , MS_NOSUID,  "subvolid=260,subvol=/qbittorrent-data");
}

