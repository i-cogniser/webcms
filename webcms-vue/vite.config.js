import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import path from 'path';
import dotenv from 'dotenv';

dotenv.config();

const proxyTarget = process.env.PROXY_TARGET || 'http://localhost:8080';

export default defineConfig({
    plugins: [vue()],
    build: {
        outDir: 'dist',
    },
    resolve: {
        alias: {
            '@': path.resolve(__dirname, 'src'),
        },
    },
    server: {
        proxy: {
            '/api': {
                target: proxyTarget,
                changeOrigin: true,
                secure: false,
                rewrite: (path) => path.replace(/^\/api/, ''),
                logLevel: 'debug',
                configure: (proxy, options) => {
                    proxy.on('proxyReq', (proxyReq, req, res) => {
                        console.log('Proxying request:', req.url);
                    });
                    proxy.on('proxyRes', (proxyRes, req, res) => {
                        console.log('Received response from target:', proxyRes.statusCode, req.url);
                    });
                    proxy.on('error', (err, req, res) => {
                        console.error('Proxy error:', err, req.url);
                    });
                }
            },
        },
    },
});
