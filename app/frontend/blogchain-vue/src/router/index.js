import Vue from 'vue';
import Router from 'vue-router';
import UserRegister from '@/components/UserRegister.vue';
import UserLogin from '@/components/UserLogin.vue';
import BlogPosts from '@/components/BlogPosts.vue';

Vue.use(Router);

const routes = [
    { path: '/register', component: UserRegister },
    { path: '/login', component: UserLogin },
    { path: '/posts', component: BlogPosts }
];

const router = new Router({
    mode: 'history',
    routes
});

export default router;
