<template>
  <div>
    <h1>Users</h1>
    <b-table :items="users" :fields="fields" :striped="true" :hover="true">
      <template #cell(name)="data">
        <b-link :to="{ name: 'UserDetail', params: { id: data.item.id } }">{{ data.item.name }}</b-link>
      </template>
    </b-table>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

const users = ref([])
const fields = ['name', 'email']

onMounted(async () => {
  try {
    const response = await axios.get('/api/users')
    users.value = response.data
  } catch (error) {
    console.error('Failed to fetch users:', error)
  }
})
</script>
