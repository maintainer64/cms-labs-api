import TokenManager from "@/helpers/api/axios-refresh";
import {OpenAPI} from "@/helpers/api/src";
OpenAPI.TOKEN = TokenManager.getToken.bind(TokenManager)
