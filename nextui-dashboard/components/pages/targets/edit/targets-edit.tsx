'use client';
import React from 'react';
import { House, Map } from 'lucide-react';
import { RoutesLocation } from '@/components/routes';
import useLanguageBrowser from '@/helpers/locale';
import { CrumbsLayout } from '@/components/layout/crumbs';
import { useParams } from 'react-router-dom';
import { TargetEditForm } from '@/components/pages/targets/edit/form';
import {useQueryTargetGet} from "@/helpers/queries/target/use-query-target-get";

export const TargetEdit = () => {
  const { id } = useParams();
  const { locale } = useLanguageBrowser();
  const {
    locale: { Target }
  } = useLanguageBrowser();
  const crumbs = [
    {
      icon: <House className='w-5 h-5 stroke-[#969696]' />,
      name: locale.Sidebar.Home,
      href: RoutesLocation.home()
    },
    {
      icon: <Map className='w-5 h-5 stroke-[#969696]' />,
      name: Target.Title,
      href: RoutesLocation.targets()
    },
    {
      icon: undefined,
      name: locale.Sidebar.Edit,
      href: '#'
    }
  ];
  const response = useQueryTargetGet({id: id || ''});
  return (
    <CrumbsLayout name={`${Target.ButtonEdit} ${response?.data?.name}`} crumbs={crumbs}>
      <div className='max-w-[95rem] mx-auto w-full'>
        <TargetEditForm id={id} />
      </div>
    </CrumbsLayout>
  );
};
