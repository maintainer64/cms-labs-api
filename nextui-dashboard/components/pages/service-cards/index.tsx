import { Button } from '@heroui/react';
import React from 'react';
import { House, ScrollText, Server } from 'lucide-react';
import { RoutesLocation } from '@/components/routes';
import useLanguageBrowser from '@/helpers/locale';
import { CrumbsLayout } from '@/components/layout/crumbs';
import { Link } from 'react-router-dom';
import { ServiceCardsTableWrapper } from '@/components/pages/service-cards/table/table';
import { useQueryServiceCardList } from '@/helpers/queries/service_card/use-query-service-card-list';

export const ServiceCardsList = () => {
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
      href: RoutesLocation.pnetServers()
    },
    {
      icon: <ScrollText className='w-5 h-5 stroke-[#969696]' />,
      name: locale.Sidebar.ServiceCards,
      href: RoutesLocation.serviceCards()
    },
    {
      icon: undefined,
      name: locale.Sidebar.AnyList,
      href: '#'
    }
  ];
  const response = useQueryServiceCardList({});
  const rows = response?.data?.model || [];
  const totalCount = response.data?.totalCount ?? 0;
  return (
    <CrumbsLayout name={`${ServiceCardsTable.Title} (${totalCount})`} crumbs={crumbs}>
      <>
        <div className='flex justify-between flex-wrap gap-4 items-center'>
          <div className='flex items-center gap-3 flex-nowrap w-full'>
            <Link to={RoutesLocation.serviceCardsCreate()}>
              <Button color='primary'>{ServiceCardsTable.ButtonAdd}</Button>
            </Link>
          </div>
        </div>
        <div className='max-w-[95rem] mx-auto w-full'>
          <ServiceCardsTableWrapper rows={rows} isLoading={response.isLoading} loadMore={() => {}} />
        </div>
      </>
    </CrumbsLayout>
  );
};
