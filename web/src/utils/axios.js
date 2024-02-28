import axios from 'axios'

const bangumiService = axios.create({
  baseURL: '/bgmapi'
})
const favaniService = axios.create({
  baseURL: '/api'
})
export default { bangumiService, favaniService }
