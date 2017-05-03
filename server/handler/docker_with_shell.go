package handler

import (
	"log"
	"os/exec"

	"../api"
)

func reloadHandler(msg api.HubMessage) {
	log.Println("received message to reload ...")
	out, err := exec.Command("../reload.sh").Output()
	if err != nil {
		log.Println("ERROR EXECUTING COMMAND IN RELOAD HANDLER!!")
		log.Println(err)
		return
	}
	log.Println("output of reload.sh is", string(out))
}
