/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type UsecasesLTIAttemptGetRequest,
  type UsecasesLTIAttemptGetResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesLTIAttemptGetRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesLTIAttemptGetResponse['result']>;

export const useQueryLtiAttemptGet = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'lti_attempt.get', params, { retry: 3 })
  );
};
