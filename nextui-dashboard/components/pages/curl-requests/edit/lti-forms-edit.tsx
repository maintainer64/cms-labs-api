'use client';
import React from 'react';
import { HouseIcon } from '@/components/icons/breadcrumb/house-icon';
import { RoutesLocation } from '@/components/routes';
import useLanguageBrowser from '@/helpers/locale';
import { CrumbsLayout } from '@/components/layout/crumbs';
import { useParams } from 'react-router-dom';
import { CurlRequestIcon } from '@/components/icons/breadcrumb/curl-request-icon';
import { CurlRequestEditForm } from '@/components/pages/curl-requests/edit/form';

export const CurlRequestEdit = () => {
  const { id } = useParams();
  const { locale } = useLanguageBrowser();
  const {
    locale: {
      Tables: { CurlRequestTable }
    }
  } = useLanguageBrowser();
  const crumbs = [
    {
      icon: <HouseIcon />,
      name: locale.Sidebar.Home,
      href: RoutesLocation.home()
    },
    {
      icon: <CurlRequestIcon />,
      name: locale.Sidebar.APIRequests,
      href: RoutesLocation.curlRequest()
    },
    {
      icon: undefined,
      name: locale.Sidebar.AnyList,
      href: '#'
    }
  ];
  return (
    <CrumbsLayout name={CurlRequestTable.Title} crumbs={crumbs}>
      <div className='max-w-[95rem] mx-auto w-full'>
        <CurlRequestEditForm id={parseInt(id ?? '', 10)} />
      </div>
    </CrumbsLayout>
  );
};
