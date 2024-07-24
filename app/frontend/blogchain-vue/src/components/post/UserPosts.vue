<template>
    <div>
        <div class="scrolling-component" ref="scrollComponent">
            <post-component v-for="post in posts" :key="post.id" :post="post" />
        </div>
        <div v-if="loading" class="loading">Loading...</div>
    </div>
</template>

<script>
import { getUserPosts } from '@/scripts/posts_handler';
import BasePostsComponent from './BasePostsComponent.vue';
import PostComponent from './PostComponent.vue';

export default {
    name: 'BlogPosts',
    extends: BasePostsComponent,
    components: {
        PostComponent
    },
    methods: {
        async fetchPosts(page, pageSize) {
            return getUserPosts(page, pageSize);
        }
    }
};
</script>

<style scoped>
</style>
