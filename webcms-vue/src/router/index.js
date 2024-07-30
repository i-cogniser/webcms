import { createRouter, createWebHistory } from 'vue-router';
import Home from '../views/Home.vue';
import Register from '../views/Register.vue';
import UserList from '../components/UserList.vue';  // Убедитесь, что здесь правильный импорт
import PageList from '../components/PageList.vue';
import PostList from '../components/PostList.vue';
import UserDetail from '../components/UserDetail.vue';
import PageDetail from '../components/PageDetail.vue';
import PostDetail from '../components/PostDetail.vue';
import EditPage from '../views/EditPage.vue';
import EditPost from '../views/EditPost.vue';
import DeleteUser from '../components/DeleteUser.vue';
import DeletePage from '../components/DeletePage.vue';
import DeletePost from '../components/DeletePost.vue';

const routes = [
    {
        path: '/',
        name: 'Home',
        component: Home
    },
    {
        path: '/register',
        name: 'Register',
        component: Register
    },
    {
        path: '/users',
        name: 'UserList',  // Обратите внимание на правильное имя
        component: UserList
    },
    {
        path: '/pages',
        name: 'PageList',
        component: PageList
    },
    {
        path: '/posts',
        name: 'PostList',
        component: PostList
    },
    {
        path: '/users/:id',
        name: 'UserDetail',
        component: UserDetail,
        props: true
    },
    {
        path: '/pages/:id',
        name: 'PageDetail',
        component: PageDetail,
        props: true
    },
    {
        path: '/posts/:id',
        name: 'PostDetail',
        component: PostDetail,
        props: true
    },
    {
        path: '/edit-page/:id',
        name: 'EditPage',
        component: EditPage,
        props: true
    },
    {
        path: '/edit-post/:id',
        name: 'EditPost',
        component: EditPost,
        props: true
    },
    {
        path: '/delete-user/:id',
        name: 'DeleteUser',
        component: DeleteUser,
        props: true
    },
    {
        path: '/delete-page/:id',
        name: 'DeletePage',
        component: DeletePage,
        props: true
    },
    {
        path: '/delete-post/:id',
        name: 'DeletePost',
        component: DeletePost,
        props: true
    },
];

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL), // Используйте import.meta.env.BASE_URL для Vite
    routes
});

export default router;
