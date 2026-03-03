const services = [
    {
        fullPathOpenApi: '../backend/docs/swagger.json',
        baseApiRpc: [
            '/core/api/v1/rpc'
        ],
        baseApiRpcConst: "CoreJsonRpcPath",
        generateTypesFile: 'core',
    },
    {
        fullPathOpenApi: '../clabgate/docs/swagger.json',
        baseApiRpc: [
            '/clabgate/api/v1/rpc'
        ],
        baseApiRpcConst: "ClabgateJsonRpcPath",
        generateTypesFile: 'clabgate',
    }
]
export default services;
