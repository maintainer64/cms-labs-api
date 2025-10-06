import { Button } from '@heroui/react';
import React, { useState } from 'react';
import { HouseIcon } from '@/components/icons/breadcrumb/house-icon';
import { RoutesLocation } from '@/components/routes';
import useLanguageBrowser from '@/helpers/locale';
import { CrumbsLayout } from '@/components/layout/crumbs';
import { Link } from 'react-router-dom';
import SearchInput from '@/components/sidebar/search-input';
import { RouterIcon } from '@/components/icons/breadcrumb/router-icon';
import { LTIRoutingTableWrapper } from '@/components/pages/lti-routings/table/table';
import { useInfinityLtiRoutingList } from '@/helpers/queries/lti_routing/use-infinity-lti-routing-list';

export const LTIRoutingList = () => {
  const { locale } = useLanguageBrowser();
  const {
    locale: {
      Tables: { LTIRoutingTable }
    }
  } = useLanguageBrowser();
  const crumbs = [
    {
      icon: <HouseIcon />,
      name: locale.Sidebar.Home,
      href: RoutesLocation.home()
    },
    {
      icon: <RouterIcon />,
      name: locale.Sidebar.LTIRouting,
      href: RoutesLocation.ltiRouting()
    },
    {
      icon: undefined,
      name: locale.Sidebar.AnyList,
      href: '#'
    }
  ];
  const [searchTerm, setSearchTerm] = useState<string>('');
  const response = useInfinityLtiRoutingList({ limit: 100, search: searchTerm });
  const rows = response?.data?.pages.flatMap((p) => p?.model ?? []) || [];
  const totalCount = response.data?.pages[0]?.totalCount ?? 0;
  return (
    <CrumbsLayout name={`${LTIRoutingTable.Title} (${totalCount})`} crumbs={crumbs}>
      <>
        <div className='flex justify-between flex-wrap gap-4 items-center'>
          <div className='flex items-center gap-3 flex-nowrap w-full'>
            <SearchInput placeholder={LTIRoutingTable.SearchBar} setValue={setSearchTerm} />
            <Link to={RoutesLocation.ltiRoutingCreate()}>
              <Button color='primary'>{LTIRoutingTable.ButtonAdd}</Button>
            </Link>
          </div>
        </div>
        <div className='max-w-[95rem] mx-auto w-full'>
          <LTIRoutingTableWrapper
            rows={rows}
            isLoading={response.isLoading}
            loadMore={response.fetchNextPage.bind(response.fetchNextPage)}
          />
        </div>
      </>
    </CrumbsLayout>
  );
};
