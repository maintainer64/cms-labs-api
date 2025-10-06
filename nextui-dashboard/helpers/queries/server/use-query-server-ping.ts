/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type ExternalPNETServerPingRequest,
  type ExternalPNETServerPingResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<ExternalPNETServerPingRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<ExternalPNETServerPingResponse['result']>;

export const useQueryServerPing = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'server.ping', params, { retry: 3 })
  );
};
