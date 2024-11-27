import useLanguageBrowser from '@/helpers/locale';
import { Button, Input, useDisclosure } from '@nextui-org/react';
import { Modal, ModalBody, ModalContent, ModalFooter, ModalHeader } from '@nextui-org/modal';
import React, { useState } from 'react';

export const LtiFormURILTIMoodle = () => {
  const { isOpen, onOpen, onOpenChange } = useDisclosure();
  const component = (onChange: (url: string) => void) => {
    const {
      locale: { LTIForm, Sidebar }
    } = useLanguageBrowser();
    const [baseUrl, setBaseUrl] = useState<string>('');
    return (
      <Modal isOpen={isOpen} onOpenChange={onOpenChange}>
        <ModalContent>
          {(onClose) => (
            <>
              <ModalHeader className='flex flex-col gap-1'>{LTIForm.ButtonBaseURIMoodle}</ModalHeader>
              <ModalBody>
                <p>{LTIForm.ButtonBaseURIMoodleDescription}</p>
                <Input
                  variant='bordered'
                  label={LTIForm.FieldBaseURI}
                  description={LTIForm.DescriptionBaseURI}
                  type='url'
                  value={baseUrl}
                  onChange={(event) => {
                    setBaseUrl(event.target.value);
                  }}
                />
              </ModalBody>
              <ModalFooter>
                <Button color='danger' variant='light' onPress={onClose}>
                  {Sidebar.Close}
                </Button>
                <Button
                  color='primary'
                  onPress={() => {
                    onClose();
                    onChange(baseUrl);
                  }}
                >
                  {Sidebar.Save}
                </Button>
              </ModalFooter>
            </>
          )}
        </ModalContent>
      </Modal>
    );
  };
  return { component, onOpen };
};
