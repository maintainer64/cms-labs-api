'use client';
import React from 'react';
import { ReactFlowProvider } from '@xyflow/react';

interface Props {
  children: React.ReactNode;
}

export const TopologyLayout = ({ children }: Props) => {
  return (
    <ReactFlowProvider>
      <div className='flex flex-col h-screen'>
        <div className='flex-1 w-full'>{children}</div>
      </div>
    </ReactFlowProvider>
  );
};
