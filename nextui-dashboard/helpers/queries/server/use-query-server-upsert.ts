/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type UsecasesPNETServerEditRequest,
  type UsecasesPNETServerEditResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesPNETServerEditRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesPNETServerEditResponse['result']>;

export const useQueryServerUpsert = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'server.upsert', params, { retry: 3 })
  );
};
