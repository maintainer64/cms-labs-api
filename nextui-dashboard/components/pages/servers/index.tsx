import { Button } from '@heroui/react';
import React, { useState } from 'react';
import { House, Server } from 'lucide-react';
import { RoutesLocation } from '@/components/routes';
import useLanguageBrowser from '@/helpers/locale';
import { CrumbsLayout } from '@/components/layout/crumbs';
import { Link } from 'react-router-dom';
import SearchInput from '@/components/sidebar/search-input';
import { ServerTableWrapper } from '@/components/pages/servers/table/table';
import { MapServerItem } from '@/helpers/queries/server/use-query-server-get';
import { useInfinityServerList } from '@/helpers/queries/server/use-infinity-server-list';

export const ServersList = () => {
  const { locale } = useLanguageBrowser();
  const {
    locale: {
      Tables: { ServersTable }
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
      icon: undefined,
      name: locale.Sidebar.AnyList,
      href: '#'
    }
  ];
  const [searchTerm, setSearchTerm] = useState<string>('');
  const response = useInfinityServerList({ limit: 100, search: searchTerm });
  const rows =
    response?.data?.pages.flatMap((p) => p?.model?.map((item) => MapServerItem(item.model, item.roles)) ?? []) || [];
  const totalCount = response.data?.pages[0]?.totalCount ?? 0;
  return (
    <CrumbsLayout name={`${ServersTable.Title} (${totalCount})`} crumbs={crumbs}>
      <>
        <div className='flex justify-between flex-wrap gap-4 items-center'>
          <div className='flex items-center gap-3 flex-nowrap w-full'>
            <SearchInput placeholder={PnetServersTable.SearchBar} setValue={setSearchTerm} />
            <Link to={RoutesLocation.serversCreate()}>
              <Button color='primary'>{PnetServersTable.ButtonAdd}</Button>
            </Link>
          </div>
        </div>
        <div className='max-w-[95rem] mx-auto w-full'>
          <ServerTableWrapper
            rows={rows}
            isLoading={response.isLoading}
            loadMore={response.fetchNextPage.bind(response.fetchNextPage)}
          />
        </div>
      </>
    </CrumbsLayout>
  );
};
