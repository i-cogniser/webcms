import { createRouter, createWebHistory } from 'vue-router';
import Home from '../views/Home.vue';
import Users from '../components/UserList.vue';
import Pages from '../components/PageList.vue';
import Posts from '../components/PostList.vue';
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
        path: '/users',
        name: 'Users',
        component: Users
    },
    {
        path: '/pages',
        name: 'Pages',
        component: Pages
    },
    {
        path: '/posts',
        name: 'Posts',
        component: Posts
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
    }
];

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
});

export default router;