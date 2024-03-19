<template>
  <el-card class="box-card"  >
    <!--折叠面板-->
    <el-collapse v-model="activeNames" accordion @change="handleCollapseChange">
      <div>
      <el-collapse-item title=" 条目" name="1">
     <!--条目表单-->
     <el-form
      ref="SubjectFormRef"
      :label-position="labelPosition"
      label-width="120px"
      :model="UserForm"
      style="max-width: 440px"
    >
     <el-form-item label=" " prop="subject">
        <el-input v-model="SubjectForm.keywords"
         placeholder="请输入查询条目">
         <template #prefix>
          <el-icon :size="20"> <Search /> </el-icon> <!--图标-->
         </template>
        </el-input>
      </el-form-item>
      </el-form>
      </el-collapse-item>
      </div>

      <div>
      <el-collapse-item title=" 用户" name="2">
    <!--用户表单-->
      <el-form
      ref="UserFormRef"
      :label-position="labelPosition"
      :rules="rules"
      label-width="120px"
      :model="UserForm"
      style="max-width: 440px"
    >
      <el-form-item label=" " prop="username">
        <el-input v-model="UserForm.username"
         placeholder="请输入用户名">
         <template #prefix>
          <el-icon :size="20"> <UserFilled />  </el-icon> <!--图标-->
         </template>
        </el-input>
      </el-form-item>
    </el-form>
    </el-collapse-item>
    </div>
    </el-collapse>

  <p></p>
<!--不同面板激活，按钮实行不同函数-->
<el-button type="primary" round @click="activeNames.indexOf('1')!==-1 ? SubMit_subject() : SubMit_user()">查询</el-button>
</el-card>

</template>

<script lang="ts" setup >
import { reactive, ref } from 'vue'
// 引入Vue Router
import { useRouter } from 'vue-router'
// ele组件
import type { FormProps } from 'element-plus'
import { UserFilled,Search } from '@element-plus/icons-vue'
import {useSearchEntryStore} from "../../store/SearchSubject.js";
// 获取路由实例
const router = useRouter()
// 折叠面板
const activeNames = ref(['1'])
const handleCollapseChange = (activeNames) => {
  // console.log('激活的面板：', activeNames)
}
// 表单标签位置
const labelPosition = ref<FormProps['labelPosition']>('top')

// 条目表单
const SubjectForm = reactive({
  keywords: 'Nier'
})

// 用户表单
const UserForm = reactive({
  // 默认输入violetmail
  username: 'violetmail'
})

// 用户名规则验证器
const UserFormValidator = (rule: any, value: any, callback: any) => {
  if (value === '') {
    callback(new Error('Please input '))
  }
  callback()
}

const rules = {
  username: [
    {
      validator: UserFormValidator,
      message: '用户名还没有填写',
      trigger: 'blur'
    }
  ]
}
// 条目查询按钮
const SubjectFormRef = ref()
// 使用查询结果的store，读取查询类型
const searchRes = useSearchEntryStore()
const SubMit_subject = () => {
// 条目查询不为空
  if (SubjectForm.keywords !== '') {
    // 跳转到条目结果页面
    router.push({
      name: 'SubjectSearch',
      params: {keywords:SubjectForm.keywords},
      query: {cat: searchRes.getSearchType(),page:1}
    })
  }
}
// 用户查询按钮
const UserFormRef = ref()
const SubMit_user = () => {
  // 用户表单验证
  UserFormRef.value.validate(async (valid) => {
    if (valid) {
      // 如果是 Bangumi 用户，跳转到 Bangumi 用户页面
      await router.push({
        name: 'SearchUser',
        params: {user: UserForm.username}
      })
    }
  })
}

</script>

<style scoped>
.box-card {
  max-height: 500px;
  max-width: 670px;
  min-width: 200px;
  text-align: center;
}
</style>
