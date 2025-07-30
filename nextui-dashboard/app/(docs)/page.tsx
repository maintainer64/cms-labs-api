import useLanguageBrowser from '@/helpers/locale';
import { Modal, ModalBody, ModalContent, ModalHeader } from '@heroui/modal';
import React from 'react';
import MarkdownViewer from '@/components/markdown';
import { useLocation, useNavigate } from 'react-router-dom';

export function DocsKubectlTopology() {
  const { lang } = useLanguageBrowser();
  const navigate = useNavigate();
  const { search } = useLocation();
  const query = new URLSearchParams(search);
  const {
    locale: { Topology }
  } = useLanguageBrowser();
  return (
    <Modal
      isOpen={true}
      onClose={() => {
        navigate(-1);
      }}
      size='full'
      scrollBehavior='inside'
    >
      <ModalContent>
        <ModalHeader className='flex flex-col gap-1'>{Topology.Menu.ConnectToKubectlTitle}</ModalHeader>
        <ModalBody>
          <MarkdownViewer
            path={`kubectl/${lang}.md`}
            variables={{
              namespace: query.get('namespace') || '<namespace>'
            }}
          />
        </ModalBody>
      </ModalContent>
    </Modal>
  );
}
