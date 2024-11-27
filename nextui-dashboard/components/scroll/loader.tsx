import React from 'react';

interface LoadingProps {
  size?: number;
}

export function Loading(props: LoadingProps) {
  const size = props.size ?? 24;
  return <div className={`animate-spin rounded-full h-${size} w-${size} border-t-4 border-black border-opacity-50`} />;
}
