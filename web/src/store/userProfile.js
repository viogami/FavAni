import { defineStore } from 'pinia'

const useUserStore = defineStore('userProfile', {
  state: () => ({
    username: '',
    nickname: '',
    avatarUrl: '',
    sign: '',
    favorList: [],
    favorList_max: 0,
    bangumiLogin: false,
  }),
  actions: {
    // 定义一个动作用于设置用户信息
    setUserInfo (username, nickname, avatarUrl, userSign, bangumiLogin) {
      this.username = username
      this.nickname = nickname
      this.avatarUrl = avatarUrl
      this.sign = userSign
      this.bangumiLogin = bangumiLogin
    },
    // 添加收藏到数组中
    addFavorList (id,name,name_cn) {
      const newItem = {
        id: id,
        name: name,
        name_cn: name_cn,
      };
      if (this.favorList.length <= this.favorList_max) {
        this.favorList.push(newItem)
      }
    },
    // 获取收藏数组
    getFavorList (page, page_items) {
      if (page === 'all') { return this.favorList } else { return this.favorList.slice(page_items * page, page_items * page + page_items) }
    },
    
    // 删除收藏
    delFavorList (id) {
      for (let i = 0; i < this.favorList.length; i++) {
        if (this.favorList[i].id === id) {
          this.favorList.splice(i, 1)
          break
        }
      }
    },

    // 修改信息
    updateUsername (newUsername) {
      this.username = newUsername
    },
    updateNickname (newNickname) {
      this.nickname = newNickname
    },
    updateAvatarUrl (newAvatarUrl) {
      this.avatarUrl = newAvatarUrl
    },
    updateSign (newUserSign) {
      this.sign = newUserSign
    }
  }
})

export { useUserStore }
