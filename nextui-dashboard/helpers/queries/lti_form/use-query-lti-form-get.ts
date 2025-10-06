/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type UsecasesLTIFormGetRequest,
  type UsecasesLTIFormGetResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesLTIFormGetRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesLTIFormGetResponse['result']>;

export const useQueryLtiFormGet = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'lti_form.get', params, { retry: 3 })
  );
};
