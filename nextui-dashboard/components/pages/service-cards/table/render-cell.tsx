import React from 'react';
import { ModelsServiceCardListItem } from '@/helpers/api';
import { Link } from 'react-router-dom';
import { RoutesLocation } from '@/components/routes';
import { Chip } from '@heroui/react';
import { CamelCasedPropertiesDeep } from 'type-fest';
import { SquarePen } from 'lucide-react';

interface Props {
  item: CamelCasedPropertiesDeep<ModelsServiceCardListItem>;
  columnKey: string | React.Key;
}

export const RenderCellWithLocale = (locale: any, { item, columnKey }: Props) => {
  const {
    locale: {
      Tables: {
        ServiceCardsTable: { ColumnStatus }
      }
    }
  } = locale;
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
            <span>{item.description}</span>
          </div>
        </div>
      );
    case 'order':
      return (
        <div>
          <div>
            <span>{item.order}</span>
          </div>
        </div>
      );
    case 'is_active':
      return (
        <div>
          <div>
            {item.isActive ? (
              <Chip size='sm' color='success'>
                {ColumnStatus.Activated}
              </Chip>
            ) : (
              <Chip size='sm' color='danger'>
                {ColumnStatus.Deactivated}
              </Chip>
            )}
          </div>
        </div>
      );
    case 'actions':
      return (
        <div className='flex items-center gap-4 '>
          <div>
            <Link to={RoutesLocation.serviceCardsEdit(item.id?.toString())}>
              <SquarePen className='w-4 p-4 stroke-[#979797]' />
            </Link>
          </div>
        </div>
      );
    default:
      return '';
  }
};
