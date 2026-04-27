'use client';

import useLanguageBrowser from '@/helpers/locale';
import React from 'react';
import { Button } from '@heroui/react';
import { Modal, ModalBody, ModalContent, ModalFooter, ModalHeader } from '@heroui/modal';
import { Link } from 'react-router-dom';
import { RoutesLocation } from '@/components/routes';

export const LTIAttemptConfirm = () => {
  const {
    locale: {
      ServersQueue: { LTIAttemptRoom, LTIAttemptConfirm }
    }
  } = useLanguageBrowser();
  return (
    <Modal isOpen={true} hideCloseButton={true}>
      <ModalContent>
        <ModalHeader className='flex flex-col gap-1'>{LTIAttemptRoom.ModalTitle}</ModalHeader>
        <ModalBody>
          <p>{LTIAttemptConfirm.ModalDescription}</p>
        </ModalBody>
        <ModalFooter>
          <Button onPress={() => {}} as={Link} to={RoutesLocation.home()} variant='light' color='danger'>
            {LTIAttemptConfirm.ButtonAdminPanel}
          </Button>
          <Button as={Link} to={RoutesLocation.ltiRedirectCreate()} color='primary'>
            {LTIAttemptConfirm.ButtonLab}
          </Button>
        </ModalFooter>
      </ModalContent>
    </Modal>
  );
};
