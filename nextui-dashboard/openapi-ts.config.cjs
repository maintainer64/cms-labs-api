/** @type {import('@hey-api/openapi-ts').UserConfig} */
module.exports = {
    client: 'legacy/axios',
    input: 'http://localhost:5000/docs/doc.json',
    output: 'helpers/api/src',
    types: {
        enums: 'typescript',
    },
};
