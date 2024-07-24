import Vue from 'vue';
import Router from 'vue-router';
import UserRegister from '@/components/UserRegister.vue';
import UserLogin from '@/components/UserLogin.vue';
import BlogPosts from '@/components/posts/BlogPosts.vue';
import PostForm from '@/components/posts/PostForm.vue';

Vue.use(Router);

const routes = [
    { path: '/register', component: UserRegister },
    { path: '/login', component: UserLogin },
    { path: '/posts', component: BlogPosts },
    { path: '/', component: PostForm }
];

const router = new Router({
    mode: 'history',
    routes
});

export default router;
