<template>
  <div class="form-container">
    <h2>Register</h2>
    <form @submit.prevent="register">
      <input v-model="username" placeholder="Username" required />
      <input v-model="email" type="email" placeholder="Email" required />
      <input v-model="password" type="password" placeholder="Password" required />
      <button type="submit" class="register-button">Register</button>
    </form>
  </div>
</template>

<script>
import axios from '@/axios';
import '@/assets/css/form-styles.css';

export default {
  name: 'UserRegister',
  data() {
    return {
      username: '',
      password: '',
      email: ''
    };
  },
  methods: {
    async register() {
      try {
        await axios.post('/users/register', {
          username: this.username,
          password: this.password,
          email:    this.email
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
