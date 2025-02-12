import { Dropdown, DropdownItem, DropdownMenu, DropdownTrigger, NavbarItem } from '@nextui-org/react';
import React, { useCallback } from 'react';
import { DarkModeSwitch } from './darkmodeswitch';
import { userClearCookies } from '@/helpers/queries/jwt/userClearCookies';
import useLanguageBrowser from '@/helpers/locale';
import { useNavigate } from 'react-router-dom';
import { useUserProfile } from '@/components/providers/auth-jwt/hooks';
import CustomAvatar from '@/components/sidebar/avatar';
import { RoutesLocation } from '@/components/routes';

export const UserDropdown = () => {
  const navigate = useNavigate();
  const { locale } = useLanguageBrowser();
  const user = useUserProfile();

  const handleLogout = useCallback(async () => {
    await userClearCookies();
    navigate(RoutesLocation.login());
  }, [navigate]);

  const handleChangeLanguage = useCallback(() => {
    navigate(RoutesLocation.language());
  }, [navigate]);

  const handleToServicesCards = useCallback(() => {
    navigate(RoutesLocation.serviceCards());
  }, [navigate]);

  const handleToPasswordChange = useCallback(() => {
    navigate(RoutesLocation.profileChangePassword());
  }, [navigate]);

  return (
    <Dropdown>
      <NavbarItem>
        <DropdownTrigger>
          {CustomAvatar({
            tooltip: false,
            as: 'button',
            size: 'md',
            username: user.name,
            name: user.name,
            email: user.email
          })}
        </DropdownTrigger>
      </NavbarItem>
      <DropdownMenu aria-label='User menu actions'>
        <DropdownItem key='profile' className='flex flex-col justify-start w-full items-start'>
          <p>{locale.UserNavBar.SignedAs}</p>
          <p>{user.email}</p>
        </DropdownItem>
        <DropdownItem key='password' onPress={handleToPasswordChange}>
          {locale.UserNavBar.PasswordChange}
        </DropdownItem>
        <DropdownItem key='language' onPress={handleChangeLanguage}>
          {locale.UserNavBar.LanguageChange}
        </DropdownItem>
        <DropdownItem key='services' onPress={handleToServicesCards}>
          {locale.Sidebar.ServiceCards}
        </DropdownItem>
        <DropdownItem key='logout' color='danger' className='text-danger' onPress={handleLogout}>
          {locale.UserNavBar.Logout}
        </DropdownItem>
        <DropdownItem key='switch'>
          <DarkModeSwitch />
        </DropdownItem>
      </DropdownMenu>
    </Dropdown>
  );
};
