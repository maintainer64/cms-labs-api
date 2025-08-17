package topology

// TopologiesNodeDataItem представляет данные о ноде
type TopologiesNodeDataItem struct {
	ID    string `json:"id"`
	Label string `json:"label"`
	Type  string `json:"type"`
	// Icon enum:cloud,router,server,switch,desktop
	Icon    string `json:"icon"`
	Kind    string `json:"kind"`
	Startup string `json:"startup"`
}

// TopologiesNode представляет узел топологии
type TopologiesNode struct {
	ID    string `json:"id"`
	Label string `json:"label"`
	Type  string `json:"type"`
	// Icon enum:cloud,router,server,switch,desktop
	Icon string                 `json:"icon"`
	Data TopologiesNodeDataItem `json:"data"`
}

// TopologiesEdgeDataItem представляет данные о конечной точке связи
type TopologiesEdgeDataItem struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Label string `json:"label"`
}

// TopologiesEdgeData содержит данные о связи
type TopologiesEdgeData struct {
	Source TopologiesEdgeDataItem `json:"source"`
	Target TopologiesEdgeDataItem `json:"target"`
}

// TopologiesEdge представляет связь между узлами
type TopologiesEdge struct {
	ID     string             `json:"id"`
	Type   string             `json:"type"`
	Source string             `json:"source"`
	Target string             `json:"target"`
	Data   TopologiesEdgeData `json:"data"`
}

// Topology содержит полное описание топологии
type Topology struct {
	Nodes     []TopologiesNode `json:"nodes"`
	Edges     []TopologiesEdge `json:"edges"`
	Direction string           `json:"direction"`
}

func (t *Topology) GetNodeByID(id string) *TopologiesNode {
	for _, node := range t.Nodes {
		if node.ID == id {
			return &node
		}
	}
	return nil
}

// ContainerLabTopology представляет структуру ContainerLab топологии
type ContainerLabTopology struct {
	Definition struct {
		ContainerLab string `yaml:"containerlab"`
	} `yaml:"definition"`
}

type endpoint struct {
	Node      string
	Interface string
}
