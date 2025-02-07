/** @type {import('@hey-api/openapi-ts').UserConfig} */
module.exports = {
    client: 'legacy/axios',
    input: '../backend/docs/swagger.json',
    output: 'helpers/api/src',
    types: {
        enums: 'typescript',
    },
};
