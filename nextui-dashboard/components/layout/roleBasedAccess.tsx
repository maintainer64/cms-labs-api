import { useUserProfile } from '@/components/providers/auth-jwt/hooks';
import { useNavigate } from 'react-router-dom';
import useLanguageBrowser from '@/helpers/locale';
import { UserRoleBase } from '@/helpers/queries/sso/auth';
import { InfoModalBlock } from '@/components/layout/infoModalBlock';

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
    <InfoModalBlock
      upTitle='403'
      title={RoleBasedAccess.Title}
      description={RoleBasedAccess.Description}
      onClick={handleClose}
      buttonText={RoleBasedAccess.Button}
    />
  );
};
