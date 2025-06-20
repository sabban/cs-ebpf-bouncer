// SPDX-License-Identifier: MIT

package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/sabban/cs-ebpf-bouncer/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
