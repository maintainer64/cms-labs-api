/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  type AuthSSOAuthorizeRequest,
  type AuthSSOAuthorizeResponse,
  CoreJsonRpcPath,
  transportWithAuth
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<AuthSSOAuthorizeRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<AuthSSOAuthorizeResponse['result']>;

export const useQuerySsoAuthorize = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'sso.authorize', params, { retry: 1 })
  );
};
