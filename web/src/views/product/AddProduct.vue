<template>
  <el-row>
    <el-col :span="18" :offset="3"><br>
      <!-- 商品基本信息 -->
      <el-descriptions direction="vertical" :column="4" border>
        <el-descriptions-item label="基本信息"><br>
          <el-form ref="form" :model="product" label-width="100px">
            <el-form-item label="类目">
              <el-cascader v-model="categoryValue"
                           :options="product.options"
                           placeholder="请选择"
                           @focus="getCategoryOption"
                           @change="changeCategoryId"
                           clearable />
            </el-form-item>
            <el-form-item label="类型">
              <el-radio v-model="product.kind" :label="value1">全新</el-radio>
              <el-radio v-model="product.kind" :label="value2">二手</el-radio>
            </el-form-item>
            <el-form-item label="标题">
              <el-input v-model="product.title"
                        maxlength="30" style="width: 80%;"
                        show-word-limit></el-input>
            </el-form-item>
            <el-form-item label="品牌" label-width="100px">
              <el-select v-model="brandValue"
                         @focus="getBrandOption"
                         @change="changeBrandId"
                         placeholder="请选择">
                <el-option
                    v-for="item in options"
                    :key="item.id"
                    :label="item.name"
                    :value="item.id">
                </el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="名称">
              <el-input v-model="product.name"
                        maxlength="20"
                        style="width: 60%;"
                        show-word-limit></el-input>
            </el-form-item>
            <el-form-item label="售价" label-width="100px">
              <el-input v-model.number="product.price" style="width: 30%;">
                <template #append>元</template>
              </el-input>
            </el-form-item>
            <el-form-item label="库存" label-width="100px">
              <el-input v-model.number="product.amount" style="width: 30%;">
                <template #append>件</template>
              </el-input>
            </el-form-item>
          </el-form>
        </el-descriptions-item>
      </el-descriptions>
    </el-col>
    <el-col :span="18" :offset="3"><br><br>
      <!-- 商品图文描述 -->
      <el-descriptions direction="vertical" :column="4" border>
        <el-descriptions-item label="图文描述"><br>
          <el-form ref="form" :model="product" label-width="100px">
            <el-form-item label="商品主图">
                  <el-upload
                      action="http://localhost:8000/upload"
                      :headers="{'token': token}"
                      list-type="picture-card"
                      limit="1"
                      :on-success="handleAlbumSuccess"
                      name="file">
                    <i class="el-icon-plus"></i>
                    <el-dialog v-model="dialogVisible">
                      <el-image
                          style="width: 100%; height: 100%"
                          :src="product.imageUrl"></el-image>
                      <img width="100%" :src="product.imageUrl" alt="">
                    </el-dialog>
                  </el-upload>
            </el-form-item>
            <br>
            <el-form-item>
              <el-alert
                  title="最多上传4张图，首张图为主图，商品主图大小不能超过3MB，支持任意格式的图片"
                  type="info"
                  show-icon>
              </el-alert>
            </el-form-item>
          </el-form>
        </el-descriptions-item>
      </el-descriptions>
    </el-col>
    <el-col :span="18" :offset="3"><br><br>
      <!-- 商品物流信息 -->
      <el-descriptions direction="vertical" :column="4" border>
        <el-descriptions-item label="物流信息"><br>
          <el-form ref="form" :model="product" label-width="100px">
            <el-form-item label="发货地址">
              <el-input v-model="product.sendAddress" style="width: 60%;"></el-input>
            </el-form-item>
            <el-form-item label="快递类型">
              <el-input v-model="product.parcelType" style="width: 60%;"></el-input>
            </el-form-item>
          </el-form>
        </el-descriptions-item>
      </el-descriptions>
    </el-col>
    <el-col :span="18" :offset="3"><br><br>
      <!-- 商品售后服务 -->
      <el-descriptions direction="vertical" :column="4" border>
        <el-descriptions-item label="售后服务"><br>
          <el-form ref="form" :model="product" label-width="100px">
            <el-form-item label="选择服务">
              <el-checkbox-group v-model="product.salesService">
                <el-checkbox label="提供发票" name="type"></el-checkbox>
                <el-checkbox label="保修服务" name="type"></el-checkbox>
                <el-checkbox label="退换货承诺" name="type"></el-checkbox>
                <el-checkbox label="服务承诺：该类商品，必须支持【七天退货】服务" name="type"></el-checkbox>
              </el-checkbox-group>
            </el-form-item>
          </el-form>
        </el-descriptions-item>
      </el-descriptions>
    </el-col>
    <!--    -->
    <el-col :span="24"><br><br>
      <el-form-item style="text-align: center;">
        <el-button type="primary" @click="onSubmit(0)">保存草稿</el-button>
        <el-button type="primary" @click="onSubmit(1)">发布商品</el-button>
      </el-form-item>
    </el-col>
  </el-row>
</template>

<script>

export default {
  name: "AddProduct",
  data() {
    return {
      value1: 1,
      value2: 0,
      categoryValue: null,
      brandValue: null,
      product: {
        options: null,
        categoryId: '',
        kind: '',
        title: '',
        brandId: '',
        name: '',
        price: '',
        amount: '',
        imageUrl: '',
        sendAddress: '',
        parcelType: '',
        salesService: [],
      },
      options: null,
      dialogVisible: false,
      token: ''
    }
  },
  mounted() {
    this.token = localStorage.getItem('token')
  },
  methods: {
    // 选择器更改时触发
    changeCategoryId() {
      this.product.categoryId = this.categoryValue[2];
    },
    // 选择器更改时触发
    changeBrandId() {
      this.product.brandId = this.brandValue;
    },
    // 获取类目选项
    getCategoryOption() {
      this.$axios.get('/category/option').then((response) => {
        this.product.options = response.data.data;
        console.log(response.data.data)
      }).catch((error) => {
        console.log(error)
      })
    },
    // 获取品牌选项
    getBrandOption() {
      this.$axios.get('/brand/option').then((response) => {
        this.options = response.data.data;
        console.log(response.data.data)
      }).catch((error) => {
        console.log(error)
      })
    },
    // 图片响应结果处理
    handleAlbumSuccess(response) {
      if (response.code === 200){
        this.product.imageUrl = response.data;
      }
      console.log(response.data);
    },
    // 创建商品（表单提交）
    onSubmit(status) {
      let salesServiceStr = '';
      for (let i = 0; i < this.product.salesService.length; i++) {
        salesServiceStr = salesServiceStr + this.product.salesService[i] + ',';
      }
      if (status === 0){ this.$store.commit('setPageTitle', '保存成功')}
      else { this.$store.commit('setPageTitle', '发布成功' )}

      this.$axios.post('/product/create', {
        categoryId: this.product.categoryId,
        kind: this.product.kind,
        title: this.product.title,
        brandId: this.product.brandId,
        name: this.product.name,
        price: this.product.price,
        amount: this.product.amount,
        imageUrl: this.product.imageUrl,
        sendAddress: this.product.sendAddress,
        parcelType: this.product.parcelType,
        salesService: salesServiceStr,
        creatorId: localStorage.getItem('USER_ID'),
        status: status
      }).then((response) => {
        if (response.data.code === 200){
          this.$router.push({name: 'resultPage'})
        }
        console.log(response)
      }).catch((error) => {
        console.log(error)
      })
    }
  }
}
</script>

<style scoped>

</style>