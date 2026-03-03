import { Image } from '@heroui/react';
import { Divider } from '@heroui/divider';
import useLanguageBrowser from '@/helpers/locale';
import { Link } from 'react-router-dom';
import { RoutesLocation } from '@/components/routes';

interface Props {
  children: React.ReactNode;
}

export const AuthLayoutWrapper = ({ children }: Props) => {
  const { locale } = useLanguageBrowser();
  return (
    <div className='flex h-screen'>
      <div className='flex-1 flex-col flex items-center justify-center p-6'>{children}</div>

      <div className='hidden my-10 md:block'>
        <Divider orientation='vertical' />
      </div>

      <div className='hidden md:flex flex-1 relative flex items-center justify-center p-6'>
        <div className='z-10'>
          <h1 className='font-bold text-[45px]'>{locale.Auth.MainTitle}</h1>
          <div className='font-light text-slate-400 mt-4'>{locale.Auth.MainDescription}</div>
          <div className='font-light underline text-slate-400 mt-10'>
            <Link to={RoutesLocation.language()}>{locale.Auth.MainChangeLanguage}</Link>
          </div>
        </div>
      </div>
    </div>
  );
};
