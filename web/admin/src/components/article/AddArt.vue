<template>
  <div>
    <a-card>
      <h3>{{id? '编辑文章':'新增文章'}}</h3>
      <a-form-model :model="artInfo" ref="artInfoRef" :rules="artInfoRules">
        <a-form-model-item label="文章标题" prop="title">
          <a-input style="width: 300px" v-model="artInfo.title"></a-input>
        </a-form-model-item>
        <a-form-model-item label="文章分类" prop="category_id">
          <a-select style="width: 200px" defaultValue="请选择分类" v-model="artInfo.category_id" @change="cateChange">
            <a-select-option v-for="item in cateList" :key="item.id" :value="item.id">{{item.name}}</a-select-option>
          </a-select>
        </a-form-model-item>
        <a-form-model-item label="文章描述" prop="desc">
          <a-input type="textarea" v-model="artInfo.desc"></a-input>
        </a-form-model-item>
        <a-form-model-item label="文章缩略图" prop="img">
          <a-upload name="file" :action="upUrl" :headers="headers" @change="upChange" listType="picture">
            <a-button>
              <a-icon type="upload" /> 上传文章缩略图
            </a-button>
            <br>
            <template v-if="id">
              <img :src="artInfo.img" style="width: 120px;" height="100px" />
            </template>
          </a-upload>
        </a-form-model-item>
        <a-form-model-item label="文章内容" prop="content">
          <a-input style="height: 300px" type="textarea" v-model="artInfo.content" prop="content"></a-input>
        </a-form-model-item>
        <a-form-model-item>
          <a-button type="danger" style="margin-right: 15px" @click="artOk(artInfo.id)">{{artInfo.id? '更新':'提交'}}</a-button>
          <a-button type="primary" @click="addCancel">取消</a-button>
        </a-form-model-item>
      </a-form-model>
    </a-card>
  </div>
</template>

<script>
import { Url } from '../../plugin/http'
// import Editor from '../editor/indexPrc'

export default {
  name: 'AddArt',
  props: ['id'],
  // components: { Editor },
  data() {
    return {
      artInfo: {
        id: 0,
        title: '',
        category_id: '',
        desc: '',
        content: '',
        img: ''
      },
      cateList: [],
      headers: {},
      upUrl: Url + 'upload',  // 后端上传接口
      artInfoRules: {
        title: [{ required: true, message: '请输入文章标题', trigger: 'blur' }],
        category_id: [{ required: true, message: '请选择文章分类', trigger: 'blur' }],
        desc: [
          { required: true, message: '请输入文章描述', trigger: 'blur' },
          { max: 120, message: '描述最多可写120个字符', trigger: 'blur' },
        ],
        img: [{ required: true, message: '请选择文章缩略图', trigger: 'blur' }],
        content: [{ required: true, message: '请输入文章内容', trigger: 'blur' }],
      }
    }
  },
  created() {
    this.getCateList()
    this.headers = { Authorization: `Bearer ${window.sessionStorage.getItem('token')}` }
    if (this.id) {
      this.getArtInfo(this.id)
    }
  },
  methods: {
    // 查询单个文章信息
    async getArtInfo(id) {
      const { data: res } = await this.$http.get(`article/info/${id}`)
      if (res.status !== 200) {
        return this.$message.error(res.message)
      }
      this.artInfo = res.data
      this.artInfo.id = res.data.ID
    },
    // 获取分类列表
    async getCateList() {
      const { data: res } = await this.$http.get('category')
      if (res.status !== 200) {
        return this.$message.error(res.message)
      }
      this.cateList = res.data
    },
    // 选择分类
    cateChange(value) {
      this.artInfo.category_id = value
    },
    // 上传图片
    upChange(info) {
      if (info.file.status !== 'uploading') {
        console.log(info.file.response)
      }
      if (info.file.status === 'done') {
        this.$message.success('图片上传成功');
        const imgUrl = info.file.response.url
        this.artInfo.img = imgUrl
      } else if (info.file.status === 'error') {
        this.$message.error('图片上传失败');
      }
    },
    // 添加或编辑文章
    artOk(id) {
      this.$refs.artInfoRef.validate(async (valid) => {
        if (!valid) return this.$message.error('参数验证未通过，请按要求录入文章内容')
        if (id === 0) {
          const { data: res } = await this.$http.post(`article/add`, this.artInfo)
          if (res.status !== 200) {
            return this.$message.error(res.message)
          }
          this.$message.success('添加文章成功')
          this.$router.push('/artlist')
        } else {
          const { data: res } = await this.$http.put(`article/${id}`, this.artInfo)
          if (res.status !== 200) {
            return this.$message.error(res.message)
          }
          this.$message.success('编辑文章成功')
          this.$router.push('/artlist')
        }
      })
    },
    addCancel() {
      this.$refs.artInfoRef.resetFields()
      this.$router.push('/artlist')
    }
  }
}
</script>

<style>
</style>