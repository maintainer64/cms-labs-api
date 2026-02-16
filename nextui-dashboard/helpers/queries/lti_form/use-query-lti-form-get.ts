/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type UsecasesAuthProviderGetRequest,
  type UsecasesAuthProviderGetResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesAuthProviderGetRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesAuthProviderGetResponse['result']>;

export const useQueryAuthProviderGet = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'lti_form.get', params, { retry: 3 })
  );
};
