'use client';

import React, { memo } from 'react';
import { Handle, NodeProps, Position } from '@xyflow/react';
import { ArrowUpRight, ExternalLink } from 'lucide-react';
import {RoutesLocation} from "@/components/routes";

interface TargetsNodeData {
  label: string;
  nodeId: string;
  description?: string;
  type: string;
  icon: React.ReactNode;
  isActive: boolean;
  isHighlighted?: boolean;
  links?: Array<{ type: string; value: string }>;
  colors: {
    bg: string;
    border: string;
    text: string;
    dimmed: string;
  };
}

export const TargetsNode = memo(({ data, selected }: NodeProps) => {
  const nodeData = data as unknown as TargetsNodeData;
  const { label, description, icon, isActive, isHighlighted, links, colors, nodeId } = nodeData;
  const linksFiltered = links?.filter((link) => link.value.startsWith('http') || link.type === 'ip') || [];

  return (
    <div
      className={`
                relative px-4 py-3 rounded-lg shadow-md border-2 min-w-[180px] max-w-[220px]
                transition-all duration-300 cursor-pointer
                ${selected || isHighlighted ? 'ring-2 ring-offset-2 ring-blue-500 shadow-lg' : ''}
                hover:shadow-lg hover:scale-[1.02]
            `}
      style={{
        backgroundColor: isActive ? colors.bg : '#f9fafb',
        borderColor: isActive ? colors.border : '#d1d5db'
      }}
    >
      {/* @ts-ignore */}
      <Handle
        type='target'
        position={Position.Top}
        className='!w-3 !h-3 !border-2'
        style={{
          backgroundColor: isActive ? colors.border : '#d1d5db',
          borderColor: '#fff'
        }}
      />

      <div className='flex items-start gap-2'>
        <div
          className='flex-shrink-0 p-1.5 rounded'
          style={{
            backgroundColor: isActive ? colors.border : '#d1d5db',
            color: '#fff'
          }}
        >
          {icon}
        </div>
        <div className='flex-1 min-w-0'>
          <div
            className='font-medium text-sm leading-tight truncate pr-6'
            style={{ color: isActive ? colors.text : '#9ca3af' }}
            title={label}
          >
            {label}
          </div>
          {description && (
            <div className='text-xs mt-0.5 line-clamp-2' style={{ color: isActive ? '#6b7280' : '#d1d5db' }}>
              {description}
            </div>
          )}
        </div>
      </div>

      {linksFiltered && linksFiltered.length > 0 && (
        <div className='mt-2 pt-2 border-t border-gray-200/50 flex flex-wrap gap-1'>
          {linksFiltered.slice(0, 2).map((link, idx) => (
            <a
              key={idx}
              href={link.value}
              target='_blank'
              rel='noopener noreferrer'
              onClick={(e) => e.stopPropagation()}
              className='inline-flex items-center gap-1 text-xs px-1.5 py-0.5 rounded transition-colors hover:opacity-80'
              style={{
                backgroundColor: isActive ? `${colors.border}20` : '#f3f4f6',
                color: isActive ? colors.text : '#9ca3af'
              }}
              title={link.value}
            >
              <ExternalLink className='h-2.5 w-2.5' />
              <span className='truncate max-w-[60px]'>{link.type}</span>
            </a>
          ))}
          {linksFiltered.length > 2 && (
            <span className='text-xs px-1' style={{ color: isActive ? colors.text : '#9ca3af' }}>
              +{linksFiltered.length - 2}
            </span>
          )}
        </div>
      )}

      {/* @ts-ignore */}
      <Handle
        type='source'
        position={Position.Bottom}
        className='!w-3 !h-3 !border-2'
        style={{
          backgroundColor: isActive ? colors.border : '#d1d5db',
          borderColor: '#fff'
        }}
      />

      {!isActive && <div className='absolute -top-1 -right-1 w-3 h-3 bg-gray-400 rounded-full border-2 border-white' />}

      <a
        href={RoutesLocation.targetsEdit(nodeId)}
        target='_blank'
        onClick={(e) => e.stopPropagation()}
        className='absolute -top-2 -right-2 w-6 h-6 bg-white rounded-full shadow-md border border-gray-200 flex items-center justify-center hover:bg-gray-50 transition-colors'
        title='Перейти к таргету'
        rel='noreferrer'
      >
        <ArrowUpRight className='h-3.5 w-3.5 text-gray-600' />
      </a>
    </div>
  );
});

TargetsNode.displayName = 'TargetsNode';
