'use client';

import React from 'react';
import { SidebarWrapper } from '../sidebar/sidebar';
import { SidebarContext } from './layout-context';
import useCollapseBrowser from '@/components/navbar/useCollapse';

interface Props {
  children: React.ReactNode;
}

export const Layout = ({ children }: Props) => {
  const sidebarState = useCollapseBrowser();

  return (
    <SidebarContext.Provider
      value={{
        collapsed: sidebarState.collapsed,
        setCollapsed: sidebarState.setCollapsed
      }}
    >
      <section className='flex'>
        <SidebarWrapper />
        <div className='relative flex flex-col flex-1 overflow-y-auto overflow-x-hidden'>{children}</div>
      </section>
    </SidebarContext.Provider>
  );
};
