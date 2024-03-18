import axios from '../utils/axios.js'

// 获取用户收藏
function GetFav(username){
    return axios.favaniService.get('/getfav',{"username":username} )
  }

// 添加收藏
function AddFav(username,id,data){
    return axios.favaniService.post('/addfav',{ "username":username,"anime_id":id,"data_anime":data } )
  }

// 删除收藏
function DelFav(username,data){
    return axios.favaniService.post('/delfav',{ "username":username,"data_anime":data } )
  }

  export { GetFav,AddFav,DelFav}