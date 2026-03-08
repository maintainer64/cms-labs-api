import React from 'react';
import { Link } from 'react-router-dom';
import { RoutesLocation } from '@/components/routes';
import { RolesChip } from '@/components/base-forms/roles';
import { UserItem } from '@/helpers/queries/user/use-infinity-user-list';
import { SquarePen } from 'lucide-react';

interface Props {
  item: UserItem;
  columnKey: string | React.Key;
}

export const RenderCell = ({ item, columnKey }: Props) => {
  switch (columnKey) {
    case 'id':
      return (
        <div>
          <div>
            <span>#{item.id ?? 0}</span>
          </div>
        </div>
      );
    case 'name':
      return (
        <div>
          <div>
            <span>{item.name}</span>
          </div>
          <div>
            <span>{item.email}</span>
          </div>
        </div>
      );
    case 'role':
      return (
        <div className='flex items-center gap-4'>
          <RolesChip roles={item.roles} maxRoles={3} />
        </div>
      );
    case 'actions':
      return (
        <div className='flex items-center gap-4'>
          <div>
            <Link to={RoutesLocation.accountsEdit(item.id?.toString())}>
              <SquarePen className='w-5 h-5 stroke-[#969696]' />
            </Link>
          </div>
        </div>
      );
    default:
      return '';
  }
};
