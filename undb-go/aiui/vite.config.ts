// vite.config.ts 或 vite.config.js
import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';

export default defineConfig({
  plugins: [react()],
  server: {
    proxy: {
      // 只要以 /api 开头的请求，都转发到后端
      '/api': {
        target: 'http://localhost:5555',
        changeOrigin: true,
        // 如果你的后端没有 /api 前缀，可以加上 rewrite
        // rewrite: (path) => path.replace(/^\/api/, ''),
      },
    },
  },
});