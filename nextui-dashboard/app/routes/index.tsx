import { useUserProfile } from '@/components/providers/auth-jwt/hooks';
import * as React from 'react';
import RoutesUnknown from '@/app/routes/auth';
import RoutesStudent from '@/app/routes/student';
import RoutesAdmin from '@/app/routes/admin';
import { UserRoleBase } from '@/helpers/queries/sso/auth';

const RoutesDynamic = () => {
  const user = useUserProfile();
  if (!user || !user.sub) return <RoutesUnknown />;
  if (user?.roles?.[0] === UserRoleBase.Admin || user?.roles?.[0] == UserRoleBase.Instructor) return <RoutesAdmin />;
  return <RoutesStudent />;
};
export default RoutesDynamic;
