<template>
  <div>
    <h2>Register</h2>
    <form @submit.prevent="register">
      <input v-model="username" placeholder="Username" required />
      <input v-model="password" type="password" placeholder="Password" required />
      <button type="submit">Register</button>
    </form>
  </div>
</template>

<script>
import axios from '@/axios';

export default {
  name: 'UserRegister',
  data() {
    return {
      username: '',
      password: ''
    };
  },
  methods: {
    async register() {
      try {
        await axios.post('/api/users/register', {
          username: this.username,
          password: this.password
        });
        alert('Registration successful!');
      } catch (error) {
        console.error(error);
        alert('Registration failed.');
      }
    }
  }
};
</script>
