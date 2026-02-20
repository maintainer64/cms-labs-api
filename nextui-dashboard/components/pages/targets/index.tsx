import { Button } from '@heroui/react';
import React, { useState } from 'react';
import { House, Map } from 'lucide-react';
import { RoutesLocation } from '@/components/routes';
import useLanguageBrowser from '@/helpers/locale';
import { CrumbsLayout } from '@/components/layout/crumbs';
import { Link } from 'react-router-dom';
import SearchInput from '@/components/sidebar/search-input';
import { TargetsMap } from '@/app/(app)/targets/TargetsMap';

export const TargetsList = () => {
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
    }
  ];
  const [searchTerm, setSearchTerm] = useState<string>('');
  return (
    <CrumbsLayout crumbs={crumbs}>
      <>
        <div className='flex justify-between flex-wrap gap-4 items-center'>
          <div className='flex items-center gap-3 flex-nowrap w-full'>
            <SearchInput placeholder={Target.SearchBar} setValue={setSearchTerm} />
            <Link to={RoutesLocation.targetsCreate()}>
              <Button color='primary'>{Target.ButtonAdd}</Button>
            </Link>
          </div>
        </div>
        <TargetsMap />
      </>
    </CrumbsLayout>
  );
};
