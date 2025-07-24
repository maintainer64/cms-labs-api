import React from 'react';
import { CircularProgress } from '@heroui/react';

interface LoadingProps {
  size?: 'md' | 'sm' | 'lg' | undefined;
}

export function Loading({ size }: LoadingProps) {
  return <CircularProgress aria-label='Loading...' color='primary' size={size} />;
}

export const HorizontalInfiniteLoader = () => (
  <div className='fixed top-0 left-0 right-0 h-1 z-50 w-full h-1 overflow-hidden bg-gray-200'>
    <div className='absolute top-0 left-0 w-1/3 h-full bg-blue-500 animate-slide'></div>
  </div>
);
