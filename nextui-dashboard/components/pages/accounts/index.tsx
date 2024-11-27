import { Button } from '@nextui-org/react';
import React, { useState } from 'react';
import { HouseIcon } from '@/components/icons/breadcrumb/house-icon';
import { UsersIcon } from '@/components/icons/breadcrumb/users-icon';
import { UsersTableWrapper } from '@/components/pages/accounts/table/table';
import { RoutesLocation } from '@/components/routes';
import useLanguageBrowser from '@/helpers/locale';
import { CrumbsLayout } from '@/components/layout/crumbs';
import SearchInput from '@/components/sidebar/search-input';
import { Link } from 'react-router-dom';
import { useUsersList } from '@/helpers/queries/users/get';

export const Accounts = () => {
  const { locale } = useLanguageBrowser();
  const {
    locale: {
      Tables: { UsersTable }
    }
  } = useLanguageBrowser();
  const crumbs = [
    {
      icon: <HouseIcon />,
      name: locale.Sidebar.Home,
      href: RoutesLocation.home()
    },
    {
      icon: <UsersIcon />,
      name: locale.Sidebar.Users,
      href: RoutesLocation.accounts()
    },
    {
      icon: undefined,
      name: locale.Sidebar.AnyList,
      href: '#'
    }
  ];
  const [searchTerm, setSearchTerm] = useState<string>('');
  const response = useUsersList({ limit: 100, search: searchTerm });
  const users = response?.data?.pages.flatMap((p) => p.result?.model ?? []) || [];
  const totalCount = response.data?.pages[0].result?.total_count ?? 0;
  return (
    <CrumbsLayout name={`${UsersTable.Title} (${totalCount})`} crumbs={crumbs}>
      <>
        <div className='flex justify-between flex-wrap gap-4 items-center'>
          <div className='flex items-center gap-3 flex-nowrap w-full'>
            <SearchInput placeholder={UsersTable.SearchBar} setValue={setSearchTerm} />
            <Link to={RoutesLocation.accountsCreate()}>
              <Button color='primary'>{UsersTable.ButtonAdd}</Button>
            </Link>
          </div>
        </div>
        <div className='max-w-[95rem] mx-auto w-full'>
          <UsersTableWrapper
            users={users}
            isLoading={response.isLoading}
            loadMore={response.fetchNextPage.bind(response.fetchNextPage)}
          />
        </div>
      </>
    </CrumbsLayout>
  );
};
