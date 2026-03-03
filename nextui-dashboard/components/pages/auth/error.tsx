import useLanguageBrowser from '@/helpers/locale';
import { Modal, ModalBody, ModalContent, ModalHeader } from '@heroui/modal';
import { ModalFooter } from '@heroui/react';
import React from 'react';

type ErrorModalProps = {
  title: string;
  description: string;
  children?: React.ReactNode;
};
export const ErrorModal = ({ title, description, children }: ErrorModalProps) => {
  return (
    <Modal isOpen={true} hideCloseButton={true}>
      <ModalContent>
        <ModalHeader className='flex flex-col gap-1'>{title}</ModalHeader>
        <ModalBody>
          <p>{description}</p>
        </ModalBody>
        <ModalFooter>{children}</ModalFooter>
      </ModalContent>
    </Modal>
  );
};
export default function AuthError() {
  const { locale } = useLanguageBrowser();
  return <ErrorModal title={locale.Auth.ErrorPageTitle} description={locale.Auth.ErrorPageDescription} />;
}
