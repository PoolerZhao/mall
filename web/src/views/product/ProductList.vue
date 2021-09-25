<template>
  <el-row :gutter="12">
    <el-col :span="24">
      <el-card shadow="never" class="card-box">
        <el-form ref="query" :model="query" label-width="80px">
          <el-row :gutter="12">
            <el-col :span="6">
              <el-form-item label="商品标题" prop="productTitle">
                <el-input v-model="query.productTitle" size="small"></el-input>
              </el-form-item>
              <el-form-item label="商品ID" prop="productId">
                <el-input v-model="query.productId" size="small"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="商品类目" prop="productKindValue">
                <el-cascader size="small"
                             v-model="query.productKindValue"
                             :options="product.options"
                             placeholder="请选择"
                             @focus="getCategory"
                             @change="changeProductKind"
                             clearable></el-cascader>
              </el-form-item>
              <el-form-item label="商品品牌" prop="productBrandValue">
                <el-select size="small"
                           v-model="query.productBrandValue"
                           @focus="getBrand"
                           @change="changeProductBrand"
                           placeholder="请选择">
                  <el-option
                      v-for="item in brand.options"
                      :key="item.id"
                      :label="item.name"
                      :value="item.id">
                  </el-option>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="商品状态" prop="productStateValue">
                <el-select v-model="query.productStateValue"
                           @change="changeProductState"
                           placeholder="请选择" size="small">
                  <el-option value="1" label="已上架" >已上架</el-option>
                  <el-option value="0" label="未上架" >未上架</el-option>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item style="margin-left: 10px;">
                <el-button @click="resetForm('query')" size="small">重置</el-button>
                <el-button type="primary" @click="selectGoodsList" size="small">查询</el-button>
              </el-form-item>
            </el-col>
          </el-row>
        </el-form>
      </el-card>
    </el-col>
    <el-col :span="24">
      <el-card shadow="never" style="border-top: none;">
        <el-table
            :highlight-current-row="true"
            :data="tableData"
            border
            style="width: 100%;">
          <el-table-column
              fixed
              type="selection"
              width="40">
          </el-table-column>
          <el-table-column
              label="商品名称"
              width="300">
            <template #default="scope">
              <div style="display: inline-flex;">
                <div style="padding: 5px;">
                  <el-image style="width: 80px; height: 80px"
                            :src="scope.row.imageUrl">
                  </el-image>
                </div>
                <div style="padding: 5px;">
                  <div style="font-weight: 450;">{{scope.row.title}}</div>
                  <div style="padding-top: 8px;">ID：{{scope.row.id}}</div>
                </div>
              </div>
            </template>
          </el-table-column>
          <el-table-column
              prop="price"
              label="价格"
              width="100">
            <template #default="scope">
              <span>¥ {{scope.row.price}}</span>
            </template>
          </el-table-column>
          <el-table-column
              prop="amount"
              label="库存"
              width="100">
          </el-table-column>
          <el-table-column
              prop="sales"
              label="销量"
              width="100">
          </el-table-column>
          <el-table-column
              label="状态"
              width="70">
            <template #default="scope">
              <el-tag size="mini" v-if="scope.row.status === 1" type="primary">已上架</el-tag>
              <el-tag size="mini" v-if="scope.row.status === 0" type="success">仓库中</el-tag>
            </template>
          </el-table-column>
          <el-table-column
              label="创建时间"
              width="180">
            <template #default="scope">
              <i class="el-icon-time"></i>
              <span style="margin-left: 10px">{{ scope.row.created }}</span>
            </template>
          </el-table-column>
          <el-table-column label="操作" width="147">
            <template #default="scope">
              <el-button
                  size="mini"
                  type="text"
                  @click="editProduct(scope.$index, scope.row)">编辑
              </el-button>
              <el-button
                  size="mini"
                  type="text"
                  style="color: red;"
                  @click="deleteGoods(scope.$index, scope.row)">删除
              </el-button>
              <el-button
                  size="mini"
                  v-if="scope.row.status === 0"
                  @click="upProduct(scope.$index, scope.row)"
                  type="text">上架
              </el-button>
              <el-button
                  size="mini"
                  v-if="scope.row.status === 1"
                  @click="downProduct(scope.$index, scope.row)"
                  type="text">下架
              </el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-card>
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
    </el-col>
  </el-row>
</template>

<script>
export default {
  name: "ProductList",
  data() {
    return {
      query: {
        productTitle: '',
        productId: '',
        productKind: '',
        productBrand: '',
        productState: '',
        productKindValue: '',
        productBrandValue: '',
        productStateValue: '',
      },
      product: {
        options: null
      },
      brand: {
        options: null
      },
      total: null,
      currentPage: 1,
      size: 6,
      tableData: null,
    }
  },
  mounted() {
    this.selectGoodsList();
  },
  methods: {
    changeProductKind() {
      this.query.productKind = this.query.productKindValue[2];
    },
    changeProductBrand() {
      this.query.productBrand = this.query.productBrandValue;
    },
    changeProductState() {
      this.query.productState = this.query.productStateValue;
    },
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
    getCategory() {
      this.$axios.get('/category/option').then((response) => {
        this.product.options = response.data.data;
        console.log(response.data.data)
      }).catch((error) => {
        console.log(error)
      })
    },
    getBrand() {
      this.$axios.get('/brand/option').then((response) => {
        this.brand.options = response.data.data;
        console.log(response.data.data)
      }).catch((error) => {
        console.log(error)
      })
    },
    //按条件查询商品信息，显示商品列表
    selectGoodsList() {
      this.$axios.get('/product/list', {
        params: {
          pageNum: this.currentPage,
          pageSize: this.size,
          title: this.query.productTitle,
          id: this.query.productId,
          categoryId: this.query.productKind,
          brandId: this.query.productBrand,
          creatorId: localStorage.getItem("USER_ID"),
          status: this.query.productState
        }
      }).then((response) => {
        this.total = response.data.data.total;
        this.tableData = response.data.data.list;
      }).catch((error) => {
        console.log(error);
      })
    },
    onSubmit(){
      console.log();
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
    },
    editProduct(index, row) {
      console.log(index)
      this.$router.push({
        name: 'editProduct',
        params: {
          id: row.id
        }
      });
    },
    upProduct(index, row) {
      this.$axios.get('/product/update', {
        params: {
          id: row.id,
          status: 1
        }
      }).then((response) => {
        if (response.data.code === 200){
          this.selectGoodsList();
        }
      }).catch((error) => {
        console.log(error);
      })
    },
    downProduct(index, row) {
      this.$axios.get('/product/update', {
        params: {
          id: row.id,
          status: 0
        }
      }).then((response) => {
        if (response.data.code === 200){
          this.selectGoodsList();
        }
      }).catch((error) => {
        console.log(error);
      })
    },
    deleteGoods(index, row) {
      this.$confirm('此操作将永久删除该信息, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.$axios.delete('/product/delete',{
          params: {
            id: row.id
          }
        }).then((response) => {
          this.selectGoodsList();
          console.log(response)
        }).catch((error) => {
          console.log(error);
        })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        });
      });
      console.log(index,row);
    },
  }
}
</script>

<style scoped>
.card-box{
  background-color: #F2F4F7;
  margin: 18px;
  border-radius: 6px;
}
</style>