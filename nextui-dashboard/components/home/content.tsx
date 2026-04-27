'use client';
import React from 'react';
import { CardLastAttempt } from './card-last-attempt';
import HomeUsersWidget from '@/app/(app)/home/users-table';
import { CardServers } from '@/components/home/card-servers';
import { ContentCardWrapperMain } from '@/components/home/card-wrapper';
import useLanguageBrowser from '@/helpers/locale';
import { RoutesLocation } from '@/components/routes';
import CardServersDistribute from '@/components/home/card-servers-distribute';

export const Content = () => {
  const {
    locale: { Servers, ServersQueue }
  } = useLanguageBrowser();
  return (
    <div className='h-full lg:px-6'>
      <div className='flex justify-center gap-4 xl:gap-6 lg:px-0 flex-wrap xl:flex-nowrap max-w-[90rem] mx-auto w-full'>
        <div className='mt-6 gap-6 flex flex-col w-full'>
          {/* ServersCharts */}
          <ContentCardWrapperMain
            title={Servers.StatsChart.Title}
            link={RoutesLocation.servers()}
            wrapChildren={true}
          >
            <CardServers />
          </ContentCardWrapperMain>
          <ContentCardWrapperMain title={ServersQueue.RoundRobinChart.Title} wrapChildren={true}>
            <CardServersDistribute />
          </ContentCardWrapperMain>
        </div>

        {/* Left Section */}
        <div className='mt-4 gap-2 flex flex-col xl:max-w-md w-full'>
          <div>
            <CardLastAttempt />
          </div>
        </div>
      </div>
      <div className='gap-4 xl:gap-6 pt-3 px-4 lg:px-0 xl:flex-nowrap sm:pt-10 max-w-[90rem] mx-auto w-full'>
        <HomeUsersWidget />
      </div>
    </div>
  );
};
