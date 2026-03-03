import { Edge as RFEdge, Node as RFNode } from '@xyflow/react';
import { TopologyTopologiesEdge, TopologyTopologiesNode } from '@/helpers/api';
import { CSSProperties } from 'react';
import { CamelCasedPropertiesDeep } from 'type-fest';

export type RFNodeTopology = RFNode<CamelCasedPropertiesDeep<TopologyTopologiesNode>>;

export type RFEdgeTopology = RFEdge<CamelCasedPropertiesDeep<TopologyTopologiesEdge>>;

export const defaultStyleNodes: CSSProperties = {
  width: 'auto',
  height: 'auto',
  background: 'none',
  border: 'none',
  padding: 0,
  margin: 0,
  cursor: 'pointer',
  zIndex: 10,
  boxShadow: 'none'
};

export const defaultStyleEdge: CSSProperties = {
  stroke: '#0066AAFF',
  strokeWidth: 2
};
