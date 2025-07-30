import { Button } from '@heroui/react';
import { ReactNode } from 'react';
import { SmartLink } from '@/components/navbar/smartLink';

interface Props {
  onPress?: () => void;
  href?: string;
  className?: string;
  children?: ReactNode;
}

export const InfoModalButton = ({ onPress, href, className, children }: Props) => {
  if (href)
    return (
      <SmartLink className={className} to={href}>
        {children}
      </SmartLink>
    );
  return (
    <Button onPress={onPress} className={className}>
      {children}
    </Button>
  );
};
