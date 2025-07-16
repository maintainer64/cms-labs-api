import { addToast, Navbar, NavbarContent, NavbarMenu, NavbarMenuToggle } from '@heroui/react';
import { useParamsConnectTopology } from '@/components/topology/utils';
import { RoutesLocation } from '@/components/routes';
import { useTopologyGet, useTopologyTokenJson } from '@/helpers/queries/topology/get';
import copy from 'copy-to-clipboard';
import { TopologyMenuItem } from '@/components/topology/menu/item';

export default function TopologyMenu() {
  const params = useParamsConnectTopology();
  const queryTopology = useTopologyGet(params.namespace);
  const queryTokenJson = useTopologyTokenJson();
  return (
    <div className='relative w-full'>
      <Navbar
        className='w-full'
        classNames={{
          wrapper: 'w-full max-w-full',
          menu: 'z-[100]'
        }}
        disableAnimation
      >
        <NavbarContent justify='start'>
          <NavbarMenuToggle />
        </NavbarContent>
        <NavbarMenu>
          <TopologyMenuItem
            title='Открыть логи топологии'
            href={queryTopology.data?.result?.web_url || '#'}
            target='_blank'
          />
          <TopologyMenuItem
            title='Перезапустить топологию'
            color='warning'
            href={RoutesLocation.topologyConnect() + `?username=${params.username}&taskId=${params.taskId}&redeploy=1`}
          />
          <TopologyMenuItem
            title='Скопировать токен доступа'
            onClick={() => {
              copy(queryTokenJson.data?.result?.token || '-');
              addToast({
                title: 'Скопировано',
                description: 'Токен доступа скопирован в буффер обмена',
                color: 'success',
                timeout: 9999999999999
              });
            }}
          />
          <TopologyMenuItem
            title='Подключиться к kubectl'
            onClick={() => {
              console.log('Подключиться к kubectl');
            }}
          />
          <TopologyMenuItem
            title='Удалить топологию'
            color='danger'
            onClick={() => {
              console.log('Удалить топологию');
            }}
          />
        </NavbarMenu>
      </Navbar>
    </div>
  );
}
