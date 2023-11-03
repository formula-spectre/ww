package main

import (
	"log"
	"os/exec"
	"Syscall"
)

func main() {
	//cmd := exec.Command("mount", "/dev/sdc1", "/media/sdc1")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Mount successful")
}


