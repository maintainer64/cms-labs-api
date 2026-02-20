'use client';

import { createContext, useContext } from 'react';

interface SidebarContext {
  collapsed: boolean;
  setCollapsed: (collapsed: boolean) => void;
}

export const SidebarContext = createContext<SidebarContext>({
  collapsed: false,
  setCollapsed: (collapsed: boolean) => {}
});

export const useSidebarContext = () => {
  return useContext(SidebarContext);
};
