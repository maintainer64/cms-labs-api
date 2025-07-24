import { useLocation, useParams } from 'react-router-dom';
import { useUserProfile } from '@/components/providers/auth-jwt/hooks';

export function useParamsConnectTopology() {
  const { namespace } = useParams();
  const { search } = useLocation();
  const user = useUserProfile();

  const query = new URLSearchParams(search);
  return {
    username: query.get('username') || user.username || '',
    taskId: query.get('taskId') || user.username || '',
    redeploy: !!query.get('redeploy'),
    namespace: namespace || ''
  };
}
