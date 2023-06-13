package infra

import "math"

type Infra struct {
	Nodes            int
	ConvergesPerDay  int
	ConvergeSizeMB   float64
	CompliancePerDay int
	ComplianceSizeMB float64
	RetentionDays    int

	ConvergeIndicesPerDay   int
	ComplianceIndicesPerDay int
	PrimaryShards           int
	ReplicaShards           int
	HeapPerNode             int
}

type DataTotals struct {
	PerDayGB    float64
	ShardSizeGB float64
	// Raw data does not include replication
	TotalRawGB float64
	// Raw data with 25% overhead for growth, does not include replication
	TotalDataGB float64
}

func CalculateDataSize(nodes int, perDay int, sizeMB float64, primary_shards int, retention int) DataTotals {
	daily_data := float64(nodes) * float64(perDay) * sizeMB / 1024
	total_raw := daily_data * float64(retention)
	shard_size := daily_data / float64(primary_shards)

	ds := DataTotals{
		PerDayGB:    daily_data,
		ShardSizeGB: shard_size,
		TotalRawGB:  total_raw,
		TotalDataGB: total_raw * 1.25,
	}

	return ds
}

func New() Infra {
	d := Infra{
		Nodes:         10000,
		RetentionDays: 30,

		ConvergesPerDay:  24,
		ConvergeSizeMB:   0.3,
		CompliancePerDay: 1,
		ComplianceSizeMB: 4,

		ConvergeIndicesPerDay:   2,
		ComplianceIndicesPerDay: 2,
		PrimaryShards:           5,
		ReplicaShards:           1,
		HeapPerNode:             32,
	}

	return d
}

// Calculate the number of allowed shards based on heap size
// ElasticSearch and AWS documentation recommends 25 shards/GB heap
// in practice we find when we hit 80 shards/GB Heap is when we begin to see
// performance issues with OS
//
// To ensure adequate wiggle room using 50 shards/GB Heap
func (i Infra) shardsPerNode() int {
	return i.HeapPerNode * 50
}

// Calculate the total number of shards for the compliance indices for the retention period
func (i Infra) ComplianceShards() int {
	return i.ComplianceIndicesPerDay * i.RetentionDays * (i.PrimaryShards + i.PrimaryShards*i.ReplicaShards)
}

// Calculate the total number of shards for the converge indices for the retention period
func (i Infra) ConvergeShards() int {
	return i.ConvergeIndicesPerDay * i.RetentionDays * (i.PrimaryShards + i.PrimaryShards*i.ReplicaShards)
}

// Sum the total number of converge and compliance shards
func (i Infra) TotalShards() int {
	return i.ConvergeShards() + i.ComplianceShards()
}

// Based on the number of shards calculate how many nodes we need
func (i Infra) OSNodesForShards(s int) float64 {
	nodes := math.Ceil(float64(s) / float64(i.shardsPerNode()))

	return nodes
}

// Calculate number of shards needed to ensure we limit shard size to no more than
// 50gb per shard
//
// Typically we'll find that when the shard size goes above 30gb the search performance
// will start to decrease but with a write heavy environment like we have in Automate going
// up to 50gb per shard is better for write performance.
func calcShards(data_per_day float64) int {
	// Number of 50gb shards
	shards := math.Ceil(data_per_day / 50)

	return int(shards)
}

// Calculate how many primary shards we need for compliance data to maintain < 30gb shard size
func (i Infra) calcPrimaryComplianceShards() int {
	return calcShards(i.ComplianceData().PerDayGB)
}

// Calculate total number of shards for compliance indices with retention period
func (i Infra) CalculatedComplianceShardsTotal() int {
	shards := i.calcPrimaryComplianceShards()
	total_shards := shards + (shards * i.ReplicaShards)
	return i.ComplianceIndicesPerDay * total_shards * i.RetentionDays
}

// Calculate how many primary shards we need for converge data to maintain < 30gb shard size
func (i Infra) calcPrimaryConvergeShards() int {
	return calcShards(i.ConvergeData().PerDayGB)
}

// Calculate total number of shards for converge indices with retention period
func (i Infra) CalculatedConvergeShardsTotal() int {
	shards := i.calcPrimaryConvergeShards()
	total_shards := shards + (shards * i.ReplicaShards)
	return i.ConvergeIndicesPerDay * total_shards * i.RetentionDays
}

// Determine the number of OpenSearch nodes based on the default number of
// primary and replica shards
func (i Infra) OSNodes() int {
	nodes := i.OSNodesForShards(i.TotalShards())

	// Make sure we always have an odd number of nodes
	if int(nodes)%2 == 0 {
		nodes += 1
	}

	return int(math.Max(nodes, 3))
}

// If the default shard size doesn't match the number of shards we calculated then
// the indicies won't have the optimal size and will likely be too large.
func (i Infra) IsOptimalShardCount() bool {
	return i.OSNodes() == i.OSNodesOptimalShards()
}

// Determine the number of OpenSearch nodes by calculating the optimal number of shards
// necessary to keep shard size limited to the optimal size.
func (i Infra) OSNodesOptimalShards() int {
	nodes := i.OSNodesForShards(i.CalculatedConvergeShardsTotal() + i.CalculatedComplianceShardsTotal())

	// Make sure we always have an odd number of nodes
	if int(nodes)%2 == 0 {
		nodes += 1
	}

	return int(math.Max(nodes, 3))
}

// Calculate the amount of converge data generated by the chef-client nodes
func (i Infra) ConvergeData() DataTotals {
	return CalculateDataSize(i.Nodes, i.ConvergesPerDay, i.ConvergeSizeMB, i.PrimaryShards, i.RetentionDays)
}

// Calculate the amount of compliance data generated by the chef-client nodes
func (i Infra) ComplianceData() DataTotals {
	return CalculateDataSize(i.Nodes, i.CompliancePerDay, i.ComplianceSizeMB, i.PrimaryShards, i.RetentionDays)
}
