import Vue from 'vue'
import VueRouter from 'vue-router'
import LoginPrc from '../views/LoginPrc'
import AdminPrc from '../views/AdminPrc'

// 页面路由组件
import IndexPrc from '../components/admin/IndexPrc'
import AddArt from '../components/article/AddArt'
import ArtList from '../components/article/ArtList'
import CateList from '../components/category/CateList'
import UserList from '../components/user/UserList'

Vue.use(VueRouter)

const routes = [
  {
    path: '/login',
    name: 'login',
    component: LoginPrc
  },
  {
    path: '/',
    name: 'admin',
    component: AdminPrc,
    children: [
      { path: 'index', name: 'index', component: IndexPrc },  // dashboard页面
      { path: 'addart', name: 'addart', component: AddArt },  // 新增文章页面
      { path: 'editart/:id', name: 'editart', component: AddArt, props: true },  // 文章列表跳转新增文章页面
      { path: 'artlist', name: 'artlist', component: ArtList },  // 文章列表页面
      { path: 'catelist', name: 'catelist', component: CateList },  // 分类列表页面
      { path: 'userlist', name: 'userlist', component: UserList }  // 用户列表页面
    ]
  }
]

const router = new VueRouter({
  routes
})

router.beforeEach((to, from, next) => {
  const token = window.sessionStorage.getItem('token')
  if (to.path === '/login') {
    next()
  } else if (!token) {
    next('/login')
  } else {
    next()
  }
})

export default router
