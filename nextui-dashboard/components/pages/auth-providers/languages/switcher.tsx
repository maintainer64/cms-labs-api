'use client';
import { LanguageType } from '@/helpers/locale/locale';
import { Image } from '@heroui/react';
import FlagEn from './flagEn.svg';
import FlagRu from './flagRu.svg';
import useLanguageBrowser from '@/helpers/locale';
import { useNavigate } from 'react-router-dom';
import { Modal, ModalBody, ModalContent, ModalHeader } from '@heroui/modal';
import React from 'react';

const languages = [
  {
    text: 'Русский',
    img: FlagRu,
    lang: 'ru' as LanguageType
  },
  {
    text: 'English',
    img: FlagEn,
    lang: 'en' as LanguageType
  }
];

const LanguageSwitcher = () => {
  const navigate = useNavigate();
  const { locale, setLang } = useLanguageBrowser();
  const handleClose = () => {
    navigate(-1);
  };

  const langBlock = languages.map((item, index) => (
    <button
      key={`lang-${index}`}
      onClick={() => {
        setLang(item.lang);
      }}
      className='flex flex-col items-center p-4 border rounded-lg shadow bg-default-50 hover:bg-default-100 cursor-pointer'
    >
      <Image src={item.img} alt={item.lang} width='20' height='20' className='w-16 h-16 mb-2' />
      <span>{item.text}</span>
    </button>
  ));

  return (
    <Modal isOpen={true} onClose={handleClose}>
      <ModalContent>
        <ModalHeader className='flex flex-col gap-1'>{locale.LanguageSwitcher.LanguageSwitch}</ModalHeader>
        <ModalBody>{langBlock}</ModalBody>
      </ModalContent>
    </Modal>
  );
};

export default LanguageSwitcher;
