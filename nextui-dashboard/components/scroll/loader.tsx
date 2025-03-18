import React from 'react';
import { CircularProgress } from '@heroui/react';

interface LoadingProps {
  size?: 'md' | 'sm' | 'lg' | undefined;
}

export function Loading({ size }: LoadingProps) {
  return <CircularProgress aria-label='Loading...' color='primary' size={size} />;
}
