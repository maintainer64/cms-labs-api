import React from 'react';
import { useSidebarContext } from '@/components/layout/layout-context';

interface Props {
  title: string;
  children?: React.ReactNode;
}

export const SidebarMenu = ({ title, children }: Props) => {
  const { collapsed } = useSidebarContext();
  return (
    <div className='flex gap-2 flex-col'>
      {collapsed ? null : <span className='text-xs font-normal'>{title}</span>}
      {children}
    </div>
  );
};
