<template>
  <el-row :gutter="12">
    <el-col :span="24">
      <el-card shadow="never" style="border-top: none;">
      <el-table
          :highlight-current-row="true"
          :data="tableData"
          border
          style="width: 100%;">
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
            prop="created"
            label="提交时间"
            width="180">
          <template #default="scope">
            <i class="el-icon-time"></i>
            <span style="margin-left: 10px">{{ scope.row.created }}</span>
          </template>
        </el-table-column>
        <el-table-column
            prop="paymentStatus"
            label="支付状态"
            width="100">
          <template #default="scope">
            <el-tag v-if="scope.row.paymentStatus === 1" size="mini" type="success">已支付</el-tag>
            <el-tag v-if="scope.row.paymentStatus === 0" size="mini" type="danger">未支付</el-tag>
          </template>
        </el-table-column>
        <el-table-column
            prop="totalPrice"
            label="总计"
            width="100">
        </el-table-column>
        <el-table-column
            prop="status"
            label="订单状态"
            width="100">
        </el-table-column>
        <el-table-column label="操作" width="300">
          <template #default="scope">
            <el-button
                size="mini"
                @click="checkOrder(scope.$index, scope.row)">查看
            </el-button>
            <el-button
                size="mini"
                type="primary"
                @click="orderDelivery(scope.$index, scope.row)">发货
            </el-button>
            <el-button
                size="mini"
                @click="deleteOrder(scope.$index, scope.row)">退款
            </el-button>
            <el-button
                size="mini"
                type="danger"
                @click="deleteOrder(scope.$index, scope.row)">删除
            </el-button>
          </template>
        </el-table-column>
      </el-table><br><br>
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
    </el-col>
  </el-row>
</template>

<script>
export default {
  name: "OrderList",
  data() {
    return {
      tableData: null,
      currentPage: 1,
      size: 5,
      total: ''
    }
  },
  mounted() {
    this.queryOrder();
  },
  methods: {
    handleCurrentChangePrev(val) {
      this.currentPage = val;
      console.log(`上一页: ${val}`);
    },
    handleCurrentChange(val) {
      this.currentPage = val;
      this.selectGoodsList();
      console.log(`当前页: ${val}`);
    },
    handleCurrentChangeNext(val) {
      this.currentPage = val;
      console.log(`下一页: ${val}`);
    },
    queryOrder() {
      this.$axios.get('/order/list', {
        params: {
          pageNum: this.currentPage,
          pageSize: this.size
        }
      }).then((response) => {
        this.total = response.data.data.total;
        this.tableData = response.data.data.list;
      }).catch((error) => {
        console.log(error);
      })
    },
    checkOrder(index, row) {
      console.log(index)
      this.$router.push({
        name: 'orderDetail',
        params: {
          id: row.id,
          totalPrice: row.totalPrice
        }
      });
    },
    deleteOrder(index, row) {
      console.log(index)
      this.$axios.get('/order/delete', {
        params: {
          orderId: row.id,
        }
      }).then((response) => {
         if (response.data.code === 200){
           this.$message({
             type: 'success',
             message: '删除成功!'
           });
           this.queryOrder();
         }
      }).catch((error) => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        });
        console.log(error);
      })

    },
    orderDelivery(index, row) {
      this.$axios.get('/order/update', {
        params: {
          status: '已发货',
          orderId: row.id,
        }
      }).then((response) => {
        if (response.data.code === 200){
          this.$message({
            type: 'success',
            message: '发货成功!'
          });
          this.queryOrder();
        }
      }).catch((error) => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        });
        console.log(error);
      })
    }
  }
}
</script>

<style scoped>

</style>