package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

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

func irssi() error {
	// Ensure we have an up-to-date clock, which in turn also means that
	// networking is up. This is relevant because podman takes what’s in
	// /etc/resolv.conf (nothing at boot) and holds on to it, meaning your
	// container will never have working networking if it starts too early.
	//gokrazy.WaitForClock()


	if err := podman("kill", "irssi"); err != nil {
		log.Print(err)
	}

	if err := podman("rm", "irssi"); err != nil {
		log.Print(err)
	}

	// You could podman pull here.

	if err := podman("run",
		"-td",
		"-v", "/perm/irssi:/home/michael/.irssi",
		"-v", "/perm/irclogs:/home/michael/irclogs",
		"-e", "TERM=rxvt-unicode",
		"-e", "LANG=C.UTF-8",
		"--network", "host",
		"--name", "irssi",
		"docker.io/stapelberg/irssi:latest",
		"screen", "-S", "irssi", "irssi"); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := irssi(); err != nil {
		log.Fatal(err)
	}
}

// mountVar bind-mounts /perm/container-storage to /var if needed.
// This could be handled by an fstab(5) feature in gokrazy in the future.
// expandPath returns env, but with PATH= modified or added
// such that both /user and /usr/local/bin are included, which podman needs.
func expandPath(env []string) []string {
	extra := "/user:/usr/local/bin"
	found := false
	for idx, val := range env {
		parts := strings.Split(val, "=")
		if len(parts) < 2 {
			continue // malformed entry
		}
		key := parts[0]
		if key != "PATH" {
			continue
		}
		val := strings.Join(parts[1:], "=")
		env[idx] = fmt.Sprintf("%s=%s:%s", key, extra, val)
		found = true
	}
	if !found {
		const busyboxDefaultPATH = "/usr/local/sbin:/sbin:/usr/sbin:/usr/local/bin:/bin:/usr/bin"
		env = append(env, fmt.Sprintf("PATH=%s:%s", extra, busyboxDefaultPATH))
	}
	return env
}
