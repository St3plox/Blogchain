<template>
  <div class="form-container" >
    <h2>Register</h2>
    <form @submit.prevent="register">
      <input v-model="username" placeholder="Username" required />
      <input v-model="email" placeholder="Email " required />
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
      password: '',
      email:    ''
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

input {
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
