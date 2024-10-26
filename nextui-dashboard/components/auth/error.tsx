import useLanguageBrowser from "@/helpers/locale";

export default function AuthError() {
    const {locale} = useLanguageBrowser();
    return (
        <div className="flex items-center justify-center h-screen">
            <div className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded relative max-w-md mx-auto">
                <strong className="font-bold">{locale.Auth.ErrorPageTitle}</strong><br/>
                <span className="block sm:inline">{locale.Auth.ErrorPageDescription}</span>
            </div>
        </div>
    );
}