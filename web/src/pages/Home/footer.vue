<script setup>

</script>

<template>
  <div class="container">
  <el-footer class="footer">
    © 2023 FavAni by <a href="https://github.com/viogami">viogami</a>
  </el-footer>
  </div>
</template>

<style scoped>
.footer {
  background: pink;
  text-align: center;
  line-height: 60px;
  position: fixed;
  bottom: 0;
  width: 100%;
}
.footer a {
  text-decoration: none; /* 取消链接的下划线 */
  color: #333; /* 设置链接颜色 */
}

.footer a:hover {
  color: #555; /* 鼠标悬停时改变链接颜色 */
}
</style>
