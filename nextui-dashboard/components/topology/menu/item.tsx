import { Link, NavbarMenuItem } from '@heroui/react';

interface TopologyMenuItemProps {
  title: string;
  color?: 'foreground' | 'danger' | 'warning';
  target?: string;
  onClick?: () => void;
  href?: string;
}

export function TopologyMenuItem({ title, color, href, onClick, target }: TopologyMenuItemProps) {
  return (
    <NavbarMenuItem>
      <Link className='w-full' color={color || 'foreground'} href={href} size='lg' target={target} onPress={onClick}>
        {title}
      </Link>
    </NavbarMenuItem>
  );
}
