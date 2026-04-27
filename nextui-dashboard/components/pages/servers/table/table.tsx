import { Table, TableBody, TableCell, TableColumn, TableHeader, TableRow } from '@heroui/react';
import React from 'react';
import { RenderCellWithLocale } from './render-cell';
import useLanguageBrowser from '@/helpers/locale';
import InfiniteScroll from '@/components/scroll/infinity-scroll';
import { ServerItem } from '@/helpers/queries/server/use-query-server-get';

interface ServerTableWrapperProps {
  loadMore?: () => void;
  isLoading?: boolean;
  isInitialLoading?: boolean;
  rows?: ServerItem[];
}

export const ServerTableWrapper = ({ rows, isLoading, loadMore }: ServerTableWrapperProps) => {
  const {
    locale: {
      Tables: { ServersTable }
    }
  } = useLanguageBrowser();
  const RenderCell = RenderCellWithLocale.bind(RenderCellWithLocale, useLanguageBrowser());
  return (
    <InfiniteScroll loadMore={loadMore} isLoading={isLoading}>
      <div className=' w-full flex flex-col gap-4'>
        <Table aria-label='Users table'>
          <TableHeader columns={ServersTable.Columns}>
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
            {(item) => (
              <TableRow>
                {(columnKey) => (
                  <TableCell>
                    {RenderCell({
                      item,
                      columnKey
                    })}
                  </TableCell>
                )}
              </TableRow>
            )}
          </TableBody>
        </Table>
      </div>
    </InfiniteScroll>
  );
};
