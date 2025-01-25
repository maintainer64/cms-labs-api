'use client';

import React from 'react';
import { Link } from 'react-router-dom';

export interface CrumbsItem {
  icon?: React.ReactNode;
  name: string;
  href: string;
}

interface CrumbsLayoutProps {
  crumbs: CrumbsItem[];
  name?: string;
  children: React.ReactNode;
}

export const CrumbsLayout = ({ children, name, crumbs }: CrumbsLayoutProps) => {
  const pageTitle = name ? <h3 className='text-xl font-semibold'>{name}</h3> : null;
  const crumbsBlock = crumbs.map((item, index, items) => (
    <li key={index} className='flex gap-2'>
      {item.icon}
      <Link to={item.href}>
        <span>{item.name}</span>
      </Link>
      {index !== items.length - 1 ? <span>&nbsp;/&nbsp;&nbsp;</span> : null}
    </li>
  ));
  return (
    <div className='my-10 px-4 lg:px-6 max-w-[95rem] mx-auto w-full flex flex-col gap-4'>
      <ul className='flex'>{crumbsBlock}</ul>
      {pageTitle}
      {children}
    </div>
  );
};
