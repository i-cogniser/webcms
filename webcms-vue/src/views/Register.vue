<template>
  <div>
    <h1>Register</h1>
    <form @submit.prevent="register">
      <label for="name">Name</label>
      <input id="name" v-model="name" type="text" required />

      <label for="email">Email</label>
      <input id="email" v-model="email" type="email" required />

      <label for="password">Password</label>
      <input id="password" v-model="password" type="password" required />

      <button type="submit">Register</button>

      <div v-if="errorMessage" class="error">{{ errorMessage }}</div>
    </form>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import axios from 'axios'

const name = ref('')
const email = ref('')
const password = ref('')
const errorMessage = ref('')

const register = async () => {
  try {
    await axios.post('/api/register', { name: name.value, email: email.value, password: password.value })
    // Handle successful registration
    alert('Registration successful!')
  } catch (error) {
    errorMessage.value = 'Registration failed: ' + error.response.data.error
  }
}
</script>

<style scoped>
.error {
  color: red;
}
</style>
