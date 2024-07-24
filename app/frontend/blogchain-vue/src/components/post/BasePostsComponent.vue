<template>
    <div>
        <div class="scrolling-component" ref="scrollComponent">
            <post-component v-for="post in posts" :key="post.id" :post="post" />
        </div>
        <div v-if="loading" class="loading">Loading...</div>
    </div>
</template>

<script>
import PostComponent from './PostComponent.vue';

export default {
    name: 'BasePostsComponent',
    components: {
        PostComponent
    },
    data() {
        return {
            posts: [],
            page: 0,
            pageSize: 10,
            loading: false,
            allPostsLoaded: false
        };
    },
    async created() {
        await this.loadPosts();
        window.addEventListener('scroll', this.handleScroll);
    },
    beforeUnmount() {
        window.removeEventListener('scroll', this.handleScroll);
    },
    methods: {
        async loadPosts() {
            if (this.loading || this.allPostsLoaded) return;

            this.loading = true;
            try {
                const newPosts = await this.fetchPosts(this.page, this.pageSize);
                if (newPosts.length > 0) {
                    this.posts = [...this.posts, ...newPosts];
                    this.page += 1;
                } else {
                    this.allPostsLoaded = true; // No more posts to load
                }
            } catch (error) {
                console.error('Failed to fetch posts:', error);
            } finally {
                this.loading = false;
            }
        },
        handleScroll() {
            const bottomOfWindow =
                window.innerHeight + window.scrollY >= document.documentElement.scrollHeight - 200;
            if (bottomOfWindow && !this.loading) {
                this.loadPosts();
            }
        },
        async fetchPosts(page, pageSize) {
            console.log(page);
            console.log(pageSize);
            // This method should be overridden by the extending components
            throw new Error('fetchPosts method not implemented');
        }
    }
};
</script>

<style scoped>
.scrolling-component {
    overflow-y: auto;
    max-height: 80vh;
}

.loading {
    text-align: center;
    padding: 1em;
    font-size: 1.2em;
}
</style>
