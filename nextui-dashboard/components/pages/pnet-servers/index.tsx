import { Button } from '@heroui/react';
import React, { useState } from 'react';
import { HouseIcon } from '@/components/icons/breadcrumb/house-icon';
import { RoutesLocation } from '@/components/routes';
import useLanguageBrowser from '@/helpers/locale';
import { CrumbsLayout } from '@/components/layout/crumbs';
import { Link } from 'react-router-dom';
import SearchInput from '@/components/sidebar/search-input';
import { ServersIcon } from '@/components/icons/breadcrumb/servers-icon';
import { usePnetServerList } from '@/helpers/queries/pnet-server/get';
import { PnetServerTableWrapper } from '@/components/pages/pnet-servers/table/table';
import { MapServerItem } from '@/helpers/queries/pnet-server/model';

export const PnetServersList = () => {
  const { locale } = useLanguageBrowser();
  const {
    locale: {
      Tables: { PnetServersTable }
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
      icon: undefined,
      name: locale.Sidebar.AnyList,
      href: '#'
    }
  ];
  const [searchTerm, setSearchTerm] = useState<string>('');
  const response = usePnetServerList({ limit: 100, search: searchTerm });
  const rows =
    response?.data?.pages.flatMap((p) => p.result?.model?.map((item) => MapServerItem(item.model, item.roles)) ?? []) ||
    [];
  const totalCount = response.data?.pages[0].result?.total_count ?? 0;
  return (
    <CrumbsLayout name={`${PnetServersTable.Title} (${totalCount})`} crumbs={crumbs}>
      <>
        <div className='flex justify-between flex-wrap gap-4 items-center'>
          <div className='flex items-center gap-3 flex-nowrap w-full'>
            <SearchInput placeholder={PnetServersTable.SearchBar} setValue={setSearchTerm} />
            <Link to={RoutesLocation.pnetServersCreate()}>
              <Button color='primary'>{PnetServersTable.ButtonAdd}</Button>
            </Link>
          </div>
        </div>
        <div className='max-w-[95rem] mx-auto w-full'>
          <PnetServerTableWrapper
            rows={rows}
            isLoading={response.isLoading}
            loadMore={response.fetchNextPage.bind(response.fetchNextPage)}
          />
        </div>
      </>
    </CrumbsLayout>
  );
};
