<template>
  <div>
    <h1>Register</h1>
    <form @submit.prevent="register">
      <input type="text" v-model="username" placeholder="Username" required />
      <input type="email" v-model="email" placeholder="Email" required />
      <input type="password" v-model="password" placeholder="Password" required />
      <select v-model="role" required>
        <option value="user">User</option>
        <option value="admin">Admin</option>
      </select>
      <button type="submit">Register</button>
    </form>
    <div v-if="error" class="error">{{ error }}</div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import axios from 'axios';

const username = ref('');
const email = ref('');
const password = ref('');
const role = ref('user');
const error = ref(null);

const register = async () => {
  error.value = null;
  try {
    console.log('Sending registration request');
    const response = await axios.post('/api/register', {
      username: username.value,
      email: email.value,
      password: password.value,
      role: role.value
    });
    console.log('Registration successful:', response.data);
    alert('Registration successful');
  } catch (err) {
    console.error('Registration failed:', err);
    if (err.response) {
      error.value = `Registration failed: ${err.response.data.message || err.response.data.error || 'Unknown error'}`;
    } else if (err.request) {
      error.value = 'Registration failed: No response from server';
    } else {
      error.value = `Registration failed: ${err.message}`;
    }
  }
};

</script>

<style>
.error {
  color: red;
}
</style>
