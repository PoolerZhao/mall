<template>
  <el-container class="el-container">
    <el-aside width="200px" class="el-aside">
      <Menu />
    </el-aside>
    <el-container style="display: block;">
      <Header />
      <el-main class="el-main">
        <el-tabs v-model="editableTabsValue" type="card" closable @tab-remove="removeTab" @tab-click="clickTab">
          <el-tab-pane
              v-for="(item) in editableTabs"
              :key="item.name"
              :label="item.title"
              :name="item.name"
          >
            <router-view />
          </el-tab-pane>
        </el-tabs>
      </el-main>
    </el-container>
  </el-container>
</template>

<script>
import Menu from "../components/Menu";
import Header from "../components/Header";
export default {
  name: "Home",
  components: {Header, Menu},
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
.el-container {
  height: 100vh;
  width: 100vw;
}
.el-aside {
  /*background-color: #334761;*/
  /*background-color: #EDF1FF;*/
  background-color: #324054;
}
.el-main {
  color: #333;
  height: 90%;
  padding: 0;
  background-color: white;
}
</style>