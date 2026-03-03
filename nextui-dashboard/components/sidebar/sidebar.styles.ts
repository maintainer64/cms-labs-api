import { tv } from '@heroui/react';

export const SidebarWrapper = tv({
  base: 'bg-background transition-transform h-full fixed -translate-x-full w-64 shrink-0 z-[202] overflow-y-auto border-r border-divider flex-col pb-6 px-3 md:ml-0 md:flex md:static md:h-screen md:translate-x-0 '
});
export const Overlay = tv({
  base: 'bg-[rgb(15_23_42/0.3)] fixed inset-0 z-[201] opacity-80 transition-opacity md:hidden md:z-auto md:opacity-100'
});

export const Header = tv({
  base: 'flex gap-5 px-6 items-center',
  variants: {
    collapsed: {
      false: 'px-4'
    }
  }
});

export const Body = tv({
  base: 'flex flex-col gap-6 mt-3 px-2'
});

export const Sidebar = Object.assign(SidebarWrapper, {
  Header,
  Body,
  Overlay
});
