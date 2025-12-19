import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'


// https://vite.dev/config/
export default defineConfig({
  plugins: [
    react()
  ],
  server: {
    host: '0.0.0.0',
    port: 5173,
    proxy: {
      '/api': {
        // When running in Docker Compose, use service name 'backend'
        // Services on the same Docker network can communicate using service names
        target: 'http://backend:3000',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, ''),
      },
    },
    // Enable polling for file watching (required for OneDrive folders)
    watch: {
      usePolling: true,
      interval: 1000,
    },
  },
})
