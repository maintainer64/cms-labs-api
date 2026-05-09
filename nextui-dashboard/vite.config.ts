import {CommonServerOptions, defineConfig} from 'vite'
import react from '@vitejs/plugin-react'
import {resolve} from 'path';

/** @type {import('vite').UserConfig} */
export default defineConfig({
    plugins: [react()],
    resolve: {
        alias: {
            '@': resolve(__dirname),
            '@srl-labs/clab-ui': resolve(__dirname, '/node_modules/@srl-labs/clab-ui')
        },
    },
    define: {
        global: 'window'
    },
    server: {
        port: 5183,
        cors: false,
        proxy: setupProxy(),
    },
})

function setupProxy(): CommonServerOptions["proxy"] {
    return {
        "/api": {
            target: "https://auth.k8s-pre.cmslabs.ru",
            secure: false,
            changeOrigin: true,
        },
        "/clabgate/api/": {
            target: "https://auth.k8s-pre.cmslabs.ru",
            secure: false,
            changeOrigin: true,
        }
    };
}
