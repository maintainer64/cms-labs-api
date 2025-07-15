package topology

import (
	"testing"

	"github.com/goccy/go-json"
)

func TestConvert(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *Topology
		wantErr  bool
	}{
		{
			name: "basic topology conversion",
			input: `
apiVersion: clabernetes.containerlab.dev/v1alpha1
kind: Topology
spec:
  definition:
    containerlab: |
      name: test
      topology:
        nodes:
          node1:
            kind: nokia_srlinux
            labels:
              flow_icon: switch
              flow_label: S1
          node2:
            kind: linux
            labels:
              flow_icon: linux
              flow_label: PC1
        links:
          - endpoints: [node1:eth1, node2:eth1]
        defaults:
          labels:
            flow_direction: LR
`,
			expected: &Topology{
				Direction: "LR",
				Nodes: []TopologiesNode{
					{
						ID:    "node1",
						Label: "S1",
						Type:  "nokia_srlinux",
						Icon:  "switch",
					},
					{
						ID:    "node2",
						Label: "PC1",
						Type:  "linux",
						Icon:  "linux",
					},
				},
				Edges: []TopologiesEdge{
					{
						ID:     "edge-0",
						Type:   "default",
						Source: "node1",
						Target: "node2",
						Data: TopologiesEdgeData{
							Source: TopologiesEdgeDataItem{
								ID:    "node1",
								Name:  "eth1",
								Type:  "nokia_srlinux",
								Label: "S1",
							},
							Target: TopologiesEdgeDataItem{
								ID:    "node2",
								Name:  "eth1",
								Type:  "linux",
								Label: "PC1",
							},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "empty topology",
			input: `
apiVersion: clabernetes.containerlab.dev/v1alpha1
kind: Topology
spec:
  definition:
    containerlab: |
      name: test
      topology:
        defaults:
          labels:
            flow_direction: TB
`,
			expected: &Topology{
				Direction: "TB",
				Nodes:     []TopologiesNode{},
				Edges:     []TopologiesEdge{},
			},
			wantErr: false,
		},
		{
			name: "invalid YAML",
			input: `
invalid yaml: {]
`,
			expected: nil,
			wantErr:  true,
		},
		{
			name: "missing containerlab definition",
			input: `
apiVersion: clabernetes.containerlab.dev/v1alpha1
kind: Topology
spec:
  definition: {}
`,
			expected: nil,
			wantErr:  true,
		},
		{
			name: "link with single endpoint",
			input: `
apiVersion: clabernetes.containerlab.dev/v1alpha1
kind: Topology
spec:
  definition:
    containerlab: |
      name: test
      topology:
        nodes:
          node1:
            kind: nokia_srlinux
            labels:
              flow_icon: switch
              flow_label: S1
        links:
          - endpoints: [node1:eth1]
        defaults:
          labels:
            flow_direction: LR
`,
			expected: &Topology{
				Direction: "LR",
				Nodes: []TopologiesNode{
					{
						ID:    "node1",
						Label: "S1",
						Type:  "nokia_srlinux",
						Icon:  "switch",
					},
				},
				Edges: []TopologiesEdge{},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Convert([]byte(tt.input))
			if (err != nil) != tt.wantErr {
				t.Errorf("Convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			// Проверяем направление
			if got.Direction != tt.expected.Direction {
				t.Errorf("Direction = %v, want %v", got.Direction, tt.expected.Direction)
			}

			// Проверяем узлы
			if len(got.Nodes) != len(tt.expected.Nodes) {
				t.Errorf("Nodes count = %v, want %v", len(got.Nodes), len(tt.expected.Nodes))
			} else {
				for i, node := range got.Nodes {
					expectedNode := tt.expected.Nodes[i]
					if node != expectedNode {
						t.Errorf("Node[%d] = %v, want %v", i, node, expectedNode)
					}
				}
			}

			// Проверяем связи
			if len(got.Edges) != len(tt.expected.Edges) {
				t.Errorf("Edges count = %v, want %v", len(got.Edges), len(tt.expected.Edges))
			} else {
				for i, edge := range got.Edges {
					expectedEdge := tt.expected.Edges[i]
					if edge.ID != expectedEdge.ID ||
						edge.Type != expectedEdge.Type ||
						edge.Source != expectedEdge.Source ||
						edge.Target != expectedEdge.Target {
						t.Errorf("Edge[%d] basic fields mismatch, got %v, want %v", i, edge, expectedEdge)
					}

					// Сравниваем данные связи
					if edge.Data.Source != expectedEdge.Data.Source ||
						edge.Data.Target != expectedEdge.Data.Target {
						gotData, _ := json.Marshal(edge.Data)
						wantData, _ := json.Marshal(expectedEdge.Data)
						t.Errorf("Edge[%d].Data = %s, want %s", i, string(gotData), string(wantData))
					}
				}
			}
		})
	}
}

func TestParseEndpoint(t *testing.T) {
	tests := []struct {
		input    string
		expected endpoint
	}{
		{
			input: "node1:eth1",
			expected: endpoint{
				Node:      "node1",
				Interface: "eth1",
			},
		},
		{
			input: "node1",
			expected: endpoint{
				Node: "node1",
			},
		},
		{
			input: "node1:eth1:invalid", // некорректный формат
			expected: endpoint{
				Node: "node1:eth1:invalid",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := parseEndpoint(tt.input)
			if got != tt.expected {
				t.Errorf("parseEndpoint() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestFindNode(t *testing.T) {
	nodes := []TopologiesNode{
		{ID: "node1", Label: "Node 1"},
		{ID: "node2", Label: "Node 2"},
	}

	tests := []struct {
		name     string
		nodeID   string
		expected *TopologiesNode
	}{
		{
			name:     "existing node",
			nodeID:   "node1",
			expected: &nodes[0],
		},
		{
			name:     "non-existing node",
			nodeID:   "node3",
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findNode(nodes, tt.nodeID)
			if (got == nil) != (tt.expected == nil) {
				t.Errorf("findNode() = %v, want %v", got, tt.expected)
			} else if got != nil && *got != *tt.expected {
				t.Errorf("findNode() = %v, want %v", *got, *tt.expected)
			}
		})
	}
}
