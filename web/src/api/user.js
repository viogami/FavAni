import axios from '../utils/axios.js'

// bangumi登陆和默认登陆
function BangumiLogin (username) {
  return axios.bangumiService.get('/v0/users/' + username, {
    username
  })
}
function DefaultLogin (username,password) {
  return axios.favaniService.post('/login',{"username":username,"password":password})
}
// 用户注销,headers为请求头，使用jwt认证
function Logout(headers){
  return axios.favaniService.post('/auth/logout',null,{ headers } )
}
// 用户注册
function UserRegister(username,email,password){
  return axios.favaniService.post('/register',{"username":username,"email":email,"password":password})
}
// get用户头像
function getUserAvatar (username, size) {
  return axios.bangumiService.get('/v0/users/' + username + '/avatar', {
    params: {
      username,
      size
    }
  })
}

// 获得用户收藏信息
function userFavorite (username, subject_type, type, limit, offset) {
  return axios.bangumiService.get('/v0/users/' + username + '/collections', {
    params: {
      subject_type,
      type,
      limit,
      offset
    }
  })
}

// 删除用户
function DelUser (username) {
  return axios.favaniService.post('/deluser',{"Username":username})
}

export { BangumiLogin, DefaultLogin,Logout,UserRegister,userFavorite,DelUser}