package topology

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

// Convert преобразует YAML ContainerLab в выходную структуру
func Convert(yamlData []byte) (*Topology, error) {
	// Парсим основной YAML
	var clab struct {
		Spec struct {
			Definition struct {
				ContainerLab string `yaml:"containerlab"`
			} `yaml:"definition"`
		} `yaml:"spec"`
	}

	if err := yaml.Unmarshal(yamlData, &clab); err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML: %v", err)
	}

	// Проверяем наличие containerlab определения
	if clab.Spec.Definition.ContainerLab == "" {
		return nil, fmt.Errorf("containerlab definition is empty")
	}

	// Парсим вложенный YAML containerlab
	var topology struct {
		Topology struct {
			Nodes map[string]struct {
				Labels map[string]string `yaml:"labels"`
				Kind   string            `yaml:"kind"`
			} `yaml:"nodes"`
			Links []struct {
				Endpoints []string `yaml:"endpoints"`
			} `yaml:"links"`
			Defaults struct {
				Labels struct {
					FlowDirection string `yaml:"flow_direction"`
				} `yaml:"labels"`
			} `yaml:"defaults"`
		} `yaml:"topology"`
	}

	if err := yaml.Unmarshal([]byte(clab.Spec.Definition.ContainerLab), &topology); err != nil {
		return nil, fmt.Errorf("failed to unmarshal containerlab YAML: %v", err)
	}

	// Создаем выходную структуру
	output := &Topology{
		Direction: strings.ToUpper(topology.Topology.Defaults.Labels.FlowDirection),
	}

	// Обрабатываем узлы
	for nodeName, nodeData := range topology.Topology.Nodes {
		output.Nodes = append(output.Nodes, TopologiesNode{
			ID:    nodeName,
			Label: nodeData.Labels["flow_label"],
			Type:  nodeData.Kind,
			Icon:  nodeData.Labels["flow_icon"],
		})
	}

	// Обрабатываем связи
	for i, link := range topology.Topology.Links {
		if len(link.Endpoints) < 2 {
			continue
		}

		source, target := parseEndpoint(link.Endpoints[0]), parseEndpoint(link.Endpoints[1])
		sourceNode := findNode(output.Nodes, source.Node)
		targetNode := findNode(output.Nodes, target.Node)

		edgeID := fmt.Sprintf("edge-%d", i)
		edgeData := TopologiesEdgeData{}
		if sourceNode == nil && targetNode == nil {
			continue
		}
		if sourceNode != nil {
			edgeData.Source = TopologiesEdgeDataItem{
				ID:    source.Node,
				Name:  source.Interface,
				Type:  sourceNode.Type,
				Label: sourceNode.Label,
			}
		}
		if targetNode != nil {
			edgeData.Target = TopologiesEdgeDataItem{
				ID:    target.Node,
				Name:  target.Interface,
				Type:  targetNode.Type,
				Label: targetNode.Label,
			}
		}

		output.Edges = append(output.Edges, TopologiesEdge{
			ID:     edgeID,
			Type:   "default",
			Source: source.Node,
			Target: target.Node,
			Data:   edgeData,
		})
	}

	return output, nil
}

func parseEndpoint(ep string) endpoint {
	parts := strings.Split(ep, ":")
	if len(parts) != 2 {
		return endpoint{Node: ep}
	}
	return endpoint{
		Node:      parts[0],
		Interface: parts[1],
	}
}

func findNode(nodes []TopologiesNode, name string) *TopologiesNode {
	for _, node := range nodes {
		if node.ID == name {
			return &node
		}
	}
	return nil
}
