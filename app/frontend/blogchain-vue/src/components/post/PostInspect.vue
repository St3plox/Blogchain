<template>
    <div v-if="post" class="post">
        <h2>{{ post.title }}</h2>
        <p>{{ post.author }}</p>
        <p>{{ post.content }}</p>
        <p>{{ categoryName }}</p>
        <p>{{ formattedTimestamp }}</p>
        <p>{{ post.tags.join(', ') }}</p>
    </div>
    <div v-else>
        Loading...
    </div>
</template>

<script>
import { categoryNames } from '@/scripts/post_model';
import { getPostById } from '@/scripts/posts_handler';

export default {
    name: 'PostInspect',
    data() {
        return {
            post: null,
            loading: true
        };
    },
    computed: {
        categoryName() {
            return this.post ? (categoryNames[this.post.category] || 'Unknown') : '';
        },
        formattedTimestamp() {
            if (!this.post) return '';
            const date = new Date(this.post.timestamp);
            return date.toLocaleDateString() + ' ' + date.toLocaleTimeString();
        },
    },
    async created() {
        const postId = this.$route.params.postId;
        try {
            this.post = await getPostById(postId);
        } catch (error) {
            console.error('Error fetching post:', error);
        } finally {
            this.loading = false;
        }
    }
};
</script>

<style scoped>
.post {
    background: #fff;
    padding: 1.5em;
}

.post:not(:last-child) {
    border-bottom: 1px solid #ddd;
}

.post h2 {
    font-size: 1.3em;
    padding-bottom: 0.25rem;
}

.post p {
    color: #888;
}

.post-link {
    text-decoration: none;
    color: inherit;
}
</style>
