import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import UserList from '../components/UserList.vue'
import UserDetail from '../components/UserDetail.vue'
import PageList from '../components/PageList.vue'
import PageDetail from '../components/PageDetail.vue'
import PostList from '../components/PostList.vue'
import PostDetail from '../components/PostDetail.vue'
import Login from '../views/Login.vue'
import Register from '../views/Register.vue'
import AdminDashboard from '../views/AdminDashboard.vue'
import EditPage from '../views/EditPage.vue'
import EditPost from '../views/EditPost.vue'

const routes = [
    { path: '/', component: Home },
    { path: '/users', component: UserList },
    { path: '/users/:id', component: UserDetail, props: true },
    { path: '/pages', component: PageList },
    { path: '/pages/:id', component: PageDetail, props: true },
    { path: '/posts', component: PostList },
    { path: '/posts/:id', component: PostDetail, props: true },
    { path: '/login', component: Login },
    { path: '/register', component: Register },
    { path: '/dashboard', component: AdminDashboard },
    { path: '/pages/edit/:id', component: EditPage, props: true },
    { path: '/posts/edit/:id', component: EditPost, props: true },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

export default router
