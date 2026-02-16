/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type UsecasesAuthProviderEditRequest,
  type UsecasesAuthProviderEditResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesAuthProviderEditRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesAuthProviderEditResponse['result']>;

export const useQueryAuthProviderUpsert = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'lti_form.upsert', params, { retry: 3 })
  );
};
