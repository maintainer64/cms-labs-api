import { Button } from '@heroui/react';
import React, { useState } from 'react';
import { HouseIcon } from '@/components/icons/breadcrumb/house-icon';
import { RoutesLocation } from '@/components/routes';
import useLanguageBrowser from '@/helpers/locale';
import { CrumbsLayout } from '@/components/layout/crumbs';
import { Link } from 'react-router-dom';
import SearchInput from '@/components/sidebar/search-input';
import { UsersIcon } from '@/components/icons/breadcrumb/users-icon';
import { RolesIcon } from '@/components/icons/breadcrumb/roles-icon';
import { useRolesList } from '@/helpers/queries/roles/get';
import { RolesTableWrapper } from '@/components/pages/roles/table/table';

export const RolesList = () => {
  const { locale } = useLanguageBrowser();
  const {
    locale: {
      Tables: { RoleTable }
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
      icon: <RolesIcon />,
      name: locale.Sidebar.Roles,
      href: RoutesLocation.roles()
    },
    {
      icon: undefined,
      name: locale.Sidebar.AnyList,
      href: '#'
    }
  ];
  const [searchTerm, setSearchTerm] = useState<string>('');
  const queryRoles = useRolesList();
  const rows = queryRoles.data?.result?.model || [];
  const totalCount = queryRoles.data?.result?.total_count || 0;
  return (
    <CrumbsLayout name={`${RoleTable.Title} (${totalCount})`} crumbs={crumbs}>
      <>
        <div className='flex justify-between flex-wrap gap-4 items-center'>
          <div className='flex items-center gap-3 flex-nowrap w-full'>
            <SearchInput placeholder={RoleTable.SearchBar} setValue={setSearchTerm} />
            <Link to={RoutesLocation.rolesCreate()}>
              <Button color='primary'>{RoleTable.ButtonAdd}</Button>
            </Link>
          </div>
        </div>
        <div className='max-w-[95rem] mx-auto w-full'>
          <RolesTableWrapper
            rows={rows.filter((row) =>
              searchTerm.length < 3 ? true : row.name?.includes(searchTerm) || row.code?.includes(searchTerm)
            )}
            isLoading={queryRoles.isLoading}
          />
        </div>
      </>
    </CrumbsLayout>
  );
};
