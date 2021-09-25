<template>
  <el-steps :active="active" finish-status="success" align-center>
    <el-step title="基本信息"></el-step>
    <el-step title="图文描述"></el-step>
    <el-step title="物流信息"></el-step>
    <el-step title="售后服务"></el-step>
  </el-steps>
</template>

<script>
export default {
  name: "Steps",
  computed: {
    active: {
      get() {
        return this.$store.state.stepsActive
      },
      set(val) {
        this.$store.state.stepsActive = val
      }
    }
  },
}
</script>

<style scoped>

</style>