/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type AuthRenewManagerRefreshRequest,
  type AuthSwaggerSSOTokenResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<AuthRenewManagerRefreshRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<AuthSwaggerSSOTokenResponse['result']>;

export const useQueryUserTokenRefresh = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'user.token_refresh', params, { retry: 3 })
  );
};
