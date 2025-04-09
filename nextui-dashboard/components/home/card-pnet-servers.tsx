import React, { useState } from 'react';
import { Select, SelectItem } from '@heroui/react';
import { usePnetServerList } from '@/helpers/queries/pnet-server/get';
import Chart, { Props } from 'react-apexcharts';
import useLanguageBrowser from '@/helpers/locale';
import { Loading } from '@/components/scroll/loader';

export const CardPnetServers = () => {
  const {
    locale: {
      PnetServers: { StatsChart }
    }
  } = useLanguageBrowser();
  const [orderBy, setOrderBy] = useState('unitRate');
  const [status, setStatus] = useState('active');
  const response = usePnetServerList({
    limit: 100,
    order_by: orderBy,
    status: status
  });

  if (response.isLoading) return <Loading size='md' />;

  const rows = response?.data?.pages.flatMap((p) => p.result?.model ?? []) || [];

  const chartData: Props = {
    type: 'pie',
    series: rows?.map((server) => {
      if (orderBy === 'lastCountUsers') return server.model.last_count_users || 0;
      return server.model.unit_rate || 0;
    }),
    options: {
      labels: rows?.map((server) => {
        return server.model.name || `#${server.model.id}`;
      })
    }
  };

  return (
    <div>
      <div className='flex justify-between gap-4 mb-4'>
        <Select
          variant='bordered'
          label={StatsChart.Status}
          selectedKeys={[status ?? '']}
          onSelectionChange={(keys) => setStatus(keys.currentKey || '')}
        >
          <SelectItem key='active'>{StatsChart.StatusValueActive}</SelectItem>
          <SelectItem key='all'>{StatsChart.StatusValueAll}</SelectItem>
        </Select>
        <Select
          variant='bordered'
          label={StatsChart.OrderBy}
          selectedKeys={[orderBy ?? '']}
          onSelectionChange={(keys) => setOrderBy(keys.currentKey || '')}
        >
          <SelectItem key='unitRate'>{StatsChart.OrderByValueUnitRate}</SelectItem>
          <SelectItem key='lastCountUsers'>{StatsChart.OrderByValueLastCountUsers}</SelectItem>
          <SelectItem key='createdAt'>{StatsChart.OrderByValueCreatedAt}</SelectItem>
        </Select>
      </div>
      <Chart {...chartData} />
    </div>
  );
};

export default CardPnetServers;
