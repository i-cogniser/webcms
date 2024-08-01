import { createRouter, createWebHistory } from 'vue-router';
import Home from '../views/Home.vue';
import Register from '../views/Register.vue';
import UserList from '../components/UserList.vue';  // Убедитесь, что здесь правильный импорт
import UserDetail from '../components/UserDetail.vue';
import DeleteUser from '../components/DeleteUser.vue';


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
        path: '/users/:id',
        name: 'UserDetail',
        component: UserDetail,
        props: true
    },
    {
        path: '/delete-user/:id',
        name: 'DeleteUser',
        component: DeleteUser,
        props: true
    },
];

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL), // Используйте import.meta.env.BASE_URL для Vite
    routes
});

export default router;
