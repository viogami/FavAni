import axios from '../utils/axios.js'

// get条目的函数
function SearchSubject (keywords, type, responseGroup, start, max_results) {
  return axios.bangumiService.get('/search/subject/' + keywords, {
    params: {
      type,
      responseGroup,
      start,
      max_results
    }
  })
}

export { SearchSubject }
