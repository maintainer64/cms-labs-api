import { Table, TableBody, TableCell, TableColumn, TableHeader, TableRow } from '@nextui-org/react';
import React from 'react';
import { RenderCellWithLocale } from './render-cell';
import useLanguageBrowser from '@/helpers/locale';
import { models_PNETServerListItem } from '@/helpers/api';
import InfiniteScroll from '@/components/scroll/infinity-scroll';

interface PnetServerTableWrapperProps {
  loadMore?: () => void;
  isLoading?: boolean;
  isInitialLoading?: boolean;
  rows?: models_PNETServerListItem[];
}

export const PnetServerTableWrapper = ({ rows, isLoading, loadMore }: PnetServerTableWrapperProps) => {
  const {
    locale: {
      Tables: { PnetServersTable }
    }
  } = useLanguageBrowser();
  const RenderCell = RenderCellWithLocale.bind(RenderCellWithLocale, useLanguageBrowser());
  return (
    <InfiniteScroll loadMore={loadMore} isLoading={isLoading}>
      <div className=' w-full flex flex-col gap-4'>
        <Table aria-label='Users table'>
          <TableHeader columns={PnetServersTable.Columns}>
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
