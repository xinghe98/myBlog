import { fileURLToPath, URL } from "node:url";
import viteCompression from "vite-plugin-compression";
import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import AutoImport from "unplugin-auto-import/vite";
import Components from "unplugin-vue-components/vite";
import { ElementPlusResolver } from "unplugin-vue-components/resolvers";
// import compress from 'vite-plugin-compress'
// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    viteCompression(),
    AutoImport({
      resolvers: [ElementPlusResolver()],
    }),
    Components({
      resolvers: [ElementPlusResolver()],
    }),
  ],
  resolve: {
    alias: {
      "@": fileURLToPath(new URL("./src", import.meta.url)),
    },
  },
  server: {
    proxy: {
      // 字符串简写写法
      // 选项写法
      "/api": {
        target: "http://124.71.70.42:3000/api",
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, ""),
      },
    },
  },
  build: {
    assetsDir: "./static",
    chunkSizeWarningLimit: 500,
    minify: "terser",
    cssCodeSplit: true, // 如果设置为false，整个项目中的所有 CSS 将被提取到一个 CSS 文件中
    terserOptions: {
      compress: {
        // warnings: false,
        drop_console: true, //打包时删除console
        drop_debugger: true, //打包时删除 debugger
        pure_funcs: ["console.log"],
      },
      output: {
        // 去掉注释内容
        comments: true,
      },
    },
    rollupOptions: {
      output: {
        manualChunks: {
          // 拆分代码，这个就是分包，配置完后自动按需加载，现在还比不上webpack的splitchunk，不过也能用了。
          vue: ["vue", "vue-router"],
          echarts: ["echarts"],
          // echarts: ['echarts'],
        },
      },
    },
    brotliSize: false,
  },
});
