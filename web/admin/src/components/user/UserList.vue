<template>
  <div>
    <a-card>
      <a-row :gutter="10">
        <a-col :span="6">
          <a-input-search v-model="queryParam.username" placeholder="输入用户名查找" enter-button allowClear @search="getUserList" />
        </a-col>
        <a-col :span="6">
          <a-button type="primary" @click="addUserVisible = true">新增</a-button>
        </a-col>
      </a-row>
      <a-table rowKey="ID" :columns="columns" :pagination="pagination" :dataSource="userlist" bordered @change="handleTableChange">
        <span slot="role" slot-scope="data">{{data == 1 ? '管理员':'订阅者'}}</span>
        <template slot="action" slot-scope="text, record">
          <div class="actionSlot">
            <a-button type="primary" icon="edit" style="margin-right: 15px" @click="editUser(record)">编辑</a-button>
            <a-button type="danger" icon="delete" style="margin-right: 15px" @click="deleteUser(record)">删除</a-button>
            <a-button icon="redo" style="margin-right: 15px" @click="redoPassword(record)">重置密码</a-button>
          </div>
        </template>
      </a-table>
    </a-card>
    <!-- 新增用户 -->
    <a-modal closable destroyOnClose width="20%" title="新增用户" :visible="addUserVisible" @ok="addUserOk" @cancel="addUserCancel">
      <a-form-model :model="newUser" :rules="addUserRules" ref="addUserRef">
        <a-form-model-item label="用户名" prop="username">
          <a-input v-model="newUser.username" allowClear></a-input>
        </a-form-model-item>
        <a-form-model-item has-feedback label="密码" prop="password">
          <a-input-password v-model="newUser.password" allowClear></a-input-password>
        </a-form-model-item>
        <a-form-model-item has-feedback label="确认密码" prop="checkpass">
          <a-input-password v-model="newUser.checkpass" allowClear></a-input-password>
        </a-form-model-item>
      </a-form-model>
    </a-modal>
    <!-- 编辑用户 -->
    <a-modal closable destroyOnClose width="20%" title="编辑用户" :visible="editUserVisible" @ok="editUserOk" @cancel="editUserCancel">
      <a-form-model :model="userInfo" :rules="userRules" ref="editUserRef">
        <a-form-model-item label="用户名" prop="username">
          <a-input v-model="userInfo.username"></a-input>
        </a-form-model-item>
        <a-form-model-item label="是否为管理员" prop="role">
          <a-switch :checked="IsAdmin" checked-children="是" un-checked-children="否" @change="adminChange"></a-switch>
        </a-form-model-item>
      </a-form-model>
    </a-modal>
    <!-- 重置密码 -->
    <a-modal closable destroyOnClose width="20%" title="重置密码" :visible="redoPwVisible" @ok="redoPwOk" @cancel="redoPwCancel">
      <a-form-model :model="redoPwInfo" :rules="redoPwRules" ref="redoPwRef">
        <a-form-model-item has-feedback label="旧密码" prop="oldpass">
          <a-input-password v-model="redoPwInfo.oldpass" allowClear placeholder="请输入旧密码"></a-input-password>
        </a-form-model-item>
        <a-form-model-item has-feedback label="新密码" prop="password">
          <a-input-password v-model="redoPwInfo.password" allowClear placeholder="请输入新密码"></a-input-password>
        </a-form-model-item>
        <a-form-model-item has-feedback label="确认密码" prop="checkpass">
          <a-input-password v-model="redoPwInfo.checkpass" allowClear placeholder="请再次输入新密码"></a-input-password>
        </a-form-model-item>
      </a-form-model>
    </a-modal>
  </div>
</template>

