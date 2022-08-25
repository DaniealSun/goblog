<template>
  <div>
    <a-card>
      <a-row :gutter="10">
        <a-col :span="6">
          <a-input-search v-model="queryParam.title" placeholder="输入文章名查找" enter-button allowClear @keyup.enter="getArtList" @search="getArtList" />
        </a-col>
        <a-col :span="6">
          <a-button type="primary" @click="addArt">新增</a-button>
        </a-col>
        <a-col :span="4" :offset="8">
          <a-select style="width: 200px" defaultValue="请选择分类" @change="cateChange">
            <a-select-option v-for="item in cateList" :key="item.id" :value="item.id">{{item.name}}</a-select-option>
          </a-select>
        </a-col>
      </a-row>
      <a-table rowKey="ID" :columns="columns" :pagination="pagination" :dataSource="artList" bordered @change="handleTableChange">
        <span class="artImg" slot="img" slot-scope="img">
          <img :src="img" />
        </span>
        <template slot="action" slot-scope="text, record">
          <div class="actionSlot">
            <a-button type="primary" icon="edit" style="margin-right: 15px" @click="editArt(record)">编辑</a-button>
            <a-button type="danger" icon="delete" style="margin-right: 15px" @click="deleteArt(record)">删除</a-button>
          </div>
        </template>
      </a-table>
    </a-card>
  </div>
</template>

<script>
const columns = [
  {
    title: 'ID',
    dataIndex: 'ID',
    width: '5%',
    key: 'id',
    align: 'center'
  },
  {
    title: '分类',
    dataIndex: 'Category.name',
    width: '10%',
    key: 'Category.name',
    align: 'center'
  },
  {
    title: '文章标题',
    dataIndex: 'title',
    width: '15%',
    key: 'title',
    align: 'center',
  },
  {
    title: '文章描述',
    dataIndex: 'desc',
    width: '20%',
    key: 'desc',
    align: 'center',
  },
  {
    title: '缩略图',
    dataIndex: 'img',
    width: '25%',
    key: 'img',
    align: 'center',
    scopedSlots: { customRender: 'img' },
  },
  {
    title: '操作',
    dataIndex: 'action',
    width: '25%',
    key: 'action',
    align: 'center',
    scopedSlots: { customRender: 'action' },
  },
]

export default {
  name: 'ArtList',
  data() {
    return {
      // 前端分页规则
      pagination: {
        pageSizeOptions: ['5', '10', '20'],
        pageSize: 5,
        total: 0,
        showSizeChanger: true,
        showTotal: (total) => `共${total}条`,
      },
      // 后端分页传参
      queryParam: {
        title: '',
        pagesize: 5,
        pagenum: 1
      },
      artList: [],
      cateList: [],
      columns,
      visible: false,
    }
  },
  created() {
    this.getArtList()
    this.getCateList()
  },
  methods: {
    // 获取分类
    async getCateList() {
      const { data: res } = await this.$http.get('category')
      if (res.status !== 200) {
        return this.$message.error(res.message)
      }
      this.cateList = res.data
      this.pagination.total = res.total
    },
    // 获取文章列表
    async getArtList() {
      const { data: res } = await this.$http.get('article', {
        params: { title: this.queryParam.title, pagesize: this.queryParam.pagesize, pagenum: this.queryParam.pagenum }
      })
      if (res.status !== 200) {
        return this.$message.error(res.message)
      }
      this.artList = res.data
      this.pagination.total = res.total
    },
    // 更改分页
    handleTableChange(pagination) {
      var pager = { ...this.pagination }
      pager.current = pagination.current
      pager.pageSize = pagination.pageSize
      this.queryParam.pagesize = pagination.pageSize
      this.queryParam.pagenum = pagination.current
      // 如果重选每页显示条数，直接跳回到第一页
      if (pagination.pageSize !== this.pagination.pageSize) {
        this.queryParam.pagenum = 1
        pager.current = 1
      }
      this.pagination = pager
      this.getArtList()
    },
    // 删除文章
    deleteArt(record) {
      var id = record.ID
      // 确认删除
      this.$confirm({
        title: '确定删除该文章吗？',
        onOk: async () => {
          const res = await this.$http.delete(`article/${id}`)
          if (res.status !== 200) {
            return this.$message.error(res.message)
          }
          this.$message.success('删除成功')
          this.getArtList()
        },
        // 取消删除
        onCancel: () => {
          this.$message.info('已取消删除')
        },
      })
    },
    // 查询分类下的文章
    cateChange(value) {
      this.getCateArt(value)
    },
    async getCateArt(id) {
      const { data: res } = await this.$http.get(`article/list/${id}`)
      if (res.status !== 200) {
        return this.$message.error(res.message)
      }
      this.artList = res.data
      this.pagination.total = res.total
    },
    addArt() {
      this.$router.push('/addart')
    },
    editArt(record) {
      var id = record.ID
      this.$router.push(`/editart/${id}`)
    }
  }
}
</script>

<style scoped>
.actionSlot {
  display: flex;
  justify-content: center;
}
.artImg {
  width: 100px;
  height: 80px;
}
.artImg img {
  width: 100px;
  height: 80px;
}
</style>