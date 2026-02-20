import React from 'react';
import { BookPlus, House, UsersRound } from 'lucide-react';
import { RoutesLocation } from '@/components/routes';
import useLanguageBrowser from '@/helpers/locale';
import { CrumbsLayout } from '@/components/layout/crumbs';
import { LTIAttemptTableWrapper } from '@/components/pages/lti-attempts/table/table';
import { useParams } from 'react-router-dom';
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
      icon: <House className='w-5 h-5 stroke-[#969696]' />,
      name: locale.Sidebar.Home,
      href: RoutesLocation.home()
    },
    {
      icon: <UsersRound className='w-5 h-5 stroke-[#969696]' />,
      name: locale.Sidebar.Users,
      href: RoutesLocation.accountsEdit(id)
    },
    {
      icon: <BookPlus className='w-5 h-5 stroke-[#969696]' />,
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
