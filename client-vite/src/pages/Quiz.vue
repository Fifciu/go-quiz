<script setup lang="ts">
import { computed, onBeforeMount, ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useUserResultStore } from '../stores/userResult.store';
import QuizQuestion from '../components/QuizQuestion.vue';

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
const currentQuestionIndex = ref(0);
const currentQuestion = computed(() => userResultManager.questions[currentQuestionIndex.value]);

const reinitTest = async () => {
  await userResultManager.start(route.params.testId as any);
  await userResultManager.getQuestions();
};

onBeforeMount(async () => {
  if (userResultManager.id < 1) {
    isLoading.value = true;
    await reinitTest();
    isLoading.value = false;
  }
});

const onAnswer = async (answerId: number) => {
  await userResultManager.answer(answerId);
  if (userResultManager.isAlreadyFinished) {
    await userResultManager.finish();
    return router.push({
      name: 'quiz-results',
      params: {
        testId: route.params.testId
      }
    })
  }

  currentQuestionIndex.value++;
};
</script>

<template>
  <div class="quiz">
    <transition name="slide-right" mode="out-in">
      <component 
        v-if="currentQuestion"
        :is="QuizQuestion"
        :key="currentQuestion"
        :questionWithAnswers="currentQuestion"
        @answer="onAnswer"
      />
    </transition>
  </div>
</template>

<style lang="scss">

</style>
