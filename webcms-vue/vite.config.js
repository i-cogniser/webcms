import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import path from 'path';

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
                target: 'http://web:8080',
                changeOrigin: true,
                secure: false,
                rewrite: path => path.replace(/^\/api/, ''),
                logLevel: 'debug'
            },
        },
    },
});
