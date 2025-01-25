import useLanguageBrowser from '@/helpers/locale';

type ErrorModalProps = {
  title: string;
  description: string;
  children?: React.ReactNode
}
export const ErrorModal = ({ title, description, children }: ErrorModalProps) => {
  return (
    <div className="flex items-center justify-center h-screen">
      <div className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded relative max-w-md mx-auto">
        <strong className="font-bold">{title}</strong>
        <br />
        <span className="block sm:inline">{description}</span>
        {children && (
          <div>
            {children}
          </div>
        )}
      </div>
    </div>
  );
};
export default function AuthError() {
  const { locale } = useLanguageBrowser();
  return <ErrorModal
    title={locale.Auth.ErrorPageTitle}
    description={locale.Auth.ErrorPageDescription}
  />;
}