<script>
const columns = [
  {
    title: 'ID',
    dataIndex: 'ID',
    width: '10%',
    key: 'id',
    align: 'center'
  },
  {
    title: '用户名',
    dataIndex: 'username',
    width: '20%',
    key: 'username',
    align: 'center'
  },
  {
    title: '角色',
    dataIndex: 'role',
    width: '10%',
    key: 'role',
    align: 'center',
    scopedSlots: { customRender: 'role' },
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
  name: 'UserList',
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
        username: '',
        pagesize: 5,
        pagenum: 1
      },
      // 新增用户模型
      newUser: {
        id: 0,
        username: '',
        password: '',
        checkpass: '',
        role: 2
      },
      // 编辑用户模型
      userInfo: {
        id: 0,
        username: '',
        role: 2
      },
      // 重置密码模型
      redoPwInfo: {
        id: 0,
        oldpass: '', // 旧密码，待与新密码验证是否一致
        password: '',
        checkpass: '',
      },
      // 新增用户输入框校验规则
      addUserRules: {
        username: [{
          validator: (rule, value, callback) => {
            if (this.newUser.username === '') { callback(new Error('请输入用户名')) }
            if ([...this.newUser.username].length < 4 || [... this.newUser.username].length > 12) {
              callback(new Error('用户名应该在4到12位之间'))
            } else {
              callback()
            }
          }, trigger: 'blur'
        }],
        password: [{
          validator: (rule, value, callback) => {
            if (this.newUser.password === '') { callback(new Error('请输入密码')) }
            if ([...this.newUser.password].length < 6 || [... this.newUser.password].length > 20) {
              callback(new Error('密码应该在6到20位之间'))
            } else {
              callback()
            }
          }, trigger: 'blur'
        }],
        checkpass: [{
          validator: (rule, value, callback) => {
            if (this.newUser.checkpass === '') { callback(new Error('请输入密码')) }
            if (this.newUser.password !== this.newUser.checkpass) {
              callback(new Error('密码不一致，请重新输入'))
            } else {
              callback()
            }
          }, trigger: 'blur'
        }]
      },
      // 编辑用户输入框校验规则
      userRules: {
        username: [{
          validator: (rule, value, callback) => {
            if (this.userInfo.username === '') { callback(new Error('请输入用户名')) }
            if ([...this.userInfo.username].length < 4 || [... this.userInfo.username].length > 12) {
              callback(new Error('用户名应该在4到12位之间'))
            } else {
              callback()
            }
          }, trigger: 'blur'
        }]
      },
      // 重置密码输入框校验规则
      redoPwRules: {
        password: [{
          validator: (rule, value, callback) => {
            if (this.redoPwInfo.password === '') { callback(new Error('请输入密码')) }
            if ([...this.redoPwInfo.password].length < 6 || [... this.redoPwInfo.password].length > 20) {
              callback(new Error('密码应该在6到20位之间'))
            } else {
              callback()
            }
          }, trigger: 'blur'
        }],
        checkpass: [{
          validator: (rule, value, callback) => {
            if (this.redoPwInfo.checkpass === '') { callback(new Error('请输入密码')) }
            if (this.redoPwInfo.password !== this.redoPwInfo.checkpass) {
              callback(new Error('密码不一致，请重新输入'))
            } else {
              callback()
            }
          }, trigger: 'blur'
        }]
      },
      userlist: [],
      columns,
      visible: false,
      // 新增用户弹窗
      addUserVisible: false,
      // 编辑用户弹窗
      editUserVisible: false,
      // 重置密码弹窗
      redoPwVisible: false,
    }
  },
  created() {
    this.getUserList()
  },
  computed: {
    IsAdmin: function () {
      if (this.userInfo.role === 1) {
        return true
      } else {
        return false
      }
    },
  },
  methods: {
    // 获取用户列表
    async getUserList() {
      const { data: resGet } = await this.$http.get('users', {
        params: { username: this.queryParam.username, pagesize: this.queryParam.pagesize, pagenum: this.queryParam.pagenum }
      })
      if (resGet.status !== 200) {
        return this.$message.error(resGet.message)
      }
      this.userlist = resGet.data
      this.pagination.total = resGet.total
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
      this.getUserList()
    },
    // 删除用户
    deleteUser(record) {
      var id = record.ID
      // 确认删除
      this.$confirm({
        title: '确定删除该用户吗？',
        onOk: async () => {
          const resDel = await this.$http.delete(`user/${id}`)
          if (resDel.status !== 200) {
            return this.$message.error(resDel.message)
          }
          this.$message.success('删除成功')
          this.getUserList()
        },
        // 取消删除
        onCancel: () => {
          this.$message.info('已取消删除')
        },
      })
    },
    // 新增用户
    addUserOk() {
      this.$refs.addUserRef.validate(async (valid) => {
        if (!valid) {
          return this.$message.error('参数不符合要求，请重新输入')
        }
        const { data: resAdd } = await this.$http.post('user/add', {
          username: this.newUser.username,
          password: this.newUser.password,
          role: this.newUser.role
        })
        if (resAdd.status !== 200) {
          return this.$message.error(resAdd.message)
        }
        this.$refs.addUserRef.resetFields()
        this.addUserVisible = false
        this.$message.success('添加用户成功')
        this.getUserList()
      })
    },
    addUserCancel() {
      this.$refs.addUserRef.resetFields()
      this.addUserVisible = false
      this.$message.info('新增用户已取消')
    },
    // 编辑用户
    async editUser(record) {
      var id = record.ID
      this.editUserVisible = true
      const { data: resEdit } = await this.$http.get(`user/${id}`)
      this.userInfo = resEdit.data
      this.userInfo.id = id
    },
    editUserOk() {
      this.$refs.editUserRef.validate(async (valid) => {
        if (!valid) {
          return this.$message.error('参数不符合要求，请重新输入')
        }
        const resedit = await this.$http.put(`user/${this.userInfo.id}`, {
          username: this.userInfo.username,
          role: this.userInfo.role
        })
        if (resedit.status !== 200) {
          return this.$message.error(resedit.message)
        }
        this.$refs.editUserRef.resetFields()
        this.editUserVisible = false
        this.$message.success('更新用户信息成功')
        this.getUserList()
      })
    },
    editUserCancel() {
      this.$refs.editUserRef.resetFields()
      this.editUserVisible = false
      this.$message.info('编辑用户已取消')
    },
    adminChange(checked) {
      if (checked) {
        this.userInfo.role = 1
      } else {
        this.userInfo.role = 2
      }
    },
    // 重置密码
    redoPassword(record) {
      this.redoPwVisible = true
      this.redoPwInfo.id = record.ID  // redoPwInfo模型写入id为后边redoPwOk方法的put请求做准备
    },
    redoPwOk() {
      this.$refs.redoPwRef.validate(async (valid) => {
        if (!valid) {
          return this.$message.error('参数不符合要求，请重新输入')
        }
        const { data: resredo } = await this.$http.put(`admin/redopass/${this.redoPwInfo.id}`, {
          password: this.redoPwInfo.password
        })
        if (resredo.status !== 200) {
          return this.$message.error(resredo.message)
        }
        this.$refs.redoPwRef.resetFields()
        this.redoPwVisible = false
        this.$message.success('重置密码成功')
        this.getUserList()
      })
    },
    redoPwCancel() {
      this.$refs.redoPwRef.resetFields()
      this.redoPwVisible = false
      this.$message.info('编辑用户已取消')
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