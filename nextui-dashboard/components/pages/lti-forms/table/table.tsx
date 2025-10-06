import { Table, TableBody, TableCell, TableColumn, TableHeader, TableRow } from '@heroui/react';
import React from 'react';
import { RenderCell } from './render-cell';
import useLanguageBrowser from '@/helpers/locale';
import { ModelsLTIFormListItem } from '@/helpers/api';
import InfiniteScroll from '@/components/scroll/infinity-scroll';
import { CamelCasedPropertiesDeep } from 'type-fest';

interface LTIFormsTableWrapperProps {
  loadMore?: () => void;
  isLoading?: boolean;
  isInitialLoading?: boolean;
  rows?: CamelCasedPropertiesDeep<ModelsLTIFormListItem>[];
}

export const LTIFormsTableWrapper = ({ rows, isLoading, loadMore }: LTIFormsTableWrapperProps) => {
  const {
    locale: {
      Tables: { LTIFormsTable }
    }
  } = useLanguageBrowser();
  return (
    <InfiniteScroll loadMore={loadMore} isLoading={isLoading}>
      <div className=' w-full flex flex-col gap-4'>
        <Table aria-label='Users table'>
          <TableHeader columns={LTIFormsTable.Columns}>
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
