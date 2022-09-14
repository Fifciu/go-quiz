<script setup lang="ts">
import { ref, computed } from '@vue/reactivity';
import { onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { useUserResultStore } from '../stores/userResult.store';
import { client } from 'api-client';

const mock = [
  {
    "question": "Which one is example of composite literal?",
    "answer": "x := [5]int{1,2,3,4,5}",
    "is_proper": false
  },
  {
    "question": "Can we create global variable with := operator?",
    "answer": "No",
    "is_proper": true
  },
  {
    "question": "How to create map?",
    "answer": "x := map[string]int{\"Philip\":10,\"Michael\": 15}",
    "is_proper": true
  }
];

const router = useRouter();
const stats = computed(() => userResultManager.results.length ? userResultManager.results : mock);
const correctness = computed(
  () => {
    const sum = stats.value.reduce((total, curr) => curr.is_proper ? total+1 : total, 0);
    const length = stats.value.length;
    return (sum/length*100).toFixed(2);
  }
);
const userResultManager = useUserResultStore();

</script>


<template>
  <div class="quiz-results">
    <q-markup-table>
      <thead class="bg-dark text-white shadow-5">
        <tr>
          <th colspan="5">
            <div class="row no-wrap items-start">
              <div class="text-h4 text-white">Results of Goquiz</div>
            </div>
          </th>
        </tr>
        <tr>
          <th class="text-left">Question</th>
          <th class="text-right">Answer</th>
          <th class="text-right">Is proper?</th>
        </tr>
      </thead>
      <tbody class="bg-grey-3">
        <tr 
          v-for="question in stats"
          :key="question.question"
        >
          <td class="text-left">{{ question.question }}</td>
          <td class="text-right">{{ question.answer }}</td>
          <td class="text-right">
            <q-icon 
              size="3em"
              :name="question.is_proper ? 'check' : 'close'" 
              :color="question.is_proper ? 'green' : 'red'" 
            />
          </td>
        </tr>
        <tr >
          <td colspan="2" class="text-right text-bold">Correctness</td>
          <td class="text-right">{{ correctness }}%</td>
        </tr>
      </tbody>
    </q-markup-table>
    <q-btn 
    color="green" 
    text-color="white" 
    label="Back to the dashboard" 
    class="quiz-results__btn" 
    @click.native="router.push('/user-dashboard').then(userResultManager.clearStorage)"
    />
  </div>
</template>

<style lang="scss">
.quiz-results {
  max-width: 800px;
  margin: 0 auto;
  min-height: 100vh;
  display: flex;
  align-items: center;
  flex-direction: column;
  justify-content: center;

  &__btn {
    margin-top: 15px;
  }
}
</style>
