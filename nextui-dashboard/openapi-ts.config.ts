import {defineConfig} from '@hey-api/openapi-ts';

export default defineConfig({
    input: '../backend/docs/swagger.json',
    output: {
        path: 'helpers/api/src',
        lint: false,
        format: false,
    },
    plugins: ['legacy/axios'],
});
