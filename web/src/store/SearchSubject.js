import { defineStore } from 'pinia'

const useSearchEntryStore = defineStore('SearchEntryStore', {
  state: () => ({
    searchType: 'all',
    searchResList: [],
    searchResList_max: 0
  }),

  actions: {
    // 获得搜索类型
    getSearchType () {
      switch (this.searchType) {
        case '全部类型':
          return 'all'
        case '书籍':
          return 1
        case '动画':
          return 2
        case '音乐':
          return 3
        case '游戏':
          return 4
        case '三次元':
          return 6
      }
      return this.searchType
    },
    // 添加搜索结果
    addSearchRes(type,id,img,name,name_cn,rate){
      const newItem = {
        type:type,
        id: id,
        img: img,
        name: name,
        name_cn: name_cn,
        rate: rate
      };
      if (this.searchResList.length<this.searchResList_max){
        this.searchResList.push(newItem)
      }
    },
    // 获取搜索结果列表
    getSearchResList (page,page_items) {
      if (page === 'all') { return this.searchResList }
      else { return this.searchResList.slice(page_items * page, page_items * page + page_items) }
    },
    // 添加搜索结果
    test(){
      console.log('take a test')
    }
  }
})

export { useSearchEntryStore }
