/* eslint-disable */
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

export interface QueriesTaskCodeRegistryItem {
  full_path?: string;
  id?: string;
  namespace_suffix?: string;
  title?: string;
}

export interface TopologyTopologiesEdge {
  data?: TopologyTopologiesEdgeData;
  id?: string;
  source?: string;
  target?: string;
  type?: string;
}

export interface TopologyTopologiesEdgeData {
  source?: TopologyTopologiesEdgeDataItem;
  target?: TopologyTopologiesEdgeDataItem;
}

export interface TopologyTopologiesEdgeDataItem {
  id?: string;
  label?: string;
  name?: string;
  type?: string;
}

export interface TopologyTopologiesNode {
  data?: TopologyTopologiesNodeDataItem;
  /** Icon enum:cloud,router,server,switch,desktop */
  icon?: string;
  id?: string;
  label?: string;
  type?: string;
}

export interface TopologyTopologiesNodeDataItem {
  /** Icon enum:cloud,router,server,switch,desktop */
  icon?: string;
  id?: string;
  kind?: string;
  label?: string;
  startup?: string;
  type?: string;
}

export interface TopologyTopology {
  direction?: string;
  edges?: TopologyTopologiesEdge[];
  nodes?: TopologyTopologiesNode[];
}

export interface UsecasesContainerGetInputDTO {
  deployment?: string;
  namespace?: string;
}

export interface UsecasesContainerGetItem {
  connect_url?: string;
  label?: string;
  name?: string;
  namespace?: string;
  pod?: string;
  ready?: boolean;
  restart_count?: number;
  session_id?: string;
  startup?: string;
  /** Status enum: running,waiting,terminated,unknown */
  status?: string;
  /** Type enum: containerlab,default */
  type?: string;
}

export interface UsecasesContainerGetOutputDTO {
  containers?: UsecasesContainerGetItem[];
}

export interface UsecasesContainerGetRequest {
  /** @default "1" */
  id?: string;
  /** @default "2.0" */
  jsonrpc?: string;
  /** @default "container.get" */
  method?: string;
  params?: UsecasesContainerGetInputDTO;
}

export interface UsecasesContainersGetResponse {
  error?: any;
  /** @default "1" */
  id?: string;
  /** @default "2.0" */
  jsonrpc?: string;
  result?: UsecasesContainerGetOutputDTO;
}

export type UsecasesTaskListInputDTO = object;

export interface UsecasesTaskListOutputDTO {
  model: QueriesTaskCodeRegistryItem[];
}

export interface UsecasesTaskListRequest {
  /** @default "1" */
  id?: string;
  /** @default "2.0" */
  jsonrpc?: string;
  /** @default "tasks.list" */
  method?: string;
  params?: UsecasesTaskListInputDTO;
}

export interface UsecasesTaskListResponse {
  error?: any;
  /** @default "1" */
  id?: string;
  /** @default "2.0" */
  jsonrpc?: string;
  result?: UsecasesTaskListOutputDTO;
}

export interface UsecasesTopologiesGetInputDTO {
  namespace?: string;
}

export interface UsecasesTopologiesGetOutputDTO {
  topology?: TopologyTopology;
  web_url?: string;
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

export interface UsecasesTopologyCreateInputDTO {
  redeploy?: boolean;
  task_id?: string;
  username?: string;
}

export interface UsecasesTopologyCreateOutputDTO {
  deploy_created?: boolean;
  namespace?: string;
  namespace_created?: boolean;
  user_created?: boolean;
}

export interface UsecasesTopologyCreateRequest {
  /** @default "1" */
  id?: string;
  /** @default "2.0" */
  jsonrpc?: string;
  /** @default "topology.create" */
  method?: string;
  params?: UsecasesTopologyCreateInputDTO;
}

export interface UsecasesTopologyCreateResponse {
  error?: any;
  /** @default "1" */
  id?: string;
  /** @default "2.0" */
  jsonrpc?: string;
  result?: UsecasesTopologyCreateOutputDTO;
}

export interface UsecasesTopologyDeleteInputDTO {
  namespaces?: string[];
}

export interface UsecasesTopologyDeleteOutputDTO {
  namespaces?: string[];
}

export interface UsecasesTopologyDeleteRequest {
  /** @default "1" */
  id?: string;
  /** @default "2.0" */
  jsonrpc?: string;
  /** @default "topology.delete" */
  method?: string;
  params?: UsecasesTopologyDeleteInputDTO;
}

export interface UsecasesTopologyDeleteResponse {
  error?: any;
  /** @default "1" */
  id?: string;
  /** @default "2.0" */
  jsonrpc?: string;
  result?: UsecasesTopologyDeleteOutputDTO;
}
