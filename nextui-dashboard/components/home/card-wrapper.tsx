'use client';
import React from 'react';
import { Link as LinkComponent } from '@heroui/link';
import { Link } from 'react-router-dom';
import useLanguageBrowser from '@/helpers/locale';

interface Props {
  title?: React.ReactNode;
  link?: string;
  children: React.ReactNode;
  wrapChildren?: boolean;
}

export const ContentCardWrapperMain = ({ title, link, children, wrapChildren }: Props) => {
  const {
    locale: {
      Sidebar: { ViewAll }
    }
  } = useLanguageBrowser();
  const component = wrapChildren ? (
    <div className='w-full bg-default-50 shadow-lg rounded-2xl p-6'>{children}</div>
  ) : (
    children
  );
  return (
    <div className='flex flex-wrap justify-between h-full'>
      <h3 className='text-xl font-semibold'>{title}</h3>
      {link && (
        <LinkComponent href='#' color='primary' className='cursor-pointer'>
          <Link to={link}>{ViewAll}</Link>
        </LinkComponent>
      )}
      {component}
    </div>
  );
};
