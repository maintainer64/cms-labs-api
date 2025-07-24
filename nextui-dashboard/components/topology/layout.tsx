'use client';
import React from 'react';
import TopologyMenu from '@/components/topology/menu/menu';

interface Props {
  children: React.ReactNode;
}

export const TopologyLayout = ({ children }: Props) => {
  return (
    <div className='flex flex-col h-screen'>
      <TopologyMenu />
      <div className='flex-1 w-full'>{children}</div>
    </div>
  );
};
