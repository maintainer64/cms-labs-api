import { useUserProfile } from '@/components/providers/auth-jwt/hooks';
import * as React from 'react';
import RoutesUnknown from '@/app/routes/auth';
import RoutesStudent from '@/app/routes/student';
import RoutesAdmin from '@/app/routes/admin';

const RoutesDynamic = () => {
  const user = useUserProfile();
  if (!user || !user.sub) return <RoutesUnknown />;
  if (user?.roles?.[0] === 'admin' || user?.roles?.[0] == 'instructor') return <RoutesAdmin />;
  return <RoutesStudent />;
};
export default RoutesDynamic;
