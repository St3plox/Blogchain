import Vue from 'vue';
import Router from 'vue-router';
import UserRegister from '@/components/UserRegister.vue';
import UserLogin from '@/components/UserLogin.vue';
import BlogPosts from '@/components/post/BlogPosts.vue';
import AllPosts from '@/components/post/AllPosts.vue';
import HomePage from '@/components/HomePage.vue';

Vue.use(Router);

const routes = [
    { path: '/register', component: UserRegister },
    { path: '/login', component: UserLogin },
    { path: '/posts', component: BlogPosts },
    { path: '/posts/all', component: AllPosts },
    { path: '/', component: HomePage }
];

const router = new Router({
    mode: 'history',
    routes
});

export default router;
