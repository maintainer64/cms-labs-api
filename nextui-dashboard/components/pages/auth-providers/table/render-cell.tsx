import React from 'react';
import { ModelsAuthProviderListItem } from '@/helpers/api';
import { Link } from 'react-router-dom';
import { RoutesLocation } from '@/components/routes';
import { CamelCasedPropertiesDeep } from 'type-fest';
import { SquarePen } from 'lucide-react';

interface Props {
  item: CamelCasedPropertiesDeep<ModelsAuthProviderListItem>;
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
        </div>
      );
    case 'actions':
      return (
        <div className='flex items-center gap-4 '>
          <div>
            <Link to={RoutesLocation.authProvidersEdit(item.id?.toString())}>
              <SquarePen className='w-5 h-5 stroke-[#969696]' />
            </Link>
          </div>
        </div>
      );
    default:
      return '';
  }
};
