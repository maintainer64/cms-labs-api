import { addToast, Navbar, NavbarContent, NavbarMenu, NavbarMenuToggle } from '@heroui/react';
import { useParamsConnectTopology } from '@/components/topology/utils';
import { RoutesLocation } from '@/components/routes';
import { useTopologyGet, useTopologyTokenJson } from '@/helpers/queries/topology/get';
import copy from 'copy-to-clipboard';
import { TopologyMenuItem } from '@/components/topology/menu/item';
import useLanguageBrowser from '@/helpers/locale';
import { useConfirmPopup } from '@/components/hooks/useDeletePopup';
import { useTopologyDelete } from '@/helpers/queries/topology/delete';
import { useNavigate } from 'react-router-dom';
import { zIndexClassMenu } from '@/components/providers/const';
import { useCallback } from 'react';
import { NavbarDarkModeToggle } from '@/components/navbar/darkiconswitch';

export default function TopologyMenu() {
  const {
    locale: {
      Topology: { Menu },
      Auth: { MainChangeLanguage }
    }
  } = useLanguageBrowser();
  const navigate = useNavigate();
  const params = useParamsConnectTopology();
  const queryTopology = useTopologyGet(params.namespace);
  const queryTokenJson = useTopologyTokenJson();

  const handleNavigateKubectlDocs = useCallback(() => {
    navigate(RoutesLocation.docsTopologyKubectl() + `?namespace=${params.namespace}`);
  }, [params, navigate]);

  const handleNavigateLanguage = useCallback(() => {
    navigate(RoutesLocation.language());
  }, [navigate]);

  const removeTopologyMutation = useTopologyDelete({
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
            <TopologyMenuItem title={Menu.OpenLogs} href={queryTopology.data?.result?.web_url || '#'} target='_blank' />
            <TopologyMenuItem title={Menu.RestartTopology} color='warning' onClick={restartTopologyPopup.onOpen} />
            <TopologyMenuItem
              title={Menu.CopyToken}
              onClick={() => {
                copy(queryTokenJson.data?.result?.token || '-');
                addToast({
                  title: Menu.CopyTokenModal.Title,
                  description: Menu.CopyTokenModal.Description,
                  color: 'success'
                });
              }}
            />
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
