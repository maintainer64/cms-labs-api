import React from 'react';
import { Sidebar } from './sidebar.styles';
import { ServicesDropdown } from './services-dropdown';
import { HomeIcon } from '../icons/sidebar/home-icon';
import { BalanceIcon } from '../icons/sidebar/balance-icon';
import { AccountsIcon } from '../icons/sidebar/accounts-icon';
import { SidebarItem } from './sidebar-item';
import { SidebarMenu } from './sidebar-menu';
import { useSidebarContext } from '../layout/layout-context';
import { useLocation } from 'react-router-dom';
import useLanguageBrowser from '@/helpers/locale';
import { RoutesLocation } from '@/components/routes';
import { LtiIcon } from '@/components/icons/breadcrumb/lti-icon';
import { ServersIcon } from '@/components/icons/breadcrumb/servers-icon';
import { RouterIcon } from '@/components/icons/breadcrumb/router-icon';

export const SidebarWrapper = () => {
  const { pathname } = useLocation();
  const { locale } = useLanguageBrowser();
  const { collapsed, setCollapsed } = useSidebarContext();

  return (
    <aside className='h-screen z-[20] sticky top-0'>
      {collapsed ? <div className={Sidebar.Overlay()} onClick={setCollapsed} /> : null}
      <div
        className={Sidebar({
          collapsed: collapsed
        })}
      >
        <div className={Sidebar.Header()}>
          <ServicesDropdown />
        </div>
        <div className='flex flex-col justify-between h-full'>
          <div className={Sidebar.Body()}>
            <SidebarItem title={locale.Sidebar.Home} icon={<HomeIcon />} isActive={pathname === '/'} href='/' />
            <SidebarMenu title={locale.Sidebar.MainMenu}>
              <SidebarItem
                isActive={pathname === RoutesLocation.ltiRouting()}
                title={locale.Sidebar.LTIRouting}
                icon={<RouterIcon />}
                href={RoutesLocation.ltiRouting()}
              />
              <SidebarItem
                isActive={pathname === RoutesLocation.ltiForms()}
                title={locale.Sidebar.LTIIntegrations}
                icon={<LtiIcon />}
                href={RoutesLocation.ltiForms()}
              />
              <SidebarItem
                isActive={pathname === RoutesLocation.accounts()}
                title={locale.Sidebar.Users}
                icon={<AccountsIcon />}
                href={RoutesLocation.accounts()}
              />
              <SidebarItem
                isActive={pathname === RoutesLocation.pnetServers()}
                title={locale.Sidebar.Servers}
                icon={<ServersIcon />}
                href={RoutesLocation.pnetServers()}
              />
            </SidebarMenu>
          </div>
        </div>
      </div>
    </aside>
  );
};
