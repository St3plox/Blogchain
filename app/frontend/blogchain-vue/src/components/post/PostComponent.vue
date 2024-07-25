<template>
    <router-link :to="postLink" class="post-link">
        <div class="post">
            <h2>{{ post.title }}</h2>
            <p>{{ post.author }}</p>
            <p>{{ categoryName }}</p>
            <p>{{ formattedTimestamp }}</p>
            <p>{{ post.tags.join(', ') }}</p>
        </div>
    </router-link>
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
        },
        postLink() {
            return `/posts/${this.post.id}`;
        }
    }
};
</script>

<style scoped>
.post {
    background: #fff;
    padding: 1.5em;
    cursor: pointer; /* Show pointer cursor to indicate it's clickable */
    transition: background-color 0.3s, box-shadow 0.3s; /* Smooth transition for background color and shadow */
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
    text-decoration: none; /* Remove underline from the link */
    color: inherit; /* Inherit color from the .post class */
}

.post-link:hover .post {
    background-color: #f9f9f9; /* Light grey background on hover */
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1); /* Subtle shadow to lift the element */
}
</style>
