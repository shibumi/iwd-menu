package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	dbus "github.com/godbus/dbus/v5"
	"github.com/shibumi/iwd"
)

func main() {
	conn, err := dbus.SystemBus()
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	i := iwd.New(conn)
	openNetworks := map[int]iwd.Network{}
	var index int = 0
	for _, network := range i.Networks {
		if network.Type == "open" {
			fmt.Println(index, "  ", network.Name)
			openNetworks[index] = network
			index++
		}
	}
	fmt.Println("---")
	fmt.Print("Connect > ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	target, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}
	n := openNetworks[target]
	n.Connect(conn)
}
