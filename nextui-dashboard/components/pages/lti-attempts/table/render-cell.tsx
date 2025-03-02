import React from 'react';
import { EditIcon } from '../../../icons/table/edit-icon';
import { models_LTIAttemptListItem } from '@/helpers/api';
import { Link } from 'react-router-dom';
import { RoutesLocation } from '@/components/routes';

interface Props {
  item: models_LTIAttemptListItem;
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
            <span>{item.user_name}</span>
          </div>
          <div>
            <Link to={RoutesLocation.accountsEdit(item.user_id?.toString())}>
              <span>{item.user_email}</span>
            </Link>
          </div>
        </div>
      );
    case 'server':
      return (
        <Link to={RoutesLocation.pnetServersEdit(item.pnet_server_id?.toString())}>
          <div>
            <div>
              <span>{item.pnet_server_name}</span>
            </div>
          </div>
        </Link>
      );
    case 'name':
      return (
        <Link to={RoutesLocation.ltiRoutingEdit(item.lti_routing_id?.toString())}>
          <div>
            <span>{item.lti_routing_name}</span>
          </div>
        </Link>
      );
    case 'actions':
      return (
        <div className='flex items-center gap-4 '>
          <div>
            <Link to={RoutesLocation.ltiAttemptEdit(item.id?.toString())}>
              <EditIcon size={20} fill='#979797' />
            </Link>
          </div>
        </div>
      );
    default:
      return '';
  }
};
