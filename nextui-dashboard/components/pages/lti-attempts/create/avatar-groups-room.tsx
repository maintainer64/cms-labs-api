import { ModelsUserListItem } from '@/helpers/api';
import { AvatarGroup } from '@heroui/react';
import CustomAvatar from '@/components/sidebar/avatar';
import React from 'react';
import { CamelCasedPropertiesDeep } from 'type-fest';

interface AvatarGroupsRoomProps {
  members?: Array<CamelCasedPropertiesDeep<ModelsUserListItem>>;
  collaboration?: number;
}

export const AvatarGroupsRoom = ({ members, collaboration }: AvatarGroupsRoomProps) => {
  const emptyAvatar = new Array((collaboration || 0) - (members?.length || 0)).fill(undefined);
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
