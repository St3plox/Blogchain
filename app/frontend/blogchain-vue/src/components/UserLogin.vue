<template>
  <div>
    <h2>Login</h2>
    <form @submit.prevent="login">
      <input v-model="email" placeholder="Email" required />
      <input v-model="password" type="password" placeholder="Password" required />
      <button type="submit">Login</button>
    </form>
  </div>
</template>

<script>
import axios from '@/axios';

export default {
  name: 'UserLogin',
  data() {
    return {
      email: '',
      password: ''
    };
  },
  methods: {
    async login() {
      try {
        const response = await axios.post('/api/users/login', {
          username: this.email,
          password: this.password
        });
        localStorage.setItem('token', response.data.token);
        alert('Login successful!');
      } catch (error) {
        console.error(error);
        alert('Login failed.');
      }
    }
  }
};
</script>
