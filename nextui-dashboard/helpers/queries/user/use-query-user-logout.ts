/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type AuthUserLogoutRequest,
  type AuthSwaggerSSOTokenResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<AuthUserLogoutRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<AuthSwaggerSSOTokenResponse['result']>;

export const useQueryUserLogout = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'user.logout', params, { retry: 3 })
  );
};
