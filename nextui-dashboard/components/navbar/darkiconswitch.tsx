import React from 'react';
import useThemeBrowser from '@/components/navbar/useTheme';
import { DarkModeIcon } from '@/components/icons/navbar/dark-mode-icon';
import { LightModeIcon } from '@/components/icons/navbar/light-mode-icon';

export const NavbarDarkModeToggle = () => {
  const { theme, setTheme } = useThemeBrowser();
  return (
    <div
      onClick={() => {
        setTheme(theme === 'dark' ? 'light' : 'dark');
      }}
      className='cursor-pointer border-1 border-solid border-black dark:border-white rounded-xl p-1'
    >
      {theme === 'dark' ? <LightModeIcon /> : <DarkModeIcon />}
    </div>
  );
};
