/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type AuthRenewManagerCredentialsRequest,
  type AuthSwaggerSSOTokenResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<AuthRenewManagerCredentialsRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<AuthSwaggerSSOTokenResponse['result']>;

export const useQueryUserLogin = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'user.login', params, { retry: 3 })
  );
};
