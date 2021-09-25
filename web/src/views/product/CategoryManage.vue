<template>
  <el-main>
    <!-- 查询、创建模块 -->
    <el-card shadow="never" class="card-box">
      <el-form ref="form" :model="sort">
        <el-form-item>
          <el-input size="small"
                    v-model="selectCondition.name"
                    style="width: 20%;" placeholder="按类目名称查询"></el-input>
          <el-select v-model="selectCondition.value" size="small" style="width: 20%;margin-left: 10px;"
                     placeholder="按类目等级查询">
            <el-option
                v-for="item in options"
                :key="item.value"
                :label="item.label"
                :value="item.value">
            </el-option>
          </el-select>
          <el-button size="small" type="primary" @click="queryList" style="margin-left: 28px;">查询</el-button>
          <el-button size="small" type="primary" @click="createClick" style="margin-left: 18px;">新建</el-button>
        </el-form-item>
      </el-form>
    </el-card>
    <!-- 类目列表模块 -->
      <el-card shadow="never" style="border-bottom: none;">
        <el-table
            :data="tableData"
            @selection-change="handleSelectionChange"
            stripe
            style="width: 100%" border>
          <el-table-column
              type="selection"
              width="55">
          </el-table-column>
          <el-table-column
              prop="id"
              label="编号"
              width="100">
          </el-table-column>
          <el-table-column
              prop="name"
              label="类目名称"
              width="150">
          </el-table-column>
          <el-table-column
              prop="level"
              label="类目等级"
              width="150">
          </el-table-column>
          <el-table-column
              prop="sort"
              label="排序"
              width="150">
          </el-table-column>
          <el-table-column
              label="操作">
            <template #default="scope">
              <el-button
                  size="mini"
                  @click="updateClick(scope.$index, scope.row)">修改
              </el-button>
              <el-button
                  size="mini"
                  type="primary"
                  @click="nextLevelCategory(scope.$index, scope.row)">下一级类目
              </el-button>
              <el-button
                  size="mini"
                  type="danger"
                  @click="deleteCategory(scope.$index, scope.row)">删除
              </el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-card>
    <!-- 分页查询模块 -->
      <el-card shadow="never" style="border-top: none;">
        <el-pagination
            background
            @current-change="handleCurrentChange"
            @prev-click="handleCurrentChangePrev"
            @next-click="handleCurrentChangeNext"
            :currentPage="currentPage"
            :page-size="size"
            layout="total, prev, pager, next"
            :total="total">
        </el-pagination>
      </el-card>
    <!-- 弹出框模块 -->
    <el-dialog :title="dialogTitle" width="30%" v-model="dialogFormVisible" center>
      <el-divider>{{ category.levelTitle }}</el-divider>
      <el-form :model="category">
        <el-form-item label="类目名称" :label-width="formLabelWidth">
          <el-input v-model="category.name" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="类目排序" :label-width="formLabelWidth">
          <el-input v-model.number="category.sort" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <template #footer>
    <span class="dialog-footer">
      <el-button @click="dialogFormVisible = false">取 消</el-button>
      <el-button type="primary" v-show="updateButton" @click="updateCategory">确 定</el-button>
      <el-button type="primary" v-show="createButton" @click="createCategory">下一级</el-button>
    </span>
      </template>
    </el-dialog>
  </el-main>
</template>

<script>
export default {
  name: "CategoryManage",
  data() {
    return {
      options: [{
        value: '1',
        label: '一级类目'
      }, {
        value: '2',
        label: '二级类目'
      }, {
        value: '3',
        label: '三级类目'
      }],
      selectCondition: {
        name: '',
        value: '1',
      },
      tableData: '',
      category: {
        id: '',
        name: '',
        level: 1,
        sort: '',
        parentId: 0,
        levelTitle: ''
      },
      multipleSelection: [],
      total: null,
      currentPage: 1,
      size: 5,
      dialogFormVisible: false,
      formLabelWidth: '80px',
      dialogTitle: '',
      updateButton: false,
      createButton: false
    }
  },
  mounted() {
    this.queryList();
  },
  methods: {
    handleSelectionChange(val) {
      this.multipleSelection = val;
    },
    handleCurrentChangePrev(val) {
      this.currentPage = val;
      console.log(`上一页: ${val}`);
    },
    handleCurrentChange(val) {
      this.currentPage = val;
      this.queryList();
      console.log(`当前页: ${val}`);
    },
    handleCurrentChangeNext(val) {
      this.currentPage = val;
      console.log(`下一页: ${val}`);
    },
    // 下一级类目
    nextLevelCategory(index, row) {
      console.log(index)
      this.$axios.get('/category/list', {
        params: {
          pageNum: this.currentPage,
          pageSize: this.size,
          parentId: row.id
        }
      }).then((response) => {
        if (response.data.code === 200) {
          this.total = response.data.data.total;
          this.tableData = response.data.data.list;
        } else {
          this.$message.error('查询失败');
        }
      }).catch((error) => {
        console.log(error);
      })
    },

    // 更新类目
    updateClick(index, row) {
      console.log(index, row);
      this.dialogFormVisible = true;
      this.dialogTitle = '修改类目';
      this.updateButton = true;
      this.createButton = false;
      this.category.id = row.id;
    },

    // 创建类目
    createClick() {
      this.dialogFormVisible = true;
      this.dialogTitle = '新建品牌';
      this.createButton = true;
      this.updateButton = false;
      this.category.levelTitle = '一级类目';
    },

    // 更新类目
    updateCategory() {
      this.$axios.put('/category/update', {
        id: this.category.id,
        name: this.category.name,
        sort: this.category.sort
      }).then((response) => {
        if (response.data.code === 200) {
          console.log(response)
          this.category.id = '';
          this.category.name = '';
          this.category.sort = '';
        }
        this.dialogFormVisible = false;
        this.queryList();
      }).catch((error) => {
        console.log(error);
      })
    },

    // 创建类目
    createCategory() {
      this.$axios.post('/category/create', {
        name: this.category.name,
        level: this.category.level,
        sort: this.category.sort,
        parentId: this.category.parentId
      }).then((response) => {
        if (response.data.code === 200) {
          this.category.level++;
          this.category.name = '';
          this.category.sort = '';
          if (this.category.level === 2) {
            this.category.levelTitle = '二级类目';
          }
          if (this.category.level === 3) {
            this.category.levelTitle = '三级类目';
          }
          this.category.parentId = response.data.data;
          if (this.category.level === 4) {
            this.dialogFormVisible = false;
            this.queryList();
          }
        }
      }).catch((error) => {
        console.log(error);
      })
    },

    // 查询
    queryList() {
      this.$axios.get('/category/list', {
        params: {
          pageNum: this.currentPage,
          pageSize: this.size,
          name: this.selectCondition.name,
          level: this.selectCondition.value
        }
      }).then((response) => {
        this.total = response.data.data.total;
        this.tableData = response.data.data.list;
      }).catch((error) => {
        console.log(error);
      })
    },

    // 删除
    deleteCategory(index, row) {
      this.$axios.get('/category/delete', {
        params: {
          id: row.id
        }
      }).then((response) => {
        if (response.data.code === 200) {
          this.$message({
            message: '删除成功',
            type: 'success'
          });
          this.queryList();
        } else {
          this.$message.error('删除失败');
        }
      }).catch((error) => {
        console.log(error);
      })
    }
  }
}
</script>

<style scoped>
.card-box {
  background-color: #F2F4F7;
  margin: 18px;
  border-radius: 6px;
}
</style>