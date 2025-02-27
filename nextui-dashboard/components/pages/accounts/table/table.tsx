import { Table, TableBody, TableCell, TableColumn, TableHeader, TableRow } from '@heroui/react';
import React from 'react';
import { RenderCell } from './render-cell';
import useLanguageBrowser from '@/helpers/locale';
import { models_UserListItem } from '@/helpers/api';
import InfiniteScroll from '@/components/scroll/infinity-scroll';

interface UsersTableWrapperProps {
  loadMore?: () => void;
  isLoading?: boolean;
  isInitialLoading?: boolean;
  users?: models_UserListItem[];
}

export const UsersTableWrapper = ({ users, isLoading, loadMore }: UsersTableWrapperProps) => {
  const {
    locale: {
      Tables: { UsersTable }
    }
  } = useLanguageBrowser();
  return (
    <InfiniteScroll loadMore={loadMore} isLoading={isLoading}>
      <div className=' w-full flex flex-col gap-4'>
        <Table aria-label='Users table'>
          <TableHeader columns={UsersTable.Columns}>
            {(column) => (
              <TableColumn
                key={column.uid}
                hideHeader={column.uid === 'actions'}
                align={column.uid === 'actions' ? 'center' : 'start'}
              >
                {column.name}
              </TableColumn>
            )}
          </TableHeader>
          <TableBody items={users ?? []}>
            {(item) => <TableRow>{(columnKey) => <TableCell>{RenderCell({ item, columnKey })}</TableCell>}</TableRow>}
          </TableBody>
        </Table>
      </div>
    </InfiniteScroll>
  );
};
