'use client';
import React from 'react';
import { House, ScrollText, Server } from 'lucide-react';
import { RoutesLocation } from '@/components/routes';
import useLanguageBrowser from '@/helpers/locale';
import { CrumbsLayout } from '@/components/layout/crumbs';
import { useParams } from 'react-router-dom';
import { ServiceCardEditForm } from '@/components/pages/service-cards/edit/form';

export const ServiceCardsEdit = () => {
  const { id } = useParams();
  const { locale } = useLanguageBrowser();
  const {
    locale: {
      Tables: { ServiceCardsTable }
    }
  } = useLanguageBrowser();
  const crumbs = [
    {
      icon: <House className='w-5 h-5 stroke-[#969696]' />,
      name: locale.Sidebar.Home,
      href: RoutesLocation.home()
    },
    {
      icon: <Server className='w-5 h-5 stroke-[#969696]' />,
      name: locale.Sidebar.Servers,
      href: RoutesLocation.servers()
    },
    {
      icon: <ScrollText className='w-5 h-5 stroke-[#969696]' />,
      name: locale.Sidebar.ServiceCards,
      href: RoutesLocation.serviceCards()
    },
    {
      icon: undefined,
      name: locale.Sidebar.Edit,
      href: '#'
    }
  ];
  return (
    <CrumbsLayout name={ServiceCardsTable.Title} crumbs={crumbs}>
      <div className='max-w-[95rem] mx-auto w-full'>
        <ServiceCardEditForm id={parseInt(id ?? '', 10)} />
      </div>
    </CrumbsLayout>
  );
};
