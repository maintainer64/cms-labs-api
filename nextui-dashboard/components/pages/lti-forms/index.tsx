import { Button } from '@heroui/react';
import React, { useState } from 'react';
import { HouseIcon } from '@/components/icons/breadcrumb/house-icon';
import { RoutesLocation } from '@/components/routes';
import useLanguageBrowser from '@/helpers/locale';
import { CrumbsLayout } from '@/components/layout/crumbs';
import { Link } from 'react-router-dom';
import { LtiIcon } from '@/components/icons/breadcrumb/lti-icon';
import SearchInput from '@/components/sidebar/search-input';
import { LTIFormsTableWrapper } from '@/components/pages/lti-forms/table/table';
import { RoleBasedAccess } from '@/components/layout/roleBasedAccess';
import { UserRoleBase } from '@/helpers/queries/sso/auth';
import { useInfinityLtiFormList } from '@/helpers/queries/lti_form/use-infinity-lti-form-list';

export const LTIFormsList = () => {
  const { locale } = useLanguageBrowser();
  const {
    locale: {
      Tables: { LTIFormsTable }
    }
  } = useLanguageBrowser();
  const crumbs = [
    {
      icon: <HouseIcon />,
      name: locale.Sidebar.Home,
      href: RoutesLocation.home()
    },
    {
      icon: <LtiIcon />,
      name: locale.Sidebar.LTIIntegrations,
      href: RoutesLocation.ltiForms()
    },
    {
      icon: undefined,
      name: locale.Sidebar.AnyList,
      href: '#'
    }
  ];
  const [searchTerm, setSearchTerm] = useState<string>('');
  const response = useInfinityLtiFormList({ limit: 100, search: searchTerm });
  const rows = response?.data?.pages.flatMap((p) => p?.model ?? []) || [];
  const totalCount = response.data?.pages[0]?.totalCount ?? 0;
  return (
    <CrumbsLayout name={`${LTIFormsTable.Title} (${totalCount})`} crumbs={crumbs}>
      <>
        <div className='flex justify-between flex-wrap gap-4 items-center'>
          <div className='flex items-center gap-3 flex-nowrap w-full'>
            <SearchInput placeholder={LTIFormsTable.SearchBar} setValue={setSearchTerm} />
            <Link to={RoutesLocation.ltiFormsCreate()}>
              <Button color='primary'>{LTIFormsTable.ButtonAdd}</Button>
            </Link>
          </div>
        </div>
        <div className='max-w-[95rem] mx-auto w-full'>
          <RoleBasedAccess allowedRoles={[UserRoleBase.Admin]}>
            <LTIFormsTableWrapper
              rows={rows}
              isLoading={response.isLoading}
              loadMore={response.fetchNextPage.bind(response.fetchNextPage)}
            />
          </RoleBasedAccess>
        </div>
      </>
    </CrumbsLayout>
  );
};
