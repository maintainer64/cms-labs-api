'use client';
import { Dropdown, DropdownItem, DropdownMenu, DropdownSection, DropdownTrigger, Image } from '@heroui/react';
import React from 'react';
import { CmsIcon } from '../icons/cms-icon';
import useLanguageBrowser from '@/helpers/locale';
import { useQueryServiceCardList } from '@/helpers/queries/service_card/use-query-service-card-list';
import { ChevronDown } from 'lucide-react';
import { useSidebarContext } from '@/components/layout/layout-context';

const DropdownServiceCards = () => {
  const response = useQueryServiceCardList({});
  const rows = response?.data?.model || [];
  return rows
    .filter((service) => service.isActive)
    .map((service, index) => {
      return (
        <DropdownItem
          key={index}
          onPress={() => {
            window.open(service.url, '_blank');
          }}
          href={service.url || '#'}
          target='_blank'
          startContent={service.imageUrl && <Image src={service.imageUrl} width={30} alt={service.imageUrl} />}
          description={service.description}
          classNames={{
            base: 'py-4',
            title: 'text-base font-semibold'
          }}
        >
          {service.name}
        </DropdownItem>
      );
    });
};

export const ServicesDropdown = () => {
  const { locale } = useLanguageBrowser();
  const { collapsed } = useSidebarContext();
  const company = {
    title: locale.CompaniesDropdown.Title,
    description: locale.CompaniesDropdown.Description,
    link: '/',
    logo: <CmsIcon />
  };
  return (
    <Dropdown
      classNames={{
        base: 'w-full min-w-[240px]'
      }}
    >
      <DropdownTrigger className='cursor-pointer'>
        <div className='flex items-center gap-2'>
          {company.logo}
          {collapsed ? null : (
            <>
              <div className='flex flex-col gap-4'>
                <h3 className='text-xl font-medium m-0 text-default-900 -mb-4 whitespace-nowrap'>{company.title}</h3>
                <span className='text-xs font-medium text-default-500'>{company.description}</span>
              </div>
              <ChevronDown className='w-3 h-3 stroke-[#969696]' />
            </>
          )}
        </div>
      </DropdownTrigger>
      <DropdownMenu aria-label='Avatar Actions'>
        <DropdownSection title={locale.Tables.ServiceCardsTable.Title}>{DropdownServiceCards()}</DropdownSection>
      </DropdownMenu>
    </Dropdown>
  );
};
