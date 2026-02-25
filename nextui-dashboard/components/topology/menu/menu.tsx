import { addToast, Navbar, NavbarContent, NavbarMenu, NavbarMenuToggle } from '@heroui/react';
import { useParamsConnectTopology } from '@/components/topology/utils';
import { RoutesLocation } from '@/components/routes';
import { TopologyMenuItem } from '@/components/topology/menu/item';
import useLanguageBrowser from '@/helpers/locale';
import { useConfirmPopup } from '@/components/hooks/useDeletePopup';
import { useNavigate } from 'react-router-dom';
import { zIndexClassMenu } from '@/components/providers/const';
import { useCallback } from 'react';
import { NavbarDarkModeToggle } from '@/components/navbar/darkiconswitch';
import { useQueryTopologyGet } from '@/helpers/queries/topology/use-query-topology-get';
import { useMutationTopologyDelete } from '@/helpers/queries/topology/use-mutation-topology-delete';

export default function TopologyMenu() {
  const {
    locale: {
      Topology: { Menu },
      Auth: { MainChangeLanguage }
    }
  } = useLanguageBrowser();
  const navigate = useNavigate();
  const params = useParamsConnectTopology();
  const queryTopology = useQueryTopologyGet({ namespace: params.namespace });

  const handleNavigateKubectlDocs = useCallback(() => {
    navigate(RoutesLocation.docsTopologyKubectl() + `?namespace=${params.namespace}`);
  }, [params, navigate]);

  const handleNavigateLanguage = useCallback(() => {
    navigate(RoutesLocation.language());
  }, [navigate]);

  const removeTopologyMutation = useMutationTopologyDelete({
    onSuccess: () => {
      navigate(RoutesLocation.home(), { replace: true });
      addToast({
        title: Menu.RemoveTopologyTitle,
        description: Menu.RemoveTopologyDescriptionSuccess,
        color: 'success'
      });
    },
    onError: (error: any) => {
      addToast({
        title: Menu.RemoveTopologyTitle,
        description: error.body.msg,
        color: 'danger'
      });
    }
  });

  const removeTopologyPopup = useConfirmPopup({
    title: Menu.RemoveTopologyTitle,
    description: Menu.RemoveTopologyDescription,
    onConfirm: removeTopologyMutation.mutate.bind(removeTopologyMutation.mutate, { namespaces: [params.namespace] })
  });

  const restartTopologyPopup = useConfirmPopup({
    title: Menu.RestartTopologyTitle,
    description: Menu.RestartTopologyDescription,
    onConfirm: () => {
      navigate(RoutesLocation.topologyConnect() + `?username=${params.username}&taskId=${params.taskId}&redeploy=1`, {
        replace: true
      });
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
            <TopologyMenuItem title={Menu.OpenLogs} href={queryTopology.data?.webUrl || '#'} target='_blank' />
            <TopologyMenuItem title={Menu.RestartTopology} color='warning' onClick={restartTopologyPopup.onOpen} />
            <TopologyMenuItem title={Menu.ConnectToKubectl} onClick={handleNavigateKubectlDocs} />
            <TopologyMenuItem title={Menu.RemoveTopology} color='danger' onClick={removeTopologyPopup.onOpen} />
            <TopologyMenuItem title={MainChangeLanguage} onClick={handleNavigateLanguage} />
          </NavbarMenu>
        </Navbar>
      </div>
      {removeTopologyPopup.component({})}
      {restartTopologyPopup.component({})}
    </>
  );
}
