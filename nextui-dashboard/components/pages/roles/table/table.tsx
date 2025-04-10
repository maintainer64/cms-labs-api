import { Table, TableBody, TableCell, TableColumn, TableHeader, TableRow } from '@heroui/react';
import React from 'react';
import { RenderCell } from './render-cell';
import useLanguageBrowser from '@/helpers/locale';
import { models_Role } from '@/helpers/api';
import InfiniteScroll from '@/components/scroll/infinity-scroll';

interface RolesTableWrapperProps {
  loadMore?: () => void;
  isLoading?: boolean;
  isInitialLoading?: boolean;
  rows?: models_Role[];
}

export const RolesTableWrapper = ({ rows, isLoading, loadMore }: RolesTableWrapperProps) => {
  const {
    locale: {
      Tables: { RoleTable }
    }
  } = useLanguageBrowser();
  return (
    <InfiniteScroll loadMore={loadMore} isLoading={isLoading}>
      <div className=' w-full flex flex-col gap-4'>
        <Table aria-label='Users table'>
          <TableHeader columns={RoleTable.Columns}>
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
          <TableBody items={rows ?? []}>
            {(item) => <TableRow>{(columnKey) => <TableCell>{RenderCell({ item, columnKey })}</TableCell>}</TableRow>}
          </TableBody>
        </Table>
      </div>
    </InfiniteScroll>
  );
};
