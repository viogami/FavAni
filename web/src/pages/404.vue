<template>
    <div>
      <el-result
        icon="warning"
        title=""
        sub-title="你的页面消失了~"
      >
        <template #extra>
          <el-button type="primary" @click="$router.push('/')">回到首页</el-button>
        </template>
      </el-result>
    </div>
</template>
<script setup>

</script>
