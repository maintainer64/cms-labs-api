import ru from './locales/ru';
import en from './locales/en';
import { useQueryUserGlobalStoreGet } from '@/helpers/queries/user/use-query-user-global-store-get';
import { useMutationUserGlobalStoreSet } from '@/helpers/queries/user/use-mutation-user-global-store-set';

export type LanguageType = 'ru' | 'en';

const languageResource = (lang: LanguageType) => {
  return lang === 'en' ? en : ru;
};

const useLanguageBrowser = () => {
  const globalStoreQuery = useQueryUserGlobalStoreGet({});
  const { mutate } = useMutationUserGlobalStoreSet();
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
