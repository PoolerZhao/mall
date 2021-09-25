<template>
  <el-row>
    <el-col :span="18" :offset="3">
      <el-descriptions direction="vertical" :column="4" border>
        <el-descriptions-item label="基本信息"><br>
          <el-form ref="form" :model="product" label-width="100px">
            <el-form-item label="选择类目">
              <el-cascader v-model="categoryValue"
                           :options="product.options"
                           placeholder="请选择"
                           @change="changeCategoryId"
                           clearable />
            </el-form-item>
            <el-form-item label="商品类型">
                <el-radio v-model="product.kind" :label="value1">全新</el-radio>
                <el-radio v-model="product.kind" :label="value2">二手</el-radio>
            </el-form-item>
            <el-form-item label="商品标题">
              <el-input v-model="product.title"
                        maxlength="30" style="width: 80%;"
                        show-word-limit></el-input>
            </el-form-item>
            <el-form-item label="品牌" label-width="100px">
              <el-select v-model="brandValue"
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
            <el-form-item label="商品名称">
              <el-input v-model="product.name" maxlength="20" style="width: 60%;" show-word-limit></el-input>
            </el-form-item>
            <el-form-item label="售价" label-width="100px">
              <el-input v-model="product.price" style="width: 30%;">
                <template #append>元</template>
              </el-input>
            </el-form-item>
            <el-form-item label="库存" label-width="100px">
              <el-input v-model="product.amount" style="width: 30%;">
                <template #append>件</template>
              </el-input>
            </el-form-item>
          </el-form>
        </el-descriptions-item>
      </el-descriptions>
    </el-col>
    <el-col :span="18" :offset="3"><br><br>
      <el-descriptions direction="vertical" :column="4" border>
        <el-descriptions-item label="图文描述"><br>
          <el-form ref="form" :model="product" label-width="100px">
            <el-form-item label="商品图片">
              <el-row>
                <el-col :span="4" :offset="1" v-show="show">
                <el-image
                    style="width: 100%; height: 100%;"
                    :src="product.imageUrl"></el-image>
                </el-col>
                <el-col :span="4" :offset="1">
              <el-upload
                  action="http://localhost:8800/product/picture/upload"
                  :headers="{'Authorization': getHeader}"
                  list-type="picture-card"
                  limit="1"
                  :on-preview="handlePictureCardPreview"
                  :on-remove="handleRemove"
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
                </el-col>
              </el-row>
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
      <el-descriptions direction="vertical" :column="4" border>
        <el-descriptions-item label="售后服务"><br>
          <el-form ref="form" :model="product" label-width="100px">
            <el-form-item label="售后服务">
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
    <el-col :span="18" :offset="3"><br><br>
    <el-button type="primary" @click="onSubmit">提交商品</el-button><br><br>
    </el-col>
  </el-row>
</template>

<script>
export default {
  name: "EditProduct",
  data() {
    return {
      value1: 1,
      value2: 0,
      show: true,
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
    }
  },
  mounted() {
    this.getCategory();
    this.getBrand();
    this.selectedStatus();
    this.getUpdateInfo();
  },
  computed: {
    getHeader: {
      get() {
        return this.$store.state.token
      },
    },
  },
  methods: {
    changeCategoryId() {
      this.product.categoryId = this.categoryValue[2];
    },
    changeBrandId() {
      this.product.brandId = this.brandValue;
    },
    selectedStatus() {
      this.categoryValue = this.product.categoryId
      this.brandValue = this.product.brandId;
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
        this.options = response.data.data;
        console.log(response.data.data)
      }).catch((error) => {
        console.log(error)
      })
    },
    handleRemove(file, fileList) {
      console.log(file, fileList);
    },
    handlePictureCardPreview(file) {
      this.product.imageUrl = file.url;
      this.dialogVisible = true;
    },
    handleAlbumSuccess(response) {
      // res表示后端的返回值，其中应包含文件上传后的url路径
      if (response.code === 200){
        this.product.imageUrl = response.data;
      }
      console.log(response.data);
      this.show = false;
    },
    getUpdateInfo() {
      this.$axios.get('/product/info', {
        params: {
          id: this.$route.params.id
        }
      }).then(response => {
        let res = response.data.data;
        this.categoryValue = res.categoryId;
        this.product.kind = res.kind;
        this.product.title = res.title;
        this.brandValue = res.brandId;
        this.product.name = res.name;
        this.product.price = res.price;
        this.product.amount = res.amount;
        this.product.imageUrl = res.imageUrl;
        this.product.sendAddress = res.sendAddress;
        this.product.parcelType = res.parcelType;
        this.product.salesService = res.salesService.split(",");
      })
    },
    onSubmit() {
      let salesServiceStr = '';
      for (let i = 0; i < this.product.salesService.length; i++) {
        salesServiceStr = salesServiceStr + this.product.salesService[i] + ',';
      }
      console.log("VVVVVV" + this.$route.params.id)
      this.$axios.put('/product/update', {
        id: this.$route.params.id,
        categoryId: this.product.categoryId,
        kind:  this.product.kind,
        title:  this.product.title,
        brandId:  this.product.brandId,
        name:  this.product.name,
        price:  this.product.price,
        amount:  this.product.amount,
        imageUrl:  this.product.imageUrl,
        sendAddress:  this.product.sendAddress,
        parcelType:  this.product.parcelType,
        salesService: salesServiceStr,
        creatorId: localStorage.getItem('USER_ID')
      }).then((response) => {
        if (response.data.code === 200){
          this.$router.push({path: '/result/page'})
        }
        console.log(response.data.data)
      }).catch((error) => {
        console.log(error)
      })
    }
  }
}
</script>

<style scoped>

</style>