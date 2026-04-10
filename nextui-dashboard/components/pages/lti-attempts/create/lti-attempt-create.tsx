'use client';

import useLanguageBrowser from '@/helpers/locale';
import { InputOtp } from '@heroui/input-otp';
import React, { useEffect, useState } from 'react';
import { addToast, Button, Progress, Tooltip } from '@heroui/react';
import { Modal, ModalBody, ModalContent, ModalFooter, ModalHeader } from '@heroui/modal';
import copy from 'copy-to-clipboard';
import { RoutesLocation } from '@/components/routes';
import { ErrorModal } from '@/components/pages/auth/error';
import { Loading } from '@/components/scroll/loader';
import { Link } from 'react-router-dom';
import { AvatarGroupsRoom } from '@/components/pages/lti-attempts/create/avatar-groups-room';
import { useMutationLtiAttemptCreate } from '@/helpers/queries/lti_attempt/use-mutation-lti-attempt-create';
import { useQueryLtiAttemptCreate } from '@/helpers/queries/lti_attempt/use-query-lti-attempt-create';
import { ClipboardCopy, Undo2 } from 'lucide-react';

export const LTIAttemptCreate = () => {
  const {
    locale: {
      PnetServersQueue: { LTIAttemptRoom },
      Auth: { MainChangeLanguage }
    }
  } = useLanguageBrowser();
  const [roomNumber, setRoomNumber] = useState<string>('');
  const response = useQueryLtiAttemptCreate({});
  const result = response?.data;
  const change = useMutationLtiAttemptCreate({
    onSuccess: (data) => {
      addToast({
        title: LTIAttemptRoom.RoomChangeSuccess,
        color: 'success'
      });
    },
    onError: (error: any) => {
      addToast({
        title: LTIAttemptRoom.RoomChangeError,
        description: error.data.message,
        color: 'danger'
      });
    }
  });
  const defaultRoomNumber = result?.roomNumber?.toString() ?? '';
  useEffect(() => {
    setRoomNumber(defaultRoomNumber);
  }, [defaultRoomNumber]);
  if (response.isLoading) return <Loading size='md' />;
  if (response.error)
    return (
      <ErrorModal title={LTIAttemptRoom.ErrorPageTitle} description={response?.error?.data?.message || response?.error?.message}>
        <Button onPress={() => response.refetch()} href='#' variant='light' color='primary'>
          {LTIAttemptRoom.ErrorPageRefresh}
        </Button>
      </ErrorModal>
    );
  if (result?.autoRedirect === true) {
    window.location.href = result.nextUrl || '/';
    return <></>;
  }
  return (
    <Modal isOpen={true} hideCloseButton={true}>
      <ModalContent>
        <ModalHeader className='flex flex-col gap-1'>{LTIAttemptRoom.ModalTitle}</ModalHeader>
        <ModalBody>
          <AvatarGroupsRoom members={result?.members} collaboration={result?.collaboration} />
          <p>
            {LTIAttemptRoom.ModalDescription}
            <br />
            {LTIAttemptRoom.ModalDescriptionChange}
          </p>
          <div className='flex flex-col items-start gap-2'>
            <div className='flex items-center justify-center gap-4'>
              <InputOtp id='LTIAttemptCreateId' length={5} value={roomNumber} onValueChange={setRoomNumber} />
              {roomNumber === defaultRoomNumber && (
                <Tooltip showArrow={true} content={LTIAttemptRoom.ButtonCopy}>
                  <Button
                    isIconOnly
                    color='default'
                    variant='light'
                    onPress={() => {
                      copy(roomNumber);
                    }}
                    disableRipple={true}
                    className='cursor-pointer'
                  >
                    <ClipboardCopy className='w-4 h-4 stroke-[#969696]' />
                  </Button>
                </Tooltip>
              )}
              {roomNumber !== defaultRoomNumber && (
                <Tooltip showArrow={true} content={LTIAttemptRoom.ButtonRollback}>
                  <Button
                    isIconOnly
                    color='default'
                    variant='light'
                    onPress={() => {
                      setRoomNumber(defaultRoomNumber);
                    }}
                    disableRipple={true}
                    className='cursor-pointer'
                  >
                    <Undo2 className='w-4 h-4 stroke-[#969696]' />
                  </Button>
                </Tooltip>
              )}
            </div>
          </div>
          <div className='text-small text-default-500 cursor-pointer'>
            <Link to={RoutesLocation.language()}>{MainChangeLanguage}</Link>
          </div>
        </ModalBody>
        <ModalFooter>
          <Button
            onPress={() => {
              change.mutate({ roomNumber: parseInt(roomNumber) });
            }}
            isDisabled={roomNumber === defaultRoomNumber}
            color='secondary'
          >
            {LTIAttemptRoom.ButtonChange}
          </Button>
        </ModalFooter>
        <Progress isIndeterminate className='max-w-md' size='sm' />
      </ModalContent>
    </Modal>
  );
};
