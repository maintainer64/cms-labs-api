import {Edge as RFEdge, Node as RFNode} from '@xyflow/react';
import {topology_TopologiesEdge, topology_TopologiesNode} from "@/helpers/api";
import {CSSProperties} from "react";


export type RFNodeTopology = RFNode<topology_TopologiesNode>;

export type RFEdgeTopology = RFEdge<topology_TopologiesEdge>;


export const defaultStyleNodes: CSSProperties = {
    width: 'auto',
    height: 'auto',
    background: 'none',
    border: 'none',
    padding: 0,
    margin: 0,
    cursor: 'pointer',
    zIndex: 10,
    boxShadow: 'none',
};

export const defaultStyleEdge: CSSProperties = {
    stroke: '#0066AAFF',
    strokeWidth: 2
};