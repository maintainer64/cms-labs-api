import React from 'react';
import { ModelsLTIAttemptListItem } from '@/helpers/api';
import { Link } from 'react-router-dom';
import { RoutesLocation } from '@/components/routes';
import { CamelCasedPropertiesDeep } from 'type-fest';
import { SquarePen } from 'lucide-react';

interface Props {
  item: CamelCasedPropertiesDeep<ModelsLTIAttemptListItem>;
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
    case 'user':
      return (
        <div>
          <div>
            <span>{item.userName}</span>
          </div>
          <div>
            <Link to={RoutesLocation.accountsEdit(item.userId?.toString())}>
              <span>{item.userEmail}</span>
            </Link>
          </div>
        </div>
      );
    case 'server':
      return (
        <Link to={RoutesLocation.pnetServersEdit(item.pnetServerId?.toString())}>
          <div>
            <div>
              <span>{item.pnetServerName}</span>
            </div>
          </div>
        </Link>
      );
    case 'name':
      return (
        <Link to={RoutesLocation.ltiRoutingEdit(item.ltiRoutingId?.toString())}>
          <div>
            <span>{item.ltiRoutingName}</span>
          </div>
        </Link>
      );
    case 'actions':
      return (
        <div className='flex items-center gap-4 '>
          <div>
            <Link to={RoutesLocation.ltiAttemptEdit(item.id?.toString())}>
              <SquarePen className='w-4 p-4 stroke-[#979797]' />
            </Link>
          </div>
        </div>
      );
    default:
      return '';
  }
};
