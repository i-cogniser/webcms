<template>
  <div>
    <h1>Posts</h1>
    <b-table :items="posts" :fields="fields" :striped="true" :hover="true">
      <template #cell(title)="data">
        <b-link :to="{ name: 'PostDetail', params: { id: data.item.id } }">{{ data.item.title }}</b-link>
      </template>
    </b-table>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

const posts = ref([])
const fields = ['title', 'author']

onMounted(async () => {
  try {
    const response = await axios.get('/api/posts')
    posts.value = response.data
  } catch (error) {
    console.error('Failed to fetch posts:', error)
  }
})
</script>
