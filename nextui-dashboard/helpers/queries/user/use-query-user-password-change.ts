/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type AuthUserPasswordChangeRequest,
  type AuthUserPasswordChangeResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<AuthUserPasswordChangeRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<AuthUserPasswordChangeResponse['result']>;

export const useQueryUserPasswordChange = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'user.password_change', params, { retry: 3 })
  );
};
