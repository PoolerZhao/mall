<template>
  <el-tabs v-model="editableTabsValue" type="card" closable @tab-remove="removeTab" @tab-click="clickTab">
    <el-tab-pane
        v-for="(item) in editableTabs"
        :key="item.name"
        :label="item.title"
        :name="item.name"
        style="background-color: #409EFF;"
    >
    </el-tab-pane>
  </el-tabs>
</template>

<script>
export default {
  name: "Tabs",
  data() {
    return {

    }
  },
  computed: {
    editableTabs: {
      get() {
        return this.$store.state.editableTabs
      },
      set(val) {
        this.$store.state.editableTabs = val
      }
    },
    editableTabsValue: {
      get() {
        return this.$store.state.editableTabsValue
      },
      set(val) {
        this.$store.state.editableTabsValue = val
      }
    }
  },
  methods: {
    removeTab(targetName) {
      let tabs = this.editableTabs;
      console.log(tabs + "AAAA")
      let activeName = this.editableTabsValue;
      if (activeName === 'Index') {
        return
      }
      if (activeName === targetName) {
        tabs.forEach((tab, index) => {
          if (tab.name === targetName) {
            let nextTab = tabs[index + 1] || tabs[index - 1];
            if (nextTab) {
              activeName = nextTab.name;
            }
          }
        });
      }
      this.editableTabsValue = activeName;
      this.editableTabs = tabs.filter(tab => tab.name !== targetName);
      this.$router.push({path: this.$store.state.editableTabsValue})
    },
    clickTab(targetName) {
      console.log(targetName)
      this.$router.push({path: this.$store.state.editableTabsValue})
    }
  }
}
</script>

<style scoped>

</style>