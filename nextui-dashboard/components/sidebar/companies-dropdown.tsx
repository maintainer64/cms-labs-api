'use client';
import { Dropdown, DropdownItem, DropdownMenu, DropdownSection, DropdownTrigger } from '@nextui-org/react';
import React from 'react';
import { AcmeIcon } from '../icons/acme-icon';
import { BottomIcon } from '../icons/sidebar/bottom-icon';
import useLanguageBrowser from '@/helpers/locale';

interface Company {
  title: string;
  description: string;
  link: string;
  logo?: React.ReactNode;
}

interface CompaniesDropdownProps {
  companies?: Company[];
}

export const CompaniesDropdown = ({ companies }: CompaniesDropdownProps) => {
  const { locale } = useLanguageBrowser();
  const company: Company = {
    title: locale.CompaniesDropdown.Title,
    description: locale.CompaniesDropdown.Description,
    link: '/',
    logo: <AcmeIcon />
  };
  return (
    <Dropdown
      classNames={{
        base: 'w-full min-w-[260px]'
      }}
    >
      <DropdownTrigger className='cursor-pointer'>
        <div className='flex items-center gap-2'>
          {company.logo}
          <div className='flex flex-col gap-4'>
            <h3 className='text-xl font-medium m-0 text-default-900 -mb-4 whitespace-nowrap'>{company.title}</h3>
            <span className='text-xs font-medium text-default-500'>{company.description}</span>
          </div>
          <BottomIcon />
        </div>
      </DropdownTrigger>
      <DropdownMenu aria-label='Avatar Actions'>
        <DropdownSection title={locale.CompaniesDropdown.ContentService}>
          {(companies || []).map((item, index) => (
            <DropdownItem
              key={index}
              href={item.link}
              startContent={item.logo}
              description={item.description}
              classNames={{
                base: 'py-4',
                title: 'text-base font-semibold'
              }}
            >
              {item.title}
            </DropdownItem>
          ))}
        </DropdownSection>
      </DropdownMenu>
    </Dropdown>
  );
};
