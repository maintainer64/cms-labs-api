import ru from './locales/ru';
import en from './locales/en';
import { useGlobalStoreGet, useGlobalStoreSet } from '@/helpers/queries/users/store';

export type LanguageType = 'ru' | 'en';

const languageResource = (lang: LanguageType) => {
  return lang === 'en' ? en : ru;
};

const useLanguageBrowser = () => {
  const globalStoreQuery = useGlobalStoreGet();
  const { mutate } = useGlobalStoreSet();
  // @ts-ignore
  const lang = (globalStoreQuery?.data?.['lang'] || 'ru') as LanguageType;
  return {
    locale: languageResource(lang),
    lang: lang as LanguageType,
    setLang: (lang: LanguageType) => {
      mutate({ ...(globalStoreQuery?.data || {}), lang: lang });
    }
  };
};
export default useLanguageBrowser;
