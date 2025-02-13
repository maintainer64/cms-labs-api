import ru from './locales/ru';
import en from './locales/en';
import { useLocalStorage } from '@uidotdev/usehooks';

export type LanguageType = 'ru' | 'en';

const languageResource = (lang: LanguageType) => {
  return lang === 'en' ? en : ru;
};

const useLanguageBrowser = () => {
  try {
    const [lang, setLang] = useLocalStorage('lang', 'ru' as LanguageType);
    return {
      locale: languageResource(lang),
      setLang: (lang: LanguageType) => setLang(lang)
    };
  } catch (error) {
    return {
      locale: languageResource('ru'),
      setLang: (lang: LanguageType) => {
        console.debug(lang);
      }
    };
  }
};
export default useLanguageBrowser;
