<template>
    <div class="form-container">
        <h2>Create New Post</h2>
        <form @submit.prevent="submitPost">
            <input v-model="title" placeholder="Title" required />
            <textarea v-model="content" placeholder="Content" required></textarea>
            <select v-model="category">
                <option :value="0">Blog</option>
                <option :value="1">News</option>
                <option :value="2">Article</option>
            </select>
            <input v-model="tagsInput" placeholder="Tags (comma separated)" />
            <button type="submit">Submit</button>
        </form>
    </div>
</template>

<script>
import axios from '@/axios';
import { getToken } from '@/utils/auth';

export default {
    name: 'PostForm',
    data() {
        return {
            title: '',
            content: '',
            category: 0,
            tagsInput: '',
        };
    },
    methods: {
        async submitPost() {
            try {
                const token = getToken();
                const tags = this.tagsInput.split(',').map(tag => tag.trim());
                const response = await axios.post(
                    '/posts',
                    {
                        title: this.title,
                        content: this.content,
                        category: this.category,
                        tags: tags
                    },
                    {
                        headers: {
                            Authorization: `Bearer ${token}`
                        }
                    }
                );
                alert('Post submitted successfully!');
                console.log(response)
            } catch (error) {
                console.error('Post submission failed:', error);
                alert('Post submission failed.');
            }

        }
    }
};
</script>

<style scoped>
.form-container {
    max-width: 400px;
    margin: 0 auto;
    padding: 20px;
    border: 1px solid #ccc;
    border-radius: 10px;
    box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
}

form {
    display: flex;
    flex-direction: column;
}

input,
textarea,
select {
    margin-bottom: 15px;
    padding: 10px;
    font-size: 16px;
    border: 1px solid #ccc;
    border-radius: 5px;
}

button {
    padding: 10px;
    font-size: 16px;
    background-color: #42b983;
    color: white;
    border: none;
    border-radius: 5px;
    cursor: pointer;
}

button:hover {
    background-color: #369d73;
}
</style>