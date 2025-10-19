import { defineConfig } from 'vite';

export default defineConfig({
  // 配置环境变量前缀
  envPrefix: 'VITE_',
  
  // 构建配置
  build: {
    outDir: 'dist',
    assetsDir: 'assets',
    sourcemap: false,
    minify: 'terser'
  },
  
  // 开发服务器配置
  server: {
    port: 3000,
    open: true
  }
});
