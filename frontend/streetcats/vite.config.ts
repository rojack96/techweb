import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

export default defineConfig({
  plugins: [react()],
  server: {
    proxy: {
      '/streetcats-service/api/v1': {
        target: 'http://localhost:8201',
        changeOrigin: true,
        secure: false,
      },
    },
  },
})