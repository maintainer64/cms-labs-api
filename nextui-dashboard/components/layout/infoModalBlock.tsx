import { InfoModalButton } from '@/components/layout/infoModalButton';

interface Props {
  upTitle?: string;
  title?: string;
  description?: string;
  onClick?: () => void;
  href?: string;
  target?: string;
  buttonText?: string;
}

export const InfoModalBlock = ({ upTitle, title, description, onClick, href, target, buttonText }: Props) => {
  return (
    <main className='grid min-h-full place-items-center bg-white px-6 py-24 sm:py-32 lg:px-8 dark:bg-gray-900'>
      <div className='text-center'>
        <p className='text-base font-semibold text-indigo-600 dark:text-indigo-400'>{upTitle}</p>
        <h1 className='mt-4 text-5xl font-semibold tracking-tight text-balance text-gray-900 sm:text-7xl dark:text-white'>
          {title}
        </h1>
        <p className='mt-6 text-lg font-medium text-pretty text-gray-500 sm:text-xl/8 dark:text-gray-400'>
          {description}
        </p>
        <div className='mt-10 flex items-center justify-center gap-x-6'>
          <InfoModalButton
            onPress={onClick}
            href={href}
            className='rounded-md bg-indigo-600 px-3.5 py-2.5 text-sm font-semibold text-white shadow-xs
                        hover:bg-indigo-500 focus-visible:outline-2 focus-visible:outline-offset-2
                        focus-visible:outline-indigo-600 dark:bg-indigo-500 dark:hover:bg-indigo-400
                        dark:focus-visible:outline-indigo-500'
          >
            {buttonText}
          </InfoModalButton>
        </div>
      </div>
    </main>
  );
};
