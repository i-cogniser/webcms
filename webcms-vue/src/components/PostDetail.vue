<template>
  <div v-if="post">
    <h1>Post Details</h1>
    <p><strong>Title:</strong> {{ post.title }}</p>
    <p><strong>Content:</strong> {{ post.content }}</p>
    <b-button to="/posts" variant="secondary">Back to Posts</b-button>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import { useRoute } from 'vue-router'

const post = ref(null)
const route = useRoute()

onMounted(async () => {
  try {
    const response = await axios.get(`/api/posts/${route.params.id}`)
    post.value = response.data
  } catch (error) {
    console.error('Failed to fetch post details:', error)
  }
})
</script>
