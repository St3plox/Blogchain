<template>
    <div>
        <h2>Posts</h2>
        <ul>
            <li v-for="post in posts" :key="post.id">{{ post.title }}</li>
        </ul>
    </div>
</template>

<script>
import axios from '@/axios';
import { getToken } from '@/utils/auth';

export default {
    data() {
        return {
            posts: [],
            page: 0, // Define page and pageSize as data properties if they might change
            pageSize: 1
        };
    },
    async created() {
        await this.fetchPosts();
    },
    methods: {
        async fetchPosts() {
            try {
                const token = getToken();
                const response = await axios.get('/posts', {
                    headers: {
                        Authorization: `Bearer ${token}`
                    },
                    params: {
                        page: this.page,
                        pageSize: this.pageSize
                    }
                });
                this.posts = response.data;
            } catch (error) {
                console.error(error);
                alert('Failed to fetch posts.');
            }
        }
    }
};
</script>
