<template>
  <div>
    <a-card>
      <a-row :gutter="10">
        <a-col :span="6">
          <a-button type="primary" @click="addCateVisible = true">新增分类</a-button>
        </a-col>
      </a-row>
      <a-table rowKey="id" :columns="columns" :pagination="pagination" :dataSource="catelist" bordered @change="handleTableChange">
        <template slot="action" slot-scope="text, record">
          <div class="actionSlot">
            <a-button type="primary" icon="edit" style="margin-right: 15px" @click="editCate(record)">编辑</a-button>
            <a-button type="danger" icon="delete" style="margin-right: 15px" @click="deleteCate(record)">删除</a-button>
          </div>
        </template>
      </a-table>
    </a-card>
    <!-- 新增分类 -->
    <a-modal closable destroyOnClose width="20%" title="新增分类" :visible="addCateVisible" @ok="addCateOk" @cancel="addCateCancel">
      <a-form-model :model="newCate" :rules="addCateRules" ref="addCateRef">
        <a-form-model-item label="分类名" prop="name">
          <a-input v-model="newCate.name" allowClear @keyup.enter="addCateOk"></a-input>
        </a-form-model-item>
      </a-form-model>
    </a-modal>
    <!-- 编辑分类 -->
    <a-modal closable destroyOnClose width="20%" title="编辑分类" :visible="editCateVisible" @ok="editCateOk" @cancel="editCateCancel">
      <a-form-model :model="cateInfo" :rules="cateRules" ref="editCateRef">
        <a-form-model-item label="分类名" prop="name">
          <a-input v-model="cateInfo.name" @keyup.enter="editCateOk"></a-input>
        </a-form-model-item>
      </a-form-model>
    </a-modal>
  </div>
</template>

<script>
const columns = [
  {
    title: 'ID',
    dataIndex: 'id',
    width: '10%',
    key: 'id',
    align: 'center'
  },
  {
    title: '分类名',
    dataIndex: 'name',
    width: '20%',
    key: 'name',
    align: 'center'
  },
  {
    title: '操作',
    dataIndex: 'action',
    width: '30%',
    key: 'action',
    align: 'center',
    scopedSlots: { customRender: 'action' },
  },
]

export default {
  name: 'CateList',
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
        pagesize: 5,
        pagenum: 1
      },
      // 新增分类模型
      newCate: {
        name: ''
      },
      // 编辑分类模型
      cateInfo: {
        id: 0,
        name: ''
      },
      // 新增分类输入框校验规则
      addCateRules: {
        name: [{
          validator: (rule, value, callback) => {
            if (this.newCate.name === '') { callback(new Error('请输入分类名')) }
            if ([...this.newCate.name].length < 2 || [... this.newCate.name].length > 12) {
              callback(new Error('分类名应该在2到12位之间'))
            } else {
              callback()
            }
          }, trigger: 'blur'
        }],
      },
      // 编辑分类输入框校验规则
      cateRules: {
        name: [{
          validator: (rule, value, callback) => {
            if (this.cateInfo.name === '') { callback(new Error('请输入分类名')) }
            if ([...this.cateInfo.name].length < 2 || [... this.cateInfo.name].length > 12) {
              callback(new Error('分类名应该在2到12位之间'))
            } else {
              callback()
            }
          }, trigger: 'blur'
        }]
      },
      catelist: [],
      columns,
      visible: false,
      // 新增分类弹窗
      addCateVisible: false,
      // 编辑分类弹窗
      editCateVisible: false,
    }
  },
  created() {
    this.getCateList()
  },
  methods: {
    // 获取分类列表
    async getCateList() {
      const { data: res } = await this.$http.get('category', {
        params: { pagesize: this.queryParam.pagesize, pagenum: this.queryParam.pagenum }
      })
      if (res.status !== 200) {
        return this.$message.error(res.message)
      }
      this.catelist = res.data
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
      this.getCateList()
    },
    // 删除分类
    deleteCate(record) {
      var id = record.id
      // 确认删除
      this.$confirm({
        title: '确定删除该分类吗？',
        onOk: async () => {
          const res = await this.$http.delete(`category/${id}`)
          if (res.status !== 200) {
            return this.$message.error(res.message)
          }
          this.$message.success('删除成功')
          this.getCateList()
        },
        // 取消删除
        onCancel: () => {
          this.$message.info('已取消删除')
        },
      })
    },
    // 新增分类
    addCateOk() {
      this.$refs.addCateRef.validate(async (valid) => {
        if (!valid) {
          return this.$message.error('参数不符合要求，请重新输入')
        }
        const { data: res } = await this.$http.post('category/add', {
          name: this.newCate.name,
        })
        if (res.status !== 200) {
          return this.$message.error(res.message)
        }
        this.$refs.addCateRef.resetFields()
        this.addCateVisible = false
        this.$message.success('添加分类成功')
        this.getCateList()
      })
    },
    addCateCancel() {
      this.$refs.addCateRef.resetFields()
      this.addCateVisible = false
      this.$message.info('新增分类已取消')
    },
    // 编辑分类
    async editCate(record) {
      var id = record.id
      this.editCateVisible = true
      const { data: res } = await this.$http.get(`category/${id}`)
      this.cateInfo = res.data
      this.cateInfo.id = id
    },
    editCateOk() {
      this.$refs.editCateRef.validate(async (valid) => {
        if (!valid) {
          return this.$message.error('参数不符合要求，请重新输入')
        }
        const { data: res } = await this.$http.put(`category/${this.cateInfo.id}`, {
          name: this.cateInfo.name,
        })
        if (res.status !== 200) {
          return this.$message.error(res.message)
        }
        this.$refs.editCateRef.resetFields()
        this.editCateVisible = false
        this.$message.success('更新分类信息成功')
        this.getCateList()
      })
    },
    editCateCancel() {
      this.$refs.editCateRef.resetFields()
      this.editCateVisible = false
      this.$message.info('编辑分类已取消')
    },
  },
}
</script>

<style scoped>
.actionSlot {
  display: flex;
  justify-content: center;
}
</style>