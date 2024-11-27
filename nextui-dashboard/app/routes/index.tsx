import { useUserProfile } from '@/components/providers/auth-jwt/hooks';
import * as React from 'react';
import RoutesUnknown from '@/app/routes/auth';
import RoutesStudent from '@/app/routes/student';
import RoutesAdmin from '@/app/routes/admin';

const RoutesDynamic = () => {
  const user = useUserProfile();
  if (!user || !user.id) return <RoutesUnknown />;
  if (user.role === 'admin' || user.role == 'instructor') return <RoutesAdmin />;
  return <RoutesStudent />;
};
export default RoutesDynamic;
