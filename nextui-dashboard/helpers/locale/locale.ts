import ru from './locales/ru';
import en from './locales/en';
import {useLocalStorage} from '@uidotdev/usehooks';

export type LanguageType = 'ru' | 'en';

const languageResource = (lang: LanguageType) => {
    return lang === 'ru' ? ru : en;
}

const useLanguageBrowser = () => {
    try {
        let [lang, setLang] = useLocalStorage("lang", "en" as LanguageType);
        return {
            locale: languageResource(lang),
            setLang: (lang: LanguageType) => setLang(lang)
        }
    } catch (error) {
        return {
            locale: languageResource("en"),
            setLang: (lang: LanguageType) => {
            }
        }
    }
}
export default useLanguageBrowser;