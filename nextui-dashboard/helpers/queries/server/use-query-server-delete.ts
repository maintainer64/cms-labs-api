/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type UsecasesPNETServerDeleteRequest,
  type UsecasesPNETServerDeleteResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesPNETServerDeleteRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesPNETServerDeleteResponse['result']>;

export const useQueryServerDelete = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'server.delete', params, { retry: 3 })
  );
};
