import { Button } from '@heroui/react';
import React, { useState } from 'react';
import { RoutesLocation } from '@/components/routes';
import useLanguageBrowser from '@/helpers/locale';
import { CrumbsLayout } from '@/components/layout/crumbs';
import { Link } from 'react-router-dom';
import SearchInput from '@/components/sidebar/search-input';
import { CurlRequestTableWrapper } from '@/components/pages/curl-requests/table/table';
import { useInfinityCurlRequestList } from '@/helpers/queries/curl_request/use-infinity-curl-request-list';
import { House, Unplug } from 'lucide-react';

export const CurlRequestList = () => {
  const { locale } = useLanguageBrowser();
  const {
    locale: {
      Tables: { CurlRequestTable }
    }
  } = useLanguageBrowser();
  const crumbs = [
    {
      icon: <House className='w-5 h-5 stroke-[#969696]' />,
      name: locale.Sidebar.Home,
      href: RoutesLocation.home()
    },
    {
      icon: <Unplug className='w-5 h-5 stroke-[#969696]' />,
      name: locale.Sidebar.APIRequests,
      href: RoutesLocation.curlRequest()
    },
    {
      icon: undefined,
      name: locale.Sidebar.AnyList,
      href: '#'
    }
  ];
  const [searchTerm, setSearchTerm] = useState<string>('');
  const response = useInfinityCurlRequestList({ limit: 100, search: searchTerm });
  const rows = response?.data?.pages.flatMap((p) => p?.model ?? []) || [];
  const totalCount = response.data?.pages[0]?.totalCount ?? 0;
  return (
    <CrumbsLayout name={`${CurlRequestTable.Title} (${totalCount})`} crumbs={crumbs}>
      <>
        <div className='flex justify-between flex-wrap gap-4 items-center'>
          <div className='flex items-center gap-3 flex-nowrap w-full'>
            <SearchInput placeholder={CurlRequestTable.SearchBar} setValue={setSearchTerm} />
            <Link to={RoutesLocation.curlRequestCreate()}>
              <Button color='primary'>{CurlRequestTable.ButtonAdd}</Button>
            </Link>
          </div>
        </div>
        <div className='max-w-[95rem] mx-auto w-full'>
          <CurlRequestTableWrapper
            rows={rows}
            isLoading={response.isLoading}
            loadMore={response.fetchNextPage.bind(response.fetchNextPage)}
          />
        </div>
      </>
    </CrumbsLayout>
  );
};
