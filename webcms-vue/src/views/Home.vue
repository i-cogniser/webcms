<template>
  <div>
    <h1>Добро пожаловать в WebCMS</h1>
    <b-button to="/users" variant="primary">Управление пользователями</b-button>
    <b-button to="/pages" variant="secondary">Управление страницами</b-button>
    <b-button to="/posts" variant="info">Управление записями</b-button>
    <b-button @click="fetchUserData" variant="success">Загрузить данные пользователя</b-button>
  </div>
</template>

<script setup>
import { computed } from 'vue';
import { useStore } from 'vuex';

const store = useStore();

const userName = computed(() => store.state.user?.name);

const fetchUserData = () => {
  store.dispatch('fetchUser');
};
</script>