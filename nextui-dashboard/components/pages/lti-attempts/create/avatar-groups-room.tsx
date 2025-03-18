import { models_UserListItem } from '@/helpers/api';
import { AvatarGroup } from '@heroui/react';
import CustomAvatar from '@/components/sidebar/avatar';
import React from 'react';

interface AvatarGroupsRoomProps {
  members?: Array<models_UserListItem>;
  collaboration?: number;
}

export const AvatarGroupsRoom = ({ members, collaboration }: AvatarGroupsRoomProps) => {
  const emptyAvatar = new Array((collaboration || 0) - (members?.length || 0)).fill(undefined);
  emptyAvatar.map(() => console.log('ddd'));
  console.log();
  return (
    <div className='mx-4'>
      <AvatarGroup isBordered>
        {members?.map((member) => {
          return CustomAvatar({
            tooltip: true,
            as: 'button',
            size: 'md',
            username: member.name,
            name: member.name,
            email: member.email
          });
        })}
        {emptyAvatar?.map(() => {
          return CustomAvatar({
            as: 'button',
            tooltip: false,
            size: 'md',
            username: '+',
            name: '+',
            email: '+'
          });
        })}
      </AvatarGroup>
    </div>
  );
};
