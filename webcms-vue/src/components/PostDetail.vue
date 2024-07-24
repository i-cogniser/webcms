<template>
  <div v-if="post">
    <h1>Post Details</h1>
    <p><strong>Title:</strong> {{ post.title }}</p>
    <p><strong>Content:</strong> {{ post.content }}</p>
    <b-button to="/posts" variant="secondary">Back to Posts</b-button>
    <b-button @click="deletePost" variant="danger">Delete Post</b-button>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import { useRoute, useRouter } from 'vue-router'

const post = ref(null)
const route = useRoute()
const router = useRouter()

onMounted(async () => {
  try {
    const response = await axios.get(`/api/posts/${route.params.id}`)
    post.value = response.data
  } catch (error) {
    console.error('Failed to fetch post details:', error)
  }
})

const deletePost = async () => {
  try {
    await axios.delete(`/api/posts/${route.params.id}`)
    router.push('/posts')
  } catch (error) {
    console.error('Failed to delete post:', error)
  }
}
</script>
