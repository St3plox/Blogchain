<template>
  <div class="form-container">
    <h2>Login</h2>
    <form @submit.prevent="login">
      <input v-model="email" type="email" placeholder="Email" required />
      <input v-model="password" type="password" placeholder="Password" required />
      <button type="submit" class="login-button">Login</button>
    </form>
  </div>
</template>

<script>
import axios from '@/axios';
import '@/assets/css/form-styles.css'; // Import the shared styles

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
        const response = await axios.post('/users/login', {
          email: this.email,
          password: this.password
        });

        // Extract the Authorization header
        const token = response.headers['authorization'];
        
        if (token) {
          localStorage.setItem('token', token);
          localStorage.setItem('username', response.data.name); // Corrected from setUsername
          alert('Login successful!');
          
          // Redirect to the desired route, e.g., homepage or dashboard
          this.$router.push('/').then(() => {
            window.location.reload(); // Reload the page after redirect
          });
        } else {
          alert('Login failed.');
        }
      } catch (error) {
        alert('Login failed.');
      }
    }
  }
};
</script>
