import { RouterIcon } from './Router';
import { SwitchIcon } from './Switch';
import { DesktopIcon } from './Desktop';
import { ServerIcon } from './Server';
import { CloudIcon } from './Cloud';
import { ReactNode } from 'react';

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
  return <div className='size-6 bg-gray-200 dark:bg-gray-700 rounded-full' />;
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
