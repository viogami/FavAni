import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'
import { prismjsPlugin } from 'vite-plugin-prismjs'

const envPrefix = 'FAVANI_' // envPrefix的值默认为VITE_,改为FAVANI_
export default ({ mode }) => {
  process.env = { ...process.env, ...loadEnv(mode, process.cwd(), envPrefix) } // 使用了对象的展开语法（spread syntax）

  return defineConfig({
    plugins: [
      vue(),
      prismjsPlugin({
        languages: ['json', 'html', 'css', 'js', 'go', 'bash', 'yaml', 'c', 'c#', 'python', 'sql'],
        plugins: [
          'line-numbers', // 在代码块中显示行号
          'show-language', // 在代码块中显示代码块对应语言的名称
          'copy-to-clipboard'// 添加一个按钮，可以在单击时将代码块内容复制到剪贴板
        ],
        theme: 'tomorrow',
        css: true
      })
    ],
    envPrefix: envPrefix,
    server: {
      proxy: { // 使用 proxy 实例
        '/bgmapi': {
          target: 'https://api.bgm.tv',
          changeOrigin: true,
          rewrite: (path) => path.replace(/^\/bgmapi/, '')
        },
        '/api':{
          target: 'http://localhost:8080',
          changeOrigin: true,
          rewrite: (path) => path.replace(/^\/api/, '')
        }
      },
      port: process.env.FAVANI_PORT
    }
  })
}
