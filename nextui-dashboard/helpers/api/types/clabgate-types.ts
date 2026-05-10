/* tslint:disable */
// @ts-nocheck
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

export interface QueriesDeploymentInfo {
  name?: string;
  ready_replicas?: number;
  replicas?: number;
  restarts?: number;
  status?: string;
}

export interface QueriesServiceInfo {
  cluster_ip?: string;
  external_ip?: string[];
  name?: string;
  type?: string;
}

export interface QueriesTTYDInfo {
  name?: string;
  url?: string;
}

export interface UsecasesTopologiesGetInputDTO {
  attempt_number?: string;
  username?: string;
}

export interface UsecasesTopologiesGetOutputDTO {
  deployments?: QueriesDeploymentInfo[];
  services?: QueriesServiceInfo[];
  topology?: string;
  ttyd?: QueriesTTYDInfo[];
}

export interface UsecasesTopologiesGetRequest {
  /** @default "1" */
  id?: string;
  /** @default "2.0" */
  jsonrpc?: string;
  /** @default "topology.get" */
  method?: string;
  params?: UsecasesTopologiesGetInputDTO;
}

export interface UsecasesTopologiesGetResponse {
  error?: any;
  /** @default "1" */
  id?: string;
  /** @default "2.0" */
  jsonrpc?: string;
  result?: UsecasesTopologiesGetOutputDTO;
}
