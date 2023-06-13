package infra

import "testing"

func TestInfra_OpenSearchNodes(t *testing.T) {
	type fields struct {
		Nodes                   int
		ConvergesPerDay         int
		ConvergeSizeMB          float64
		CompliancePerDay        int
		ComplianceSizeMB        float64
		RetentionDays           int
		ConvergeIndicesPerDay   int
		ComplianceIndicesPerDay int
		PrimaryShards           int
		ReplicaShards           int
		HeapPerNode             int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "10k_nodes_30_days",
			fields: fields{
				Nodes:                   10000,
				ConvergesPerDay:         24,
				ConvergeSizeMB:          0.5,
				CompliancePerDay:        1,
				ComplianceSizeMB:        4,
				RetentionDays:           30,
				ConvergeIndicesPerDay:   2,
				ComplianceIndicesPerDay: 2,
				PrimaryShards:           5,
				ReplicaShards:           1,
				HeapPerNode:             32,
			},
			want: 3,
		},
		{
			name: "10k_nodes_90_days",
			fields: fields{
				Nodes:                   10000,
				ConvergesPerDay:         24,
				ConvergeSizeMB:          0.5,
				CompliancePerDay:        1,
				ComplianceSizeMB:        4,
				RetentionDays:           90,
				ConvergeIndicesPerDay:   2,
				ComplianceIndicesPerDay: 2,
				PrimaryShards:           5,
				ReplicaShards:           1,
				HeapPerNode:             32,
			},
			want: 3,
		},
		{
			name: "10k_nodes_180_days",
			fields: fields{
				Nodes:                   10000,
				ConvergesPerDay:         24,
				ConvergeSizeMB:          0.5,
				CompliancePerDay:        1,
				ComplianceSizeMB:        4,
				RetentionDays:           180,
				ConvergeIndicesPerDay:   2,
				ComplianceIndicesPerDay: 2,
				PrimaryShards:           5,
				ReplicaShards:           1,
				HeapPerNode:             32,
			},
			want: 5,
		},
		{
			name: "10k_nodes_365_days",
			fields: fields{
				Nodes:                   10000,
				ConvergesPerDay:         24,
				ConvergeSizeMB:          0.5,
				CompliancePerDay:        1,
				ComplianceSizeMB:        4,
				RetentionDays:           365,
				ConvergeIndicesPerDay:   2,
				ComplianceIndicesPerDay: 2,
				PrimaryShards:           5,
				ReplicaShards:           1,
				HeapPerNode:             32,
			},
			want: 7,
		},

		{
			name: "10k_nodes_365_days_min_shards",
			fields: fields{
				Nodes:                   10000,
				ConvergesPerDay:         24,
				ConvergeSizeMB:          0.5,
				CompliancePerDay:        1,
				ComplianceSizeMB:        4,
				RetentionDays:           365,
				ConvergeIndicesPerDay:   2,
				ComplianceIndicesPerDay: 2,
				PrimaryShards:           2,
				ReplicaShards:           1,
				HeapPerNode:             32,
			},
			want: 3,
		},
		{
			name: "50k_nodes_90_days",
			fields: fields{
				Nodes:                   50000,
				ConvergesPerDay:         24,
				ConvergeSizeMB:          0.5,
				CompliancePerDay:        1,
				ComplianceSizeMB:        4,
				RetentionDays:           90,
				ConvergeIndicesPerDay:   2,
				ComplianceIndicesPerDay: 2,
				PrimaryShards:           5,
				ReplicaShards:           1,
				HeapPerNode:             32,
			},
			want: 3,
		},
		{
			name: "50k_nodes_180_days",
			fields: fields{
				Nodes:                   50000,
				ConvergesPerDay:         24,
				ConvergeSizeMB:          0.5,
				CompliancePerDay:        1,
				ComplianceSizeMB:        4,
				RetentionDays:           180,
				ConvergeIndicesPerDay:   2,
				ComplianceIndicesPerDay: 2,
				PrimaryShards:           5,
				ReplicaShards:           1,
				HeapPerNode:             32,
			},
			want: 5,
		},
		{
			name: "50k_nodes_365_days",
			fields: fields{
				Nodes:                   50000,
				ConvergesPerDay:         24,
				ConvergeSizeMB:          0.5,
				CompliancePerDay:        1,
				ComplianceSizeMB:        4,
				RetentionDays:           365,
				ConvergeIndicesPerDay:   2,
				ComplianceIndicesPerDay: 2,
				PrimaryShards:           5,
				ReplicaShards:           1,
				HeapPerNode:             32,
			},
			want: 7,
		},
		{
			name: "50k_nodes_365_days_min_shards",
			fields: fields{
				Nodes:                   50000,
				ConvergesPerDay:         24,
				ConvergeSizeMB:          0.5,
				CompliancePerDay:        1,
				ComplianceSizeMB:        4,
				RetentionDays:           365,
				ConvergeIndicesPerDay:   2,
				ComplianceIndicesPerDay: 2,
				PrimaryShards:           2,
				ReplicaShards:           1,
				HeapPerNode:             32,
			},
			want: 3,
		},
		{
			name: "100k_nodes_90_days",
			fields: fields{
				Nodes:                   100000,
				ConvergesPerDay:         24,
				ConvergeSizeMB:          0.5,
				CompliancePerDay:        1,
				ComplianceSizeMB:        4,
				RetentionDays:           90,
				ConvergeIndicesPerDay:   2,
				ComplianceIndicesPerDay: 2,
				PrimaryShards:           5,
				ReplicaShards:           1,
				HeapPerNode:             32,
			},
			want: 3,
		},
		{
			name: "100k_nodes_180_days",
			fields: fields{
				Nodes:                   100000,
				ConvergesPerDay:         24,
				ConvergeSizeMB:          0.5,
				CompliancePerDay:        1,
				ComplianceSizeMB:        4,
				RetentionDays:           180,
				ConvergeIndicesPerDay:   2,
				ComplianceIndicesPerDay: 2,
				PrimaryShards:           5,
				ReplicaShards:           1,
				HeapPerNode:             32,
			},
			want: 5,
		},
		{
			name: "100k_nodes_365_days",
			fields: fields{
				Nodes:                   100000,
				ConvergesPerDay:         24,
				ConvergeSizeMB:          0.5,
				CompliancePerDay:        1,
				ComplianceSizeMB:        4,
				RetentionDays:           365,
				ConvergeIndicesPerDay:   2,
				ComplianceIndicesPerDay: 2,
				PrimaryShards:           5,
				ReplicaShards:           1,
				HeapPerNode:             32,
			},
			want: 7,
		},
		{
			name: "100k_nodes_365_days_min_shards",
			fields: fields{
				Nodes:                   100000,
				ConvergesPerDay:         24,
				ConvergeSizeMB:          0.5,
				CompliancePerDay:        1,
				ComplianceSizeMB:        4,
				RetentionDays:           365,
				ConvergeIndicesPerDay:   2,
				ComplianceIndicesPerDay: 2,
				PrimaryShards:           2,
				ReplicaShards:           1,
				HeapPerNode:             32,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := Infra{
				Nodes:                   tt.fields.Nodes,
				ConvergesPerDay:         tt.fields.ConvergesPerDay,
				ConvergeSizeMB:          tt.fields.ConvergeSizeMB,
				CompliancePerDay:        tt.fields.CompliancePerDay,
				ComplianceSizeMB:        tt.fields.ComplianceSizeMB,
				RetentionDays:           tt.fields.RetentionDays,
				ConvergeIndicesPerDay:   tt.fields.ConvergeIndicesPerDay,
				ComplianceIndicesPerDay: tt.fields.ComplianceIndicesPerDay,
				PrimaryShards:           tt.fields.PrimaryShards,
				ReplicaShards:           tt.fields.ReplicaShards,
				HeapPerNode:             tt.fields.HeapPerNode,
			}
			got := i.OSNodes()
			if got != tt.want {
				t.Errorf("Infra.OSNodes() = %v, want %v", got, tt.want)
			}
			if !i.IsOptimalShardCount() {
				t.Errorf("Infra.OSNodes() does not have optimal shard, got %v and should be %v nodes based on optimal shard count", got, i.OSNodesCalculatedShards())
			}
		})
	}
}
