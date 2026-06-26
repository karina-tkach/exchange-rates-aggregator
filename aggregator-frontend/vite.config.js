import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import tailwindcss from '@tailwindcss/vite'

export default defineConfig({
  plugins: [react(),
    tailwindcss()],

    server: {
    port: 5173,
    strictPort: true,
    proxy: {
      "/api/v1": "http://api:8000"
    }
  }
})
