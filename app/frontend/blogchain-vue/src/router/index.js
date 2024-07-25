import Vue from 'vue';
import Router from 'vue-router';
import UserRegister from '@/components/UserRegister.vue';
import UserLogin from '@/components/UserLogin.vue';
import UserPosts from '@/components/post/UserPosts.vue';
import AllPosts from '@/components/post/AllPosts.vue';
import HomePage from '@/components/HomePage.vue';
import PostInspect from '@/components/post/PostInspect.vue';

Vue.use(Router);

const routes = [
    { path: '/register', component: UserRegister },
    { path: '/login', component: UserLogin },

    { path: '/posts', component: UserPosts },
    { path: '/posts/all', component: AllPosts },
    { path: '/posts/:postId', component: PostInspect },

    { path: '/', component: HomePage }
];

const router = new Router({
    mode: 'history',
    routes
});

export default router;
