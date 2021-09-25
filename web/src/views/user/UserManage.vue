<template>
  <el-row :gutter="12">
    <el-col :span="24">
      <el-card shadow="never" style="border-top: none;">
        <el-table
            :data="tableData"
            border
            style="width: 100%">
          <el-table-column
              prop="id"
              label="编号"
              width="100">
          </el-table-column>
          <el-table-column
              prop="avatar"
              label="头像"
              width="100">
            <template #default="scope">
              <el-avatar :src="scope.row.avatar"></el-avatar>
            </template>
          </el-table-column>
          <el-table-column
              prop="username"
              label="用户名"
              width="120">
          </el-table-column>
          <el-table-column
              prop="role"
              label="角色">
            <template #default="scope">
              <el-tag size="mini" type="success" v-if="scope.row.role === 'ROLE_admin'">管理员</el-tag>
              <el-tag size="mini" v-if="scope.row.role === '普通用户'">{{scope.row.role}}</el-tag>
            </template>
          </el-table-column>
          <el-table-column
              prop="status"
              label="状态">
            <template #default="scope">
              <el-tag size="mini" type="success" v-if="scope.row.status === 1">正常</el-tag>
              <el-tag size="mini" type="danger" v-if="scope.row.status === 0">停用</el-tag>
            </template>
          </el-table-column>
          <el-table-column
              prop="created"
              label="创建时间"
              width="180">
            <template #default="scope">
              <i class="el-icon-time"></i>
              <span style="margin-left: 10px">{{ scope.row.createTime }}</span>
            </template>
          </el-table-column>
          <el-table-column label="操作" width="220">
            <template #default="scope">
              <el-button
                  size="mini"
                  type="primary"
                  @click="updateClick(scope.$index, scope.row)">修改
              </el-button>
              <el-button
                  size="mini"
                  type="danger"
                  @click="deleteUser(scope.$index, scope.row)">删除
              </el-button>
            </template>
          </el-table-column>
        </el-table>
        <el-dialog title="修改用户信息" v-model="dialogFormVisible" width="30%">
          <el-form :model="user">
            <el-form-item label="角色" :label-width="formLabelWidth">
              <el-radio v-model="user.role" label="管理员">管理员</el-radio>
              <el-radio v-model="user.role" label="普通用户">普通用户</el-radio>
            </el-form-item>
            <el-form-item label="状态" :label-width="formLabelWidth">
              <el-radio v-model="user.status" :label="value1">正常</el-radio>
              <el-radio v-model="user.status" :label="value2">停用</el-radio>
            </el-form-item>
          </el-form>
          <template #footer>
    <span class="dialog-footer">
      <el-button size="mini" @click="dialogFormVisible = false">取 消</el-button>
      <el-button size="mini" type="primary" @click="updateUser">确 定</el-button>
    </span>
          </template>
        </el-dialog>
      </el-card>
    </el-col>
  </el-row>
</template>

<script>
export default {
  name: "UserManage",
  data() {
    return {
      value: true,
      value1: 1,
      value2: 0,
      tableData: null,
      dialogFormVisible: false,
      formLabelWidth: '120px',
      user: {
        role: '',
        status: 0,
        id: 0
      }
    }
  },
  mounted() {
    this.getUserList();
  },
  methods: {
    getUserList(){
      this.$axios.get("/user/list", {
        params: {
          pageNum: 1,
          pageSize: 10,
        }
      }).then(response => {
          this.tableData = response.data.data;
      })
    },
    updateClick(index, row){
      this.user.role = row.role;
      this.user.status = row.status;
      this.user.id = row.id;
      this.dialogFormVisible = true;
      console.log(index,row);
    },
    updateUser() {
      this.$axios.get("/user/update",{
        params: this.user
      }).then(response => {
        console.log(response);
        this.dialogFormVisible = false;
        this.getUserList();
      })
    },
    deleteUser(index, row) {
      console.log(index,row);
      this.$axios.get("/user/delete",{
        params: {
          id: row.id
        }
      }).then(response => {
        console.log(response);
        this.getUserList();
      })
    }
  }
}
</script>

<style scoped>

</style>