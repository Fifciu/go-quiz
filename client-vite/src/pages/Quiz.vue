<script setup lang="ts">
import { onBeforeMount, ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useUserResultStore } from '../stores/userResult.store';

const props = defineProps({
  id: Number,
  title: String,
  description: String,
  image_url: String,
  alreadyDone: Boolean
});

const isModalActive = ref(false);
const isLoading = ref(false);
const userResultManager = useUserResultStore();
const router = useRouter();
const route = useRoute();

const reinitTest = async () => {
  await userResultManager.start(route.params.testId as any);
  await userResultManager.getQuestions();
  userResultManager.id
};

onBeforeMount(async () => {
  if (userResultManager.id < 1) {
    isLoading.value = true;
    await reinitTest();
    isLoading.value = false;
  }
});
</script>

<template>
  <div class="quiz">
    Hello! {{ userResultManager.questions }}
  </div>
</template>