<template>
  <el-menu :default-active="Index"
           :uniqueOpened="true"
           :style=" { width: '200px' } "
           background-color="#324054"
           text-color="#F3F3F4" router>
    <el-submenu :index="menu.name" v-for="menu in menuList" :key="menu.title">
      <template #title>
        <i :class="menu.icon" class="menu-icon"></i>
        <span>{{menu.title}}</span>
      </template>
      <div class="menu-item-back">
        <el-menu-item :index="item.name" v-for="item in menu.children" :key="item.name" @click="selectMenu(item)">
            <i :class="item.icon" class="menu-icon"></i>
            <span slot="title">{{item.title}}</span>
        </el-menu-item>
      </div>
    </el-submenu>
  </el-menu>
</template>

<script>
export default {
  name: "Menu",
  data() {
    return {
      menuList: [
        {
          name: '2',
          icon: 'el-icon-s-goods',
          title: '商品',
          children: [
            {
              name: '/product/list',
              icon: 'el-icon-goods',
              title: '商品列表',
            },
            {
              name: '/product/add',
              icon: 'el-icon-circle-plus-outline',
              title: '添加商品',
            },
            {
              name: '/product/category',
              icon: 'el-icon-price-tag',
              title: '类目管理',
            },
            {
              name: '/product/brand',
              icon: 'el-icon-price-tag',
              title: '品牌管理',
            }
          ]
        },
        {
          name: '3',
          icon: 'el-icon-s-order',
          title: '订单',
          children: [
            {
              name: '/order/list',
              icon: 'el-icon-tickets',
              title: '订单列表',
            },
            {
              name: '/order/refund',
              icon: 'el-icon-refresh-right',
              title: '退款处理',
            }
          ]
        },
        {
          name: '4',
          icon: 'el-icon-user-solid',
          title: '用户',
          children: [
            {
              name: '/user/manage',
              icon: 'el-icon-coin',
              title: '用户列表',
            },
            {
              name: '/change/password',
              icon: 'el-icon-lock',
              title: '修改密码',
            }
          ]
        }
      ]
    }
  },
  computed: {
    Index: {
      get() {
        return this.$store.state.editableTabsValue
      },
      // set(val) {
      //   this.$store.state.editableTabsValue = val
      // }
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
    selectMenu(item) {
      this.$store.commit("addTab", item)
    }
  }
}
</script>

<style scoped>
.menu-icon {
  margin-bottom: 3px;
  color: #fff;
}
.menu-item-back{
  background-color: #3e4b5e;
}
</style>
<style>
.el-menu-item {
  background: none !important;
  color: #ecebeb;
}
.el-menu-item.is-active {
  background-color: #409eff !important;
  color: #fff;
  border-radius: 3px;
}
.el-submenu__title:hover {
  background: none !important;
}
</style>