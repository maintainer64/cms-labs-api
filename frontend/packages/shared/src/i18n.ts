import i18n from 'i18next';
import { initReactI18next } from 'react-i18next';
import ru from './translation/ru.json';
import en from './translation/en.json';

i18n
  .use(initReactI18next)
  .init({
    fallbackLng: 'ru',
    lng: 'en',
    interpolation: {
      escapeValue: false
    },
    resources: {
      ru: {
        translation: ru
      },
      en: {
        translation: en
      }
    }
  });

export default i18n;