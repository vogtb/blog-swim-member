package main

import (
	"log"
	"os"
	"time"

	"github.com/hashicorp/memberlist"
)

func main() {
	log.Println("Starting node.")
	hostname, _ := os.Hostname()

	// Hard code hostname and IP address
	var hostMap = make(map[string]string)
	hostMap["node01"] = "192.168.50.6"
	hostMap["node02"] = "192.168.50.10"
	hostMap["node03"] = "192.168.50.7"
	hostMap["node04"] = "192.168.50.11"

	config := memberlist.DefaultLocalConfig()
	config.BindAddr = hostMap[hostname]
	list, err := memberlist.Create(config)
	if err != nil {
		log.Println("Failed to create memberlist: " + err.Error())
	}

	// Join
	clusterCount, err := list.Join([]string{hostMap["node01"]})
	log.Println("Joing cluster of size", string(clusterCount))
	if err != nil {
		log.Println("Failed to join cluster: " + err.Error())
	}

	// Every two seconds, log the members list
	for {
		time.Sleep(time.Second * 2)
		log.Println("Members:")
		for _, member := range list.Members() {
			log.Printf("  %s %s\n", member.Name, member.Addr)
		}
	}

}
