<template>
    <div class="post">
        <h2>{{ post.title }}</h2>
        <p>{{ post.author }}</p>
        <p>{{ categoryName }}</p>
        <p>{{ formattedTimestamp }}</p>
        <p>{{ post.tags.join(', ') }}</p>
    </div>
</template>

<script>
import { categoryNames } from '@/scripts/post_model';

export default {
    name: 'PostComponent',
    props: {
        post: Object
    },
    computed: {
        categoryName() {
            return categoryNames[this.post.category] || 'Unknown';
        },
        formattedTimestamp() {
            const date = new Date(this.post.timestamp);
            return date.toLocaleDateString() + ' ' + date.toLocaleTimeString();
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
</style>
