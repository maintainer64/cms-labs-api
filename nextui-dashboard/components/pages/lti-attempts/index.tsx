import React from 'react';
import { HouseIcon } from '@/components/icons/breadcrumb/house-icon';
import { RoutesLocation } from '@/components/routes';
import useLanguageBrowser from '@/helpers/locale';
import { CrumbsLayout } from '@/components/layout/crumbs';
import { UsersIcon } from '@/components/icons/breadcrumb/users-icon';
import { LTIAttemptTableWrapper } from '@/components/pages/lti-attempts/table/table';
import { useParams } from 'react-router-dom';
import { LtiAttemptIcon } from '@/components/icons/breadcrumb/lti-attempt';
import { useQueryLtiAttemptList } from '@/helpers/queries/lti_attempt/use-query-lti-attempt-list';

export const LTIAttemptListByUser = () => {
  const { id } = useParams();
  const { locale } = useLanguageBrowser();
  const {
    locale: {
      Tables: { LTIAttemptsTable }
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
      href: RoutesLocation.accountsEdit(id)
    },
    {
      icon: <LtiAttemptIcon />,
      name: locale.Sidebar.LTIAttempts,
      href: RoutesLocation.ltiAttemptUser(id)
    },
    {
      icon: undefined,
      name: locale.Sidebar.AnyList,
      href: '#'
    }
  ];
  const response = useQueryLtiAttemptList({ limit: 100, userIds: [Number(id)] });
  const items = response.data?.model || [];
  return (
    <CrumbsLayout name={`${LTIAttemptsTable.Title}`} crumbs={crumbs}>
      <>
        <div className='max-w-[95rem] mx-auto w-full'>
          <LTIAttemptTableWrapper isLoading={response.isLoading} rows={items} />
        </div>
      </>
    </CrumbsLayout>
  );
};
