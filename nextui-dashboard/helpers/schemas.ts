import {object, string} from "yup";
import useLanguageBrowser from "@/helpers/locale";

export const LoginSchema = () => {
    const {locale} = useLanguageBrowser();
    object().shape({
        email: string()
            .email(locale.Login.ErrorFieldEmailNotEmpty)
            .required(locale.Login.ErrorFieldEmailRequired),
        password: string().required(locale.Login.ErrorFieldPasswordRequired),
    });
}