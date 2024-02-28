/** @type {import('tailwindcss').Config} */
export default {
  content: [
    './index.html',
    './src/**/*.{html,vue,js,css}',
    './src/**/**/*.{html,vue,js,css}'],
  theme: {
    extend: {}
  },
  plugins: [
    '@tailwindcss/typography' // tailwind插件位置
  ]
}
