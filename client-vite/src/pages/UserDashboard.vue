<script setup>
import { useUserStore } from '../stores/user.store';
import { onMounted } from 'vue';

onMounted(async () => {
  const { me } = useUserStore();
  await me();
});
</script>

<template>
You are logged in
</template>