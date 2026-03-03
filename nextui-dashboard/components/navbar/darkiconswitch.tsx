import React from 'react';
import useThemeBrowser from '@/components/navbar/useTheme';
import { Moon, Sun } from 'lucide-react';

export const NavbarDarkModeToggle = () => {
  const { theme, setTheme } = useThemeBrowser();
  return (
    <div
      onClick={() => {
        setTheme(theme === 'dark' ? 'light' : 'dark');
      }}
      className='cursor-pointer border-1 border-solid border-black dark:border-white rounded-xl p-1'
    >
      {theme === 'dark' ? <Sun className='w-6 h-6 stroke-[#969696]' /> : <Moon className='w-6 h-6 stroke-[#969696]' />}
    </div>
  );
};
