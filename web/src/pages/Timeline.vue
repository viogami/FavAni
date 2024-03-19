<script setup>
import Header from './Home/Header.vue'
import FooterPage from './Home/footer.vue'
import { MoreFilled } from '@element-plus/icons-vue'

const activities = [
  {
    content: 'Future',
    timestamp: '???',
    size: 'large',
    type: 'primary',
    icon: MoreFilled
  },
  {
    content: 'Add the middleware func ',
    timestamp: '2024/2',
    size: 'normal',
    type: 'primary',
    hollow: true
  },
  {
    content: 'complete the search ',
    timestamp: '2024/1/21',
    size: 'normal',
    type: 'primary',
  },
  {
    content: 'Update the userPage',
    timestamp: '2023/11/29',
    size: 'large'
  },
  {
    content: 'Build by Vite',
    timestamp: '2023/11/27',
    color: '#0bbd87'
  },
  {
    content: ' Update to Github',
    timestamp: '2023/10/25',
    size: 'large'
  },
  {
    content: ' Making sense of the Demand',
    timestamp: '2023/10/15',
    type: 'primary',
    hollow: true
  },
  {
    content: 'Initial this Project',
    timestamp: '2023/10/10'
  }
]
</script>

<template>
<!-- 添加页头 -->
<Header :menuIndex="'3'"/>
  <el-row justify="center">
    <h1>Latest Updates</h1>
  </el-row>
  <el-row justify="center" style="height: 460px;">
  <el-timeline>
  <el-timeline-item
      v-for="(activity, index) in activities"
      :key="index"
      :icon="activity.icon"
      :type="activity.type"
      :color="activity.color"
      :size="activity.size"
      :hollow="activity.hollow"
      :timestamp="activity.timestamp"
  >
    {{ activity.content }}
  </el-timeline-item>
  </el-timeline>
  </el-row>

  <footer-page/>
</template>

<style scoped>

</style>
