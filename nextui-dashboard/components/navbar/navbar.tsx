import { Input, Link, Navbar, NavbarContent } from '@heroui/react';
import React from 'react';
import { GithubIcon } from '../icons/navbar/github-icon';
import { SearchIcon } from '../icons/searchicon';
import { BurguerButton } from './burguer-button';
import { UserDropdown } from './user-dropdown';
import { NavbarDarkModeToggle } from '@/components/navbar/darkiconswitch';

interface Props {
  children: React.ReactNode;
}

export const NavbarWrapper = ({ children }: Props) => {
  return (
    <div className='relative flex flex-col flex-1 overflow-y-auto overflow-x-hidden'>
      <Navbar
        isBordered
        className='w-full'
        classNames={{
          wrapper: 'w-full max-w-full'
        }}
      >
        <NavbarContent className='md:hidden'>
          <BurguerButton />
        </NavbarContent>
        <NavbarContent className='w-full max-md:hidden'>
          <Input
            startContent={<SearchIcon />}
            isClearable
            className='w-full'
            classNames={{
              input: 'w-full',
              mainWrapper: 'w-full'
            }}
            placeholder='Search...'
          />
        </NavbarContent>
        <NavbarContent justify='end' className='w-fit data-[justify=end]:flex-grow-0'>
          <Link href='https://gitlab.com/a10869/api-modules' target={'_blank'}>
            <GithubIcon />
          </Link>

          <div className='max-md:hidden'>
            <NavbarDarkModeToggle />
          </div>
          <NavbarContent>
            <UserDropdown />
          </NavbarContent>
        </NavbarContent>
      </Navbar>
      {children}
    </div>
  );
};
