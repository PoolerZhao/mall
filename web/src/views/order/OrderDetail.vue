<template>
  <el-row>
    <el-col :span="14" :offset="5"><br><br>
      <el-descriptions direction="vertical" :column="5" border>
        <el-descriptions-item label="订单编号">{{order.id}}</el-descriptions-item>
        <el-descriptions-item label="提交时间">
          <i class="el-icon-time"></i>
          <span style="margin-left: 10px">{{ order.created }}</span>
        </el-descriptions-item>
        <el-descriptions-item label="支付状态" :span="2">
          <el-tag size="mini" v-if="order.paymentStatus === 1" type="primary">已支付</el-tag>
          <el-tag size="mini" v-if="order.paymentStatus === 0" type="success">未支付</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="订单状态" :span="2">{{order.status}}</el-descriptions-item>
      </el-descriptions><br>
    </el-col>
    <el-col :span="14" :offset="5">
      <el-descriptions direction="vertical" :column="5" border>
        <el-descriptions-item label="收货人">{{order.name}}</el-descriptions-item>
        <el-descriptions-item label="手机号码">{{order.mobile}}</el-descriptions-item>
        <el-descriptions-item label="收货地址" :span="2">
          {{order.province + order.city + order.district + order.detailedAddress}}
        </el-descriptions-item>
      </el-descriptions><br>
    </el-col>
    <el-col :span="14" :offset="5">
      <el-descriptions direction="vertical" border>
        <el-descriptions-item label="商品信息" with="300">
          <el-table
              ref="multipleTable"
              :data="productItem"
              stripe border>
            <el-table-column
                prop="imageUrl"
                label="图片"
                width="200">
              <template #default="scope">
                <el-image
                    style="width: 50px; height: 50px"
                    :src="scope.row.imageUrl"></el-image>
              </template>
            </el-table-column>
            <el-table-column
                prop="name"
                label="名称"
                width="220">
            </el-table-column>
            <el-table-column
                prop="price"
                label="价格"
                width="200">
            </el-table-column>
          </el-table><br>
        </el-descriptions-item>
      </el-descriptions><br>
      <el-descriptions direction="vertical" border>
        <el-descriptions-item label="合计" with="100">
          ¥ {{order.totalPrice}}
        </el-descriptions-item>
      </el-descriptions><br><br><br>
    </el-col>
  </el-row>
</template>

<script>
export default {
  name: "OrderDetail",
  data() {
    return {
      order: {
        id: 0,
        created: '',
        paymentStatus: 0,
        status: '',
        name: '',
        mobile: '',
        province: '',
        city: '',
        district: '',
        detailedAddress: '',
        totalPrice: ''
      },
      productItem: null
    }
  },
  mounted() {
    this.getOrderDetail();
    this.order.totalPrice = this.$route.params.totalPrice
  },
  methods: {
    getOrderDetail() {
      this.$axios.get('/order/detail',{
        params: {
          id: this.$route.params.id
        }
      }).then(response => {
        let res = response.data.data;
        this.order.id =  res.id;
        this.order.createTime = res.createTime;
        this.order.paymentStatus = res.paymentStatus;
        this.order.status = res.status;
        this.order.name = res.name;
        this.order.mobile = res.mobile;
        this.order.province = res.province;
        this.order.city = res.city;
        this.order.district = res.district;
        this.order.detailedAddress = res.detailedAddress;
        this.productItem = res.productItem;
      })
    }
  }
}
</script>

<style scoped>

</style>