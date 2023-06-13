package main

import (
	"github.com/chef/automate_ha_calculator/infra"
	log "github.com/sirupsen/logrus"
)

func printInfo(i infra.Infra) {
	log.Infof("Context: %+v\n", i)

	log.Infof("Converge Data: %+v\n", i.ConvergeData())
	log.Infof("Compliance Data: %+v\n", i.ComplianceData())

	log.Infof("Total Shards: %d\n", i.TotalShards())
	log.Infof("Optimal Shard Size?: %v\n", i.IsOptimalShardCount())
	log.Infof("OS Nodes: %d\n", i.OSNodes())
	log.Infof("OS Nodes w/ Optimal Shard Size: %d\n", i.OSNodesCalculatedShards())
}

func main() {
	i := infra.New()
	printInfo(i)

	// i = infra.New()
	i.Nodes = 100000
	i.RetentionDays = 90
	printInfo(i)
}
