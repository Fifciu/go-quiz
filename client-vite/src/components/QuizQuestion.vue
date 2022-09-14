<script setup lang="ts">
import { QuestionWithAnswers }from 'api-client';
import { defineProps, PropType } from 'vue';

const props = defineProps({
  questionWithAnswers: {
    type: Object as PropType<QuestionWithAnswers>,
    required: true
  }
});

const emit = defineEmits(['answer']);

const makeAnswer = async (answerId: number) => {
  // alert(answerId);
  emit('answer', answerId);
};
</script>

<template>
  <div class="question__box">
      <h3 class="question__question">{{ props.questionWithAnswers.content }}</h3>
      <div class="question__answers">
        <button
          v-for="answer in props.questionWithAnswers.answers"
          :key="answer.id"
          class="answer"
          @click="makeAnswer(answer.id)"
        >
          {{ answer.content }}
        </button>
      </div>
    </div>
</template>

<style lang="scss">
.answer {
  background: rgba(0,0,0,.3);
  color: #fff;
  font-size: 18px;
  display: block;
  width: 100%;
  border: none;
  margin-top: 15px;
  border-radius: 10px;
  padding: 10px 0;
  cursor: pointer;
  transition: .5s;

  &:hover {
    background: rgba(0,0,0,.8);
  }
}

.quiz {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 100vh;
}

.question__box {
  max-width: 800px;
  padding: 20px;
  margin: 0 auto;
}
</style>
