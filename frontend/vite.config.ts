import { defineConfig } from 'vite'
import { resolve } from 'path';
import vue from '@vitejs/plugin-vue'
import Unocss from 'unocss/vite'

 

// https://vitejs.dev/config/
export default defineConfig({
  envDir: 'envs',
  plugins: [
    vue(),
    Unocss({ /* options */ }),
  ],
  server:{
    host:'0.0.0.0', 
  },
  resolve: {
    alias: [{
      find: '@',
      replacement: resolve(__dirname, '/src'),
    }],

  },
  build: {
    rollupOptions: {
      output: {
        assetFileNames: 'assets/[hash][extname]',
        chunkFileNames: 'assets/[hash].js',
        manualChunks(id) {
          if (id.includes('node_modules')) {
            return id.split('node_modules/')[1].split('/')[0]
          }
        }
      }
    }
  }
})