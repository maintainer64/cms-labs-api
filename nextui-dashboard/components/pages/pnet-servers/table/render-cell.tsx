import React from 'react';
import { Chip } from '@heroui/react';
import dayjs from 'dayjs';
import { RolesChip } from '@/components/base-forms/roles';
import { Link } from 'react-router-dom';
import { RoutesLocation } from '@/components/routes';
import { EditIcon } from '@/components/icons/table/edit-icon';
import { PnetServerItem } from '@/helpers/queries/server/use-query-server-get';

interface Props {
  item: PnetServerItem;
  columnKey: string | React.Key;
}

const isConnectDistribution = (lastOnlineStatus?: string, minutesForDisconnect?: number): boolean => {
  if (!minutesForDisconnect) return true;
  if (!lastOnlineStatus) return false;
  const now = dayjs();
  const lastOnline = dayjs(lastOnlineStatus);
  return lastOnline.add(minutesForDisconnect, 'minute').isAfter(now);
};

export const RenderCellWithLocale = (locale: any, { item, columnKey }: Props) => {
  const {
    locale: {
      Tables: {
        PnetServersTable: { ColumnStatus, ColumnIndicator }
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
            <span>{item.url}</span>
          </div>
        </div>
      );
    case 'indicator':
      return (
        <div>
          {item.type === 'pnet' && (
            <>
              <div>
                <span>
                  {ColumnIndicator.UnitRate}: {item.unitRate}%
                </span>
              </div>
              <div>
                <span>
                  {ColumnIndicator.LastCountUsers}: {item.lastCountUsers}
                </span>
              </div>
            </>
          )}
        </div>
      );
    case 'status':
      return (
        <div className='flex flex-col gap-2'>
          {item.type === 'pnet' && (
            <>
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
              <div>
                {isConnectDistribution(item.lastOnlineStatus, item.minutesForDisconnect) ? (
                  <Chip size='sm' color='success'>
                    {ColumnStatus.ConnectDistribution}
                  </Chip>
                ) : (
                  <Chip size='sm' color='danger'>
                    {ColumnStatus.DisconnectDistribution}
                  </Chip>
                )}
              </div>
            </>
          )}
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
        <div className='flex items-center gap-4 '>
          <div>
            <Link to={RoutesLocation.pnetServersEdit(item.id?.toString())}>
              <EditIcon size={20} fill='#979797' />
            </Link>
          </div>
        </div>
      );
    default:
      return '';
  }
};
