/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type UsecasesLTIFormDeleteRequest,
  type UsecasesLTIFormDeleteResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesLTIFormDeleteRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesLTIFormDeleteResponse['result']>;

export const useQueryLtiFormDelete = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'lti_form.delete', params, { retry: 3 })
  );
};
