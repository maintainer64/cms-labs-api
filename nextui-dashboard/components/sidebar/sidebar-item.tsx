import { Link } from 'react-router-dom';
import React from 'react';
import { useSidebarContext } from '../layout/layout-context';
import clsx from 'clsx';

interface Props {
  title: string;
  icon: React.ReactNode;
  isActive?: boolean;
  href?: string;
  onPress?: () => void;
}

export const SidebarItem = ({ icon, title, isActive, href = '', onPress }: Props) => {
  const { collapsed } = useSidebarContext();

  return (
    <Link to={href} onClick={onPress} className='text-default-900 active:bg-none max-w-full'>
      <div
        className={clsx(
          isActive ? 'bg-primary-100 [&_svg_*]:stroke-primary-500' : 'hover:bg-default-100',
          'flex gap-2 w-full min-h-[44px] h-full items-center rounded-xl cursor-pointer transition-all duration-150 active:scale-[0.98]',
          collapsed ? 'px-5' : 'px-3.5'
        )}
      >
        {icon}
        {collapsed ? null : <span className='text-default-900'>{title}</span>}
      </div>
    </Link>
  );
};
