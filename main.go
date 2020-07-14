package main

import (
	"fmt"

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
	index := 0
	for _, network := range i.Networks {
		if network.Type == "open" {
			fmt.Println(index, "  ", network.Name)
			//path := string(network.Device) + "/" + hex.EncodeToString([]byte(network.Name)) + "_" + network.Type
			//device := conn.Object("net.connman.iwd", dbus.ObjectPath(path))
			//device.Call("net.connman.iwd.Network.Connect", 0)
			index++
		}
	}
	// all := conn.Object("net.connman.iwd", "/")
	// var managed map[dbus.ObjectPath]map[string]map[string]dbus.Variant
	// err = all.Call("org.freedesktop.DBus.ObjectManager.GetManagedObjects", 0).Store(&managed)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(managed)
	// iwdDevice := conn.Object("net.connman.iwd", "/net/connman/iwd/0/4")
	// var networks []interface{}
	// iwdDevice.Call("net.connman.iwd.Station.Scan", 0)
	// err = iwdDevice.Call("net.connman.iwd.Station.GetOrderedNetworks", 0).Store(&networks)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(networks)
	// for _, network := range networks {
	// 	n := conn.Object("net.connman.iwd", network.ObjectPath)
	// 	fmt.Println(n.GetProperty("net.connman.iwd.Network.Name"))
	// }
	// //var list []string
	// //call := iwdDevice.Call("net.connman.iwd.Station.GetOrderedNetworks", 0, "")
	// //fmt.Println(call)
}
