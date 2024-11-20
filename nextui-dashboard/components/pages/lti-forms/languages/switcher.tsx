"use client";
import {LanguageType} from "@/helpers/locale/locale";
import {Image} from '@nextui-org/react'
import FlagEn from './flagEn.svg'
import FlagRu from './flagRu.svg'
import useLanguageBrowser from "@/helpers/locale";
import {useNavigate} from "react-router-dom";

const languages = [
    {
        text: 'Русский',
        img: FlagRu,
        lang: 'ru' as LanguageType,
    },
    {
        text: 'English',
        img: FlagEn,
        lang: 'en' as LanguageType,
    }
]

const LanguageSwitcher = () => {
    const navigate = useNavigate();
    const {locale, setLang} = useLanguageBrowser();
    const handleClose = () => {
        navigate(-1);
    };

    const langBlock = languages.map((item, index) => (
        <div
            key={`lang-${index}`}
            onClick={() => {
                setLang(item.lang)
            }}
            className="flex flex-col items-center p-4 border rounded-lg shadow bg-gray-100 hover:bg-gray-300 cursor-pointer">
            <Image
                src={item.img}
                alt={item.lang}
                width="20"
                height="20"
                className="w-16 h-16 mb-2"
            />
            <span>{item.text}</span>
        </div>
    ))


    return (
        <div className="flex items-center justify-center h-screen bg-gray-100">
            <div className="relative p-6 bg-gray-100 rounded-lg shadow-lg max-w-md mx-auto">
                <button
                    onClick={handleClose}
                    className="absolute top-2 right-2 text-gray-500 hover:text-gray-700"
                >
                    &times;
                </button>
                <h2 className="text-xl font-semibold mb-4 text-center">{locale.LanguageSwitcher.LanguageSwitch}</h2>
                <div className="grid grid-cols-2 gap-4">
                    {langBlock}
                </div>
            </div>
        </div>
    );
};

export default LanguageSwitcher;
