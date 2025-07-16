import React from 'react';
import { TopologyFlowVisualization } from '@/components/topology/view';
import { TopologyConnect } from '@/components/topology/connect';

export const TopologyPageConnect = () => {
  return (
    <div className='flex h-screen'>
      <div className='flex-1 flex-col flex items-center justify-center p-6'>
        <TopologyConnect />
      </div>
    </div>
  );
};

export const TopologyPageView = () => {
  return (
    <div className='flex h-screen'>
      <div className='flex-1 flex-col flex items-center justify-center p-6'>
        <TopologyFlowVisualization />
      </div>
    </div>
  );
};
