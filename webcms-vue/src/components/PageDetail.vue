<template>
  <div v-if="page">
    <h1>Page Details</h1>
    <p><strong>Title:</strong> {{ page.title }}</p>
    <p><strong>Content:</strong> {{ page.content }}</p>
    <b-button to="/pages" variant="secondary">Back to Pages</b-button>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import { useRoute } from 'vue-router'

const page = ref(null)
const route = useRoute()

onMounted(async () => {
  try {
    const response = await axios.get(`/api/pages/${route.params.id}`)
    page.value = response.data
  } catch (error) {
    console.error('Failed to fetch page details:', error)
  }
})
</script>
