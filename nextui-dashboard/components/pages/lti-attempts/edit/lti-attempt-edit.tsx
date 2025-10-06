'use client';
import React from 'react';
import { HouseIcon } from '@/components/icons/breadcrumb/house-icon';
import { RoutesLocation } from '@/components/routes';
import useLanguageBrowser from '@/helpers/locale';
import { CrumbsLayout } from '@/components/layout/crumbs';
import { useParams } from 'react-router-dom';
import { LtiAttemptEditForm } from '@/components/pages/lti-attempts/edit/form';
import { UsersIcon } from '@/components/icons/breadcrumb/users-icon';
import { LtiAttemptIcon } from '@/components/icons/breadcrumb/lti-attempt';
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
      icon: <HouseIcon />,
      name: locale.Sidebar.Home,
      href: RoutesLocation.home()
    },
    {
      icon: <UsersIcon />,
      name: locale.Sidebar.Users,
      href: RoutesLocation.accountsEdit(attempt?.userId?.toString() || '0')
    },
    {
      icon: <LtiAttemptIcon />,
      name: locale.Sidebar.LTIAttempts,
      href: RoutesLocation.ltiAttemptUser(attempt?.userId?.toString() || '0')
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
