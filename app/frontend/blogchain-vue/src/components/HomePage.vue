<template>
    <div>
        <PostForm v-if="token != nil" />
        <AllPosts />
    </div>
</template>

<script>
import { getToken } from '@/utils/auth';
import AllPosts from './post/AllPosts.vue';
import PostForm from './post/PostForm.vue';

export default {
    name: 'HomePage',
    components: {
        PostForm,
        AllPosts
    },
    data() {
        return {
            token: null,
        };
    },
    created() {
        this.token = getToken();
    }
};
</script>
