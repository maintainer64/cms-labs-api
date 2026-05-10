import { RouterIcon } from './Router';
import { SwitchIcon } from './Switch';
import { DesktopIcon } from './Desktop';
import { ServerIcon } from './Server';
import { CloudIcon } from './Cloud';
import { ReactNode } from 'react';
import { UnknownIcon } from '@/components/topology/icons/Unknown';

const iconComponents: Record<string, ReactNode> = {
  router: <RouterIcon />,
  switch: <SwitchIcon />,
  desktop: <DesktopIcon />,
  linux: <DesktopIcon />,
  pc: <DesktopIcon />,
  server: <ServerIcon />,
  cloud: <CloudIcon />
};

const getIconRelevant = (icon?: string): ReactNode => {
  const iconKey = (icon || '').toLowerCase().replace(/\.[^/.]+$/, '');
  if (iconKey in iconComponents) {
    return iconComponents[iconKey];
  }
  return <UnknownIcon />;
};

interface Props {
  label?: string;
  icon?: string;
}

export const ImageIcon = ({ icon, label }: Props) => {
  return (
    <div
      className='
  w-full h-full
  [&>svg]:w-full
  [&>svg]:h-full
  [&>svg]:max-h-full
  [&>svg]:object-contain
'
    >
      {getIconRelevant(icon)}
    </div>
  );
};

export default ImageIcon;
