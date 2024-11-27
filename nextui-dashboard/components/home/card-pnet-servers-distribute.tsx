import React from 'react';
import Chart, { Props } from 'react-apexcharts';
import { usePnetServerQueueList } from '@/helpers/queries/pnet-server-queue/get';
import { Loading } from '@/components/scroll/loader';
import { queries_RoundQueuePoolPnetListItem } from '@/helpers/api';

interface serversSeriesProps {
  name: string;
  data: number[][];
}

const serversName = (model: Array<queries_RoundQueuePoolPnetListItem>): [serversSeriesProps[], number] => {
  const series: serversSeriesProps[] = [];
  const seriesLabel: Set<number> = new Set();
  const lastUsedServer = model.find((server) => server.last_used);
  model.forEach((server) => {
    if (seriesLabel.has(server?.id || 0)) return;
    series.push({
      name: server.name || `#${server.id}`,
      // @ts-expect-error: return nullable value
      data: model
        .map((distributionServer, index) => (server.id === distributionServer.id ? [index, 1] : null))
        .filter((x) => x !== null)
    });
    seriesLabel.add(server?.id || 0);
  });
  if (!lastUsedServer) return [series, seriesLabel.size];
  series.push({
    name: `Last ${lastUsedServer?.name}`,
    // @ts-expect-error: return nullable value
    data: model
      .map((distributionServer, index) => (distributionServer.last_used ? [index, 1.05] : null))
      .filter((x) => x !== null)
  });
  return [series, seriesLabel.size];
};
export const CardPnetServersDistribute = () => {
  const response = usePnetServerQueueList();
  if (response.isLoading) return <Loading size={8} />;
  const servers = response.data?.result?.model || [];
  const [series, seriesSize] = serversName(servers);
  const shapes = Array(seriesSize + 1)
    .fill(undefined)
    .map((_, index) => (index === seriesSize ? 'diamond' : 'circle'));
  const chartData: Props = {
    type: 'scatter',
    series: series,
    options: {
      markers: {
        size: 12,
        strokeWidth: 2,
        hover: {
          size: 12
        },
        shape: shapes
      },
      tooltip: {
        enabled: true
      },
      chart: {
        zoom: {
          enabled: true
        },
        toolbar: {
          show: true
        },
        id: 'basic-bar',
        foreColor: 'hsl(var(--nextui-default-800))'
      },
      xaxis: {
        min: 0,
        labels: {
          // show: false,
          style: {
            colors: 'hsl(var(--nextui-default-800))'
          }
        },
        axisBorder: {
          color: 'hsl(var(--nextui-nextui-default-200))'
        },
        axisTicks: {
          color: 'hsl(var(--nextui-nextui-default-200))'
        }
      },
      yaxis: {
        max: 1.1,
        min: 0.9,
        labels: {
          style: {
            // hsl(var(--nextui-content1-foreground))
            colors: 'hsl(var(--nextui-default-800))'
          }
        }
      },
      grid: {
        show: true,
        borderColor: 'hsl(var(--nextui-default-200))',
        strokeDashArray: 0,
        position: 'back'
      },
      stroke: {
        curve: 'smooth',
        fill: {
          colors: ['red']
        }
      }
    }
  };

  return <Chart {...chartData} />;
};

export default CardPnetServersDistribute;
