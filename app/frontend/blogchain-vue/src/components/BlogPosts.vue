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
            posts: []
        };
    },
    async created() {
        try {
            const token = getToken();
            const response = await axios.get('/api/posts', {
                headers: {
                    Authorization: `Bearer ${token}`
                }
            });
            this.posts = response.data;
        } catch (error) {
            console.error(error);
            alert('Failed to fetch posts.');
        }
    }
};
</script>