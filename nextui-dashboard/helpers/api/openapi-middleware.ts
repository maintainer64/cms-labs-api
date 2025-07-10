import TokenManager from '@/helpers/api/axios-refresh';
import {OpenAPI as OpenAPIBackend} from './backend';
import {OpenAPI as OpenAPIClabgate} from './clabgate';

OpenAPIBackend.TOKEN = TokenManager.getToken.bind(TokenManager);
OpenAPIClabgate.TOKEN = TokenManager.getToken.bind(TokenManager);
