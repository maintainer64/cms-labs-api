'use client';
import React from 'react';
import { BookPlus, House, UsersRound } from 'lucide-react';
import { RoutesLocation } from '@/components/routes';
import useLanguageBrowser from '@/helpers/locale';
import { CrumbsLayout } from '@/components/layout/crumbs';
import { useParams } from 'react-router-dom';
import { LtiAttemptEditForm } from '@/components/pages/lti-attempts/edit/form';
import { useQueryLtiAttemptGet } from '@/helpers/queries/lti_attempt/use-query-lti-attempt-get';

export const LtiAttemptEdit = () => {
  const { id } = useParams();
  const { locale } = useLanguageBrowser();
  const {
    locale: {
      Tables: { LTIAttemptsTable }
    }
  } = useLanguageBrowser();
  const response = useQueryLtiAttemptGet({ id: parseInt(id ?? '') });
  const attempt = response.data?.model;
  const crumbs = [
    {
      icon: <House className='w-5 h-5 stroke-[#969696]' />,
      name: locale.Sidebar.Home,
      href: RoutesLocation.home()
    },
    {
      icon: <UsersRound className='w-5 h-5 stroke-[#969696]' />,
      name: locale.Sidebar.Users,
      href: RoutesLocation.accountsEdit(attempt?.userId?.toString() || '0')
    },
    {
      icon: <BookPlus className='w-5 h-5 stroke-[#969696]' />,
      name: locale.Sidebar.LTIAttempts,
      href: RoutesLocation.ltiAttempts({ userId: attempt?.userId })
    },
    {
      icon: undefined,
      name: locale.Sidebar.Edit,
      href: '#'
    }
  ];

  return (
    <CrumbsLayout name={LTIAttemptsTable.Title} crumbs={crumbs}>
      <div className='max-w-[95rem] mx-auto w-full'>
        <LtiAttemptEditForm id={parseInt(id ?? '', 10)} />
      </div>
    </CrumbsLayout>
  );
};
