<template>
  <el-result icon="success" class="result" :title="title" subTitle="请根据提示进行操作">
    <template #extra>
      <el-button type="primary" size="medium" @click="toIndex">返回</el-button>
    </template>
  </el-result>
</template>

<script>
export default {
  name: "ResultPage",
  mounted() {
    this.title = this.$route.params.title
    console.log("AAAAa" + this.$route.params.title)
  },
  computed: {
    title: {
      get() {
        return this.$store.state.pageTitle
      },
    },
  },
  methods: {
    toIndex() {
      this.$router.push({ path: '/product/list' })
    }
  }
}
</script>

<style scoped>
.result{
  margin-top: 100px;
}
</style>