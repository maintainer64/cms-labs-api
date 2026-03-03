/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type UsecasesLTIAttemptDeleteRequest,
  type UsecasesLTIAttemptDeleteResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesLTIAttemptDeleteRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesLTIAttemptDeleteResponse['result']>;

export const useQueryLtiAttemptDelete = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'lti_attempt.delete', params, { retry: 3 })
  );
};
