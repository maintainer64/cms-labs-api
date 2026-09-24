import React, { useCallback } from 'react';
import { TerminalActionFunc } from '@/components/topology/terminal/service/context';
import { addToast, Tooltip } from '@heroui/react';
import useLanguageBrowser from '@/helpers/locale';
import { useMutationNodeAction } from '@/helpers/queries/node/use-mutation-node-action';
import { useParams } from 'react-router-dom';

export interface RFContextMenuProps {
  id?: string;
  top?: number | false;
  left?: number | false;
  right?: number | false;
  bottom?: number | false;
  dispatch?: TerminalActionFunc;
  onClick?: () => void;
}

interface MenuItem {
  label: string;
  tooltip: string;
  onClick: () => void;
}

export function RFContextMenu({ id, top, left, right, bottom, dispatch, onClick }: RFContextMenuProps) {
  if (!id) return <></>;
  const sessionId = useParams().sessionId || '';
  const {
    locale: {
      Topology: { Modal }
    }
  } = useLanguageBrowser();
  const { mutate } = useMutationNodeAction({
    onSuccess: (data) => {
      if ((data?.count || 0) > 0) {
        addToast({
          title: id,
          description: Modal.Success,
          color: 'success'
        });
      } else {
        addToast({
          title: id,
          description: Modal.Error,
          color: 'danger'
        });
      }
    },
    onError: (error: any) => {
      addToast({
        title: id,
        description: error.data.message || Modal.Error,
        color: 'danger'
      });
    }
  });

  const openTerminal = useCallback(() => {
    dispatch?.({
      type: 'ADD_CLIENT',
      payload: {
        id: id
      }
    });
    dispatch?.({
      type: 'TOGGLE_VISIBILITY',
      payload: {
        id: id,
        onChange: (visibility?: boolean) => {
          if (visibility) {
            dispatch?.({
              type: 'BRING_TO_FRONT',
              payload: {
                id: id
              }
            });
          }
        }
      }
    });
  }, [id, dispatch]);
  const restart = useCallback(() => {
    mutate?.({
      sessionId,
      actions: [
        {
          node: id,
          action: 'restart'
        }
      ]
    });
  }, [id, mutate, sessionId]);

  const wipe = useCallback(() => {
    mutate?.({
      sessionId,
      actions: [
        {
          node: id,
          action: 'wipe'
        }
      ]
    });
  }, [id, mutate, sessionId]);

  const menuItems: MenuItem[] = [
    {
      label: Modal.Terminal,
      tooltip: Modal.TerminalTooltip,
      onClick: openTerminal
    },
    {
      label: Modal.Reboot,
      tooltip: Modal.RebootTooltip,
      onClick: restart
    },
    {
      label: Modal.HardReset,
      tooltip: Modal.HardResetTooltip,
      onClick: wipe
    }
  ];

  return (
    <div
      // @ts-ignore
      style={{ top, left, right, bottom }}
      className='absolute z-50 min-w-[180px] rounded-lg border border-gray-200 bg-white py-2 shadow-xl dark:border-gray-700 dark:bg-gray-800'
      onClick={onClick}
    >
      <p className='border-b border-gray-100 px-4 py-2 text-sm font-medium text-gray-500 dark:border-gray-700 dark:text-gray-400'>
        {id}
      </p>
      {menuItems.map((item, index) => (
        <Tooltip key={index} content={item.tooltip} placement='right-start'>
          <button
            key={index}
            onClick={item.onClick}
            className='flex w-full items-center gap-3 px-4 py-2.5 text-left text-sm text-gray-700 transition-colors hover:bg-gray-100 dark:text-gray-200 dark:hover:bg-gray-700'
          >
            {item.label}
          </button>
        </Tooltip>
      ))}
    </div>
  );
}
