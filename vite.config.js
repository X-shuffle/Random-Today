import { defineConfig } from 'vite';
import { resolve } from 'path';

export default defineConfig({
  // 配置环境变量前缀
  envPrefix: 'VITE_',

  // 构建配置
  build: {
    outDir: 'dist',
    assetsDir: 'assets',
    sourcemap: false,
    minify: 'terser',
    rollupOptions: {
      input: {
        main: resolve(__dirname, 'index.html'),
        admin: resolve(__dirname, 'admin.html')
      }
    }
  },

  // 开发服务器配置
  server: {
    port: 3000,
    open: true
  }
});
