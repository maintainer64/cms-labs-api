import { Table, TableBody, TableCell, TableColumn, TableHeader, TableRow } from '@nextui-org/react';
import React from 'react';
import { RenderCell } from './render-cell';
import useLanguageBrowser from '@/helpers/locale';
import { models_LTIRoutingListItem } from '@/helpers/api';
import InfiniteScroll from '@/components/scroll/infinity-scroll';

interface LTIFormsTableWrapperProps {
  loadMore?: () => void;
  isLoading?: boolean;
  isInitialLoading?: boolean;
  rows?: models_LTIRoutingListItem[];
}

export const LTIRoutingTableWrapper = ({ rows, isLoading, loadMore }: LTIFormsTableWrapperProps) => {
  const {
    locale: {
      Tables: { LTIRoutingTable }
    }
  } = useLanguageBrowser();
  return (
    <InfiniteScroll loadMore={loadMore} isLoading={isLoading}>
      <div className=' w-full flex flex-col gap-4'>
        <Table aria-label='Users table'>
          <TableHeader columns={LTIRoutingTable.Columns}>
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
