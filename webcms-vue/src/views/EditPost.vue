<template>
  <div>
    <h1>Edit Post</h1>
    <b-form @submit.prevent="updatePost">
      <b-form-group label="Title">
        <b-form-input v-model="post.title" type="text" required />
      </b-form-group>
      <b-form-group label="Content">
        <b-form-textarea v-model="post.content" rows="10" required />
      </b-form-group>
      <b-form-group label="Author">
        <b-form-input v-model="post.author" type="text" required />
      </b-form-group>
      <b-button type="submit" variant="primary">Update Post</b-button>
      <b-button to="/posts" variant="secondary">Back to Posts</b-button>
    </b-form>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import { useRoute, useRouter } from 'vue-router'

const post = ref({
  title: '',
  content: '',
  author: '',
})

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

const updatePost = async () => {
  try {
    await axios.put(`/api/posts/${route.params.id}`, post.value)
    router.push('/posts')
  } catch (error) {
    console.error('Failed to update post:', error)
  }
}
</script>
