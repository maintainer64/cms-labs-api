import { addToast, Navbar, NavbarContent, NavbarMenu, NavbarMenuToggle } from '@heroui/react';
import { TopologyMenuItem } from '@/components/topology/menu/item';
import useLanguageBrowser from '@/helpers/locale';
import { useConfirmPopup } from '@/components/hooks/useDeletePopup';
import { useNavigate, useParams } from 'react-router-dom';
import { zIndexClassMenu } from '@/components/providers/const';
import { NavbarDarkModeToggle } from '@/components/navbar/darkiconswitch';
import { useQueryTopologyGet } from '@/helpers/queries/topology/use-query-topology-get';
import { useMutationNodeAction } from '@/helpers/queries/node/use-mutation-node-action';
import { UsecasesNodeActionItem } from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';
import { useCallback } from 'react';
import { RoutesLocation } from '@/components/routes';

export default function TopologyMenu() {
  const {
    locale: {
      Topology: { Menu },
      Auth: { MainChangeLanguage }
    }
  } = useLanguageBrowser();

  const navigate = useNavigate();

  const handleNavigateLanguage = useCallback(() => {
    navigate(RoutesLocation.language());
  }, [navigate]);

  const { username, attemptNumber, sessionId } = useParams();
  const queryTopology = useQueryTopologyGet({ username, attemptNumber, sessionId });

  const createAction = (action: string) => {
    return {
      username,
      attemptNumber,
      sessionId,
      actions: queryTopology?.data?.deployments?.map((item) => {
        return {
          action: action,
          node: item?.name || '#'
        } as CamelCasedPropertiesDeep<UsecasesNodeActionItem>;
      })
    };
  };

  const { mutate } = useMutationNodeAction({
    onSuccess: (data) => {
      addToast({
        title: Menu.Topology,
        description: Menu.Success,
        color: 'success'
      });
      // Костыль закрыватель попапов и жёсткая перезагрузка топологии
      setTimeout(() => {
        location.reload();
      }, 500);
    },
    onError: (error: any) => {
      addToast({
        title: Menu.Topology,
        description: error.data.message || Menu.Error,
        color: 'danger'
      });
    }
  });

  const restartPopup = useConfirmPopup({
    title: Menu.RestartAllDevicesTitle,
    description: Menu.RestartAllDevicesDescription,
    onConfirm: () => {
      mutate(createAction('restart'));
    }
  });

  const wipePopup = useConfirmPopup({
    title: Menu.HardResetAllDevicesTitle,
    description: Menu.HardResetAllDevicesDescription,
    onConfirm: () => {
      mutate(createAction('wipe'));
    }
  });

  return (
    <>
      <div className='relative w-full'>
        <Navbar
          className='w-full'
          classNames={{
            wrapper: 'w-full max-w-full',
            menu: zIndexClassMenu
          }}
          disableAnimation
        >
          <NavbarContent justify='start'>
            <NavbarMenuToggle />
          </NavbarContent>
          <NavbarContent justify='end'>
            <NavbarDarkModeToggle />
          </NavbarContent>
          <NavbarMenu>
            <TopologyMenuItem title={Menu.RestartAllDevicesTitle} color='warning' onClick={restartPopup.onOpen} />
            <TopologyMenuItem title={Menu.HardResetAllDevicesTitle} color='danger' onClick={wipePopup.onOpen} />
            <TopologyMenuItem title={MainChangeLanguage} onClick={handleNavigateLanguage} />
          </NavbarMenu>
        </Navbar>
      </div>
      {restartPopup.component({})}
      {wipePopup.component({})}
    </>
  );
}
