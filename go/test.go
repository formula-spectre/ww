package main

import "fmt"

/*
import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func mountNecessaryFileSystems() {
	const MS_NOSUID = 0x2
	syscall.Mount("/dev/disk/by-uuid/3dcaa245-20c7-43b1-ada9-9e65a8f290b8",
		"/srv/jellyfin", "btrfs", MS_NOSUID, "subvolid=258,subvol=/jellyfin-data")

	syscall.Mount("/dev/disk/by-uuid/3dcaa245-20c7-43b1-ada9-9e65a8f290b8",
		"/vw-data", "btrfs", MS_NOSUID, "subvolid=256,subvol=/vw-data")

	syscall.Mount("/dev/disk/by-uuid/3dcaa245-20c7-43b1-ada9-9e65a8f290b8",
		"/torrents",
		"btrfs",
		MS_NOSUID,
		"subvolid=259,subvol=/torrents")

	syscall.Mount("/dev/disk/by-uuid/3dcaa245-20c7-43b1-ada9-9e65a8f290b8",
		"/media",
		"btrfs",
		MS_NOSUID,
		"subvolid=257,subvol=/media")
	/*
		syscall.Mount("/dev/disk/by-uuid/3dcaa245-20c7-43b1-ada9-9e65a8f290b8",
			"/mnt",
			"btrfs",
			MS_NOSUID,
			"subvolid=260,subvol=/qbittorrent-data")
*/
/*

	syscall.Mount("/dev/disk/by-uuid/3dcaa245-20c7-43b1-ada9-9e65a8f290b8",
		"/srv/qbittorrent",
		"btrfs",
		MS_NOSUID,
		"subvolid=260,subvol=/qbittorrent-data")
}
func podman(args ...string) error {
	podman := exec.Command("/usr/bin/podman", args...)
	podman.Env = expandPath(os.Environ())
	podman.Env = append(podman.Env, "TMPDIR=/tmp")
	podman.Stdin = os.Stdin
	podman.Stdout = os.Stdout
	podman.Stderr = os.Stderr
	if err := podman.Run(); err != nil {
		return fmt.Errorf("%v: %v", podman.Args, err)
	}
	return nil
}
*/
const dockerComposeFile string = `
version: "3.5"
services:
  qbittorrent:
    image: lscr.io/linuxserver/qbittorrent:latest
    container_name: qbittorrent
    environment:
      - PUID=1000
      - PGID=1000
      - TZ=Etc/UTC
      - WEBUI_PORT=8080
    volumes:
      - /srv/qbittorrent:/config
      - /torrents:/downloads
    ports:
      - 8080:8080
      - 6881:6881
      - 6881:6881/udp
    restart: unless-stopped
  vaultwarden:
    image: vaultwarden/server:latest
    container_name: vaultwarden
    restart: unless-stopped
    volumes:
      - /vw-data:/data
    ports:
    - 8001:80
  jellyfin:
    iamge: jellyfin/jellyfin
    container_name: jellyfin
    network_mode: host
    volumes:
      - /srv/jellyfin:/config
      - /srv/jellyfin:/cache
      - /media/films:/films
      - /media/tvshows:/tvshows
    restart: unless-stopped
  portainer-agent:
    ports:
        - '9001:9001'
    container_name: portainer_agent
    restart: always
    volumes:
        - '/var/run/docker.sock:/var/run/docker.sock'
        - '/var/lib/docker/volumes:/var/lib/docker/volumes'
    image: 'portainer/agent:2.19.1'

`
const fstabEntriesToAppend string = `
# /srv/jellyfin
UUID=3dcaa245-20c7-43b1-ada9-9e65a8f290b8	/srv/jellyfin	btrfs     	rw,relatime,space_cache=v2,subvolid=258,subvol=/jellyfin-data	0 0

# /vw-data
UUID=3dcaa245-20c7-43b1-ada9-9e65a8f290b8	/vw-data  	btrfs     	rw,relatime,space_cache=v2,subvolid=256,subvol=/vw-data	0 0

# /torrents
UUID=3dcaa245-20c7-43b1-ada9-9e65a8f290b8	/torrents 	btrfs     	rw,relatime,space_cache=v2,subvolid=259,subvol=/torrents	0 0

# /media
UUID=3dcaa245-20c7-43b1-ada9-9e65a8f290b8	/media    	btrfs     	rw,relatime,space_cache=v2,subvolid=257,subvol=/media	0 0

# /mnt
#UUID=3dcaa245-20c7-43b1-ada9-9e65a8f290b8	/mnt      	btrfs     	rw,relatime,space_cache=v2,subvolid=5,subvol=/	0 0

# /srv/qbittorrent
UUID=3dcaa245-20c7-43b1-ada9-9e65a8f290b8	/srv/qbittorrent	btrfs     	rw,relatime,space_cache=v2,subvolid=260,subvol=/qbittorrent-data	0 0

`

func main() {
	fmt.Println(dockerComposeFile)
}
