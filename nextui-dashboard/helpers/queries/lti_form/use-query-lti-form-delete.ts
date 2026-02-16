/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type UsecasesAuthProviderDeleteRequest,
  type UsecasesAuthProviderDeleteResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesAuthProviderDeleteRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesAuthProviderDeleteResponse['result']>;

export const useQueryAuthProviderDelete = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'lti_form.delete', params, { retry: 3 })
  );
};
