import { Button } from '@heroui/react';
import React from 'react';
import { HouseIcon } from '@/components/icons/breadcrumb/house-icon';
import { RoutesLocation } from '@/components/routes';
import useLanguageBrowser from '@/helpers/locale';
import { CrumbsLayout } from '@/components/layout/crumbs';
import { Link } from 'react-router-dom';
import { ServersIcon } from '@/components/icons/breadcrumb/servers-icon';
import { ServiceCardsTableWrapper } from '@/components/pages/service-cards/table/table';
import { ServicesCardIcon } from '@/components/icons/breadcrumb/services-card-icon';
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
