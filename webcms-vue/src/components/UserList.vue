<template>
  <div>
    <h1>Users</h1>
    <b-table :items="users" :fields="fields" :striped="true" :hover="true">
      <template #cell(email)="data">
        <b-link :to="{ name: 'UserDetail', params: { id: data.item.id } }">{{ data.item.email }}</b-link>
      </template>
      <template #cell(username)="data">
        <span>{{ data.item.username }}</span>
      </template>
      <template #cell(role)="data">
        <span>{{ data.item.role }}</span>
      </template>
    </b-table>
    <div v-if="error" class="error">{{ error }}</div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

const users = ref([])
const fields = ['username', 'email', 'role']
const error = ref(null)

onMounted(async () => {
  try {
    const token = localStorage.getItem('jwtToken')
    if (!token) {
      throw new Error('No token found')
    }

    const response = await axios.get('/api/users', {
      headers: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json'
      }
    })
    users.value = response.data
  } catch (err) {
    error.value = 'Failed to fetch users: ' + err.message
    console.error('Failed to fetch users:', err)
  }
})
</script>

<style>
.error {
  color: red;
}
</style>
