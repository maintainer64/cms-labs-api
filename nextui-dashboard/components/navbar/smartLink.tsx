import React, { ReactNode, useRef } from 'react';
import { Link } from 'react-router-dom';

interface Props {
  to?: any;
  children: ReactNode;
  className?: string;
  title?: string;
}

export const SmartLink = ({ to, children, className, title }: Props) => {
  const tabRef = useRef<WindowProxy | null>(null);

  const handleClick = (e: React.MouseEvent) => {
    e.preventDefault();
    if (tabRef.current && !tabRef.current.closed) {
      tabRef.current.focus();
    } else {
      tabRef.current = window.open(to, '_blank');
    }
  };
  return (
    <Link to={to} className={className} target='_blank' onClick={handleClick} title={title}>
      {children}
    </Link>
  );
};
