<template>
  <div>
    <h1>Pages</h1>
    <b-table :items="pages" :fields="fields" :striped="true" :hover="true">
      <template #cell(title)="data">
        <b-link :to="{ name: 'PageDetail', params: { id: data.item.id } }">{{ data.item.title }}</b-link>
      </template>
    </b-table>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

const pages = ref([])
const fields = ['title', 'author']

onMounted(async () => {
  try {
    const response = await axios.get('/api/pages')
    pages.value = response.data
  } catch (error) {
    console.error('Failed to fetch pages:', error)
  }
})
</script>
