'use client';
import React from 'react';
import { HouseIcon } from '@/components/icons/breadcrumb/house-icon';
import { RoutesLocation } from '@/components/routes';
import useLanguageBrowser from '@/helpers/locale';
import { CrumbsLayout } from '@/components/layout/crumbs';
import { useParams } from 'react-router-dom';
import { ServersIcon } from '@/components/icons/breadcrumb/servers-icon';
import { ServiceCardEditForm } from '@/components/pages/service-cards/edit/form';
import { ServicesCardIcon } from '@/components/icons/breadcrumb/services-card-icon';

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
      icon: <HouseIcon />,
      name: locale.Sidebar.Home,
      href: RoutesLocation.home()
    },
    {
      icon: <ServersIcon />,
      name: locale.Sidebar.Servers,
      href: RoutesLocation.pnetServers()
    },
    {
      icon: <ServicesCardIcon />,
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
