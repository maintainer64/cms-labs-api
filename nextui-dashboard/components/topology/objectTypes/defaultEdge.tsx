import React from 'react';
import { Edge, EdgeProps, getStraightPath, useInternalNode } from '@xyflow/react';
import { calculateDimensionsLabel } from '@/components/topology/objectTypes/lineObjectUtils';
import { getEdgeParams, getLabelPosition, getReverseLabelPosition } from '@/components/topology/objectTypes/edgeUtils';
import { EdgeData } from '@/components/topology/objectTypes/types';

const EdgeLabel = ({ x, y, label }: { x: number; y: number; label?: string }) => {
  if (!label) return null;

  const dimensions = calculateDimensionsLabel(label);
  const textWidth = dimensions.width;
  const textHeight = dimensions.height;

  return (
    <g transform={`translate(${x}, ${y})`}>
      {/* Фоновая подложка */}
      <rect
        x={-textWidth / 2}
        y={-textHeight / 2}
        width={textWidth}
        height={textHeight}
        rx='8'
        ry='8'
        fill='#e6f5ff'
        stroke='#346789'
        strokeWidth='2'
        filter='url(#shadow)'
      />

      {/* Текст */}
      <text
        style={{
          fontFamily: 'helvetica, sans-serif',
          fontSize: '12px',
          fontWeight: 'bold',
          fill: '#333'
        }}
        dominantBaseline='middle'
        textAnchor='middle'
      >
        {label}
      </text>

      {/* Фильтр для тени (если нужен) */}
      <defs>
        <filter id='shadow' x='-20%' y='-20%' width='140%' height='140%'>
          <feDropShadow dx='1' dy='1' stdDeviation='1' floodColor='rgba(0,0,0,0.2)' />
        </filter>
      </defs>
    </g>
  );
};

// @ts-ignore
export const edgeTypeDefault = ({ id, source, target, markerEnd, style, data }: EdgeProps<Edge<EdgeData>>) => {
  const targetNode = useInternalNode(target);
  const sourceNode = useInternalNode(source);

  if (!sourceNode || !targetNode) {
    return null;
  }

  const { sx, sy, tx, ty } = getEdgeParams(sourceNode, targetNode);
  const [edgePath] = getStraightPath({
    sourceX: sx,
    sourceY: sy,
    targetX: tx,
    targetY: ty
  });

  // Позиции лейблов (20% от длины линии с каждого конца)
  const sourceLabelPos = getLabelPosition(sx, sy, tx, ty);
  const targetLabelPos = getReverseLabelPosition(sx, sy, tx, ty);
  // @ts-ignore
  const sourceData = data?.data?.source;
  // @ts-ignore
  const targetData = data?.data?.target;
  return (
    <>
      <path id={id} className='react-flow__edge-path' d={edgePath} markerEnd={markerEnd} style={style} />
      {/* Лейбл для source */}
      <EdgeLabel x={sourceLabelPos.x} y={sourceLabelPos.y} label={sourceData?.name || sourceData?.label} />
      {/* Лейбл для target */}
      <EdgeLabel x={targetLabelPos.x} y={targetLabelPos.y} label={targetData?.name || targetData?.label} />
    </>
  );
};
