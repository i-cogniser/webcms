<template>
  <div v-if="page">
    <h1>Page Details</h1>
    <p><strong>Title:</strong> {{ page.title }}</p>
    <p><strong>Content:</strong> {{ page.content }}</p>
    <b-button to="/pages" variant="secondary">Back to Pages</b-button>
    <b-button @click="deletePage" variant="danger">Delete Page</b-button>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import { useRoute, useRouter } from 'vue-router'

const page = ref(null)
const route = useRoute()
const router = useRouter()

onMounted(async () => {
  try {
    const response = await axios.get(`/api/pages/${route.params.id}`)
    page.value = response.data
  } catch (error) {
    console.error('Failed to fetch page details:', error)
  }
})

const deletePage = async () => {
  try {
    await axios.delete(`/api/pages/${route.params.id}`)
    router.push('/pages')
  } catch (error) {
    console.error('Failed to delete page:', error)
  }
}
</script>
