'use client';
import React from 'react';
import { House, Split } from 'lucide-react';
import { RoutesLocation } from '@/components/routes';
import useLanguageBrowser from '@/helpers/locale';
import { CrumbsLayout } from '@/components/layout/crumbs';
import { useParams } from 'react-router-dom';
import { LtiRoutingEditForm } from '@/components/pages/lti-routings/edit/form';

export const LTIRoutingEdit = () => {
  const { id } = useParams();
  const { locale } = useLanguageBrowser();
  const {
    locale: {
      Tables: { LTIRoutingTable }
    }
  } = useLanguageBrowser();
  const crumbs = [
    {
      icon: <House className='w-5 h-5 stroke-[#969696]' />,
      name: locale.Sidebar.Home,
      href: RoutesLocation.home()
    },
    {
      icon: <Split className='w-5 h-5 stroke-[#969696]' />,
      name: locale.Sidebar.LTIRouting,
      href: RoutesLocation.ltiRouting()
    },
    {
      icon: undefined,
      name: locale.Sidebar.Edit,
      href: '#'
    }
  ];
  return (
    <CrumbsLayout name={LTIRoutingTable.Title} crumbs={crumbs}>
      <div className='max-w-[95rem] mx-auto w-full'>
        <LtiRoutingEditForm id={parseInt(id ?? '', 10)} />
      </div>
    </CrumbsLayout>
  );
};
