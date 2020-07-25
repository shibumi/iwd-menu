package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	dbus "github.com/godbus/dbus/v5"
	"github.com/shibumi/iwd"
)

var errNoActiveStation error = errors.New("no active station for scanning available. Is your wifi adapter plugged in?")

// getActiveStation returns the connected wifi station on success + err=nil.
// On failure it will return an empty station object + an errNoActiveStation error.
// A station is a connected wifi device like "wlan0".
func getActiveStation(daemon iwd.Iwd) (activeStation iwd.Station, err error) {
	for _, station := range daemon.Stations {
		if station.State == "connected" {
			return activeStation, nil
		}
	}
	return iwd.Station{}, errNoActiveStation
}

func main() {
	conn, err := dbus.SystemBus()
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	daemon := iwd.New(conn)
	activeStation, err := getActiveStation(daemon)
	if err != nil {
		log.Fatalln(err)
	}
	activeStation.Scan(conn)
	openNetworks := map[int]iwd.Network{}
	var index int = 0
	for _, network := range daemon.Networks {
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
