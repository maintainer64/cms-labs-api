import React from 'react';
import { EditIcon } from '../../../icons/table/edit-icon';
import { ModelsLTIFormListItem } from '@/helpers/api';
import { Link } from 'react-router-dom';
import { RoutesLocation } from '@/components/routes';
import { CamelCasedPropertiesDeep } from 'type-fest';

interface Props {
  item: CamelCasedPropertiesDeep<ModelsLTIFormListItem>;
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
            <Link to={RoutesLocation.ltiFormsEdit(item.id?.toString())}>
              <EditIcon size={20} fill='#979797' />
            </Link>
          </div>
        </div>
      );
    default:
      return '';
  }
};
