<template>
  <div>
    <h1>Register</h1>
    <form @submit.prevent="register">
      <input type="text" v-model="username" placeholder="Username" required />
      <input type="email" v-model="email" placeholder="Email" required />
      <input type="password" v-model="password" placeholder="Password" required />
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
const error = ref(null);

const register = async () => {
  error.value = null; // Сброс ошибки перед началом регистрации
  try {
    const response = await axios.post('/api/register', {
      username: username.value,
      email: email.value,
      password: password.value
    });
    alert('Registration successful');
  } catch (err) {
    console.error('Registration failed:', err); // Выводим полную ошибку в консоль для отладки
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
