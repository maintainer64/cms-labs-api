import { useUserProfile } from '@/components/providers/auth-jwt/hooks';
import { useNavigate } from 'react-router-dom';
import useLanguageBrowser from '@/helpers/locale';
import { UserRoleBase } from '@/helpers/queries/sso/auth';
import { Button } from '@heroui/react';

interface Props {
  allowedRoles?: UserRoleBase[];
  children?: any;
}

/**
 * Компонент для контроля доступа на основе ролей
 * @param {Object} props - Пропсы компонента
 * @param {Array<string>} props.allowedRoles - Разрешенные роли
 * @param {React.ReactNode} props.children - Дочерние элементы
 */
export const RoleBasedAccess = ({ allowedRoles, children }: Props) => {
  const navigate = useNavigate();
  const handleClose = () => {
    navigate(-1);
  };
  const {
    locale: { RoleBasedAccess }
  } = useLanguageBrowser();

  const user = useUserProfile();
  // Проверяем, есть ли у пользователя хотя бы одна из разрешенных ролей
  const hasAccess = allowedRoles?.some((role) => user.roles?.includes(role));

  if (hasAccess) return children;

  // Контент, если нет доступа
  return (
    <main className='grid min-h-full place-items-center bg-white px-6 py-24 sm:py-32 lg:px-8'>
      <div className='text-center'>
        <p className='text-base font-semibold text-indigo-600'>403</p>
        <h1 className='mt-4 text-5xl font-semibold tracking-tight text-balance text-gray-900 sm:text-7xl'>
          {RoleBasedAccess.Title}
        </h1>
        <p className='mt-6 text-lg font-medium text-pretty text-gray-500 sm:text-xl/8'>{RoleBasedAccess.Description}</p>
        <div className='mt-10 flex items-center justify-center gap-x-6'>
          <Button
            onPress={handleClose}
            className='rounded-md bg-indigo-600 px-3.5 py-2.5 text-sm font-semibold text-white shadow-xs
                        hover:bg-indigo-500 focus-visible:outline-2 focus-visible:outline-offset-2
                        focus-visible:outline-indigo-600'
          >
            {RoleBasedAccess.Button}
          </Button>
        </div>
      </div>
    </main>
  );
};
