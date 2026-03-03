'use client';

import { Edge, Node } from '@xyflow/react';
import dagre from 'dagre';

export type LayoutDirection = 'TB' | 'LR';

interface LayoutOptions {
  direction?: LayoutDirection;
  nodeWidth?: number;
  nodeHeight?: number;
  nodesep?: number;
  ranksep?: number;
}

export const getLayoutedElementsTargets = (
  nodes: Node[],
  edges: Edge[],
  options: LayoutOptions = {}
): { nodes: Node[]; edges: Edge[] } => {
  const { direction = 'LR', nodeWidth = 200, nodeHeight = 80, nodesep = 80, ranksep = 120 } = options;

  const dagreGraph = new dagre.graphlib.Graph();
  dagreGraph.setDefaultEdgeLabel(() => ({}));
  dagreGraph.setGraph({
    rankdir: direction,
    nodesep,
    ranksep,
    marginx: 40,
    marginy: 40
  });

  // Добавляем узлы
  nodes.forEach((node) => {
    dagreGraph.setNode(node.id, {
      width: node.measured?.width || nodeWidth,
      height: node.measured?.height || nodeHeight
    });
  });

  // Добавляем рёбра
  edges.forEach((edge) => {
    dagreGraph.setEdge(edge.source, edge.target);
  });

  // Запускаем layout
  dagre.layout(dagreGraph);

  // Обновляем позиции узлов
  const layoutedNodes = nodes.map((node) => {
    const nodeWithPosition = dagreGraph.node(node.id);
    const width = node.measured?.width || nodeWidth;
    const height = node.measured?.height || nodeHeight;

    return {
      ...node,
      position: {
        x: nodeWithPosition.x - width / 2,
        y: nodeWithPosition.y - height / 2
      }
    };
  });

  return { nodes: layoutedNodes, edges };
};
