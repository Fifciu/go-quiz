<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { useUserResultStore } from '../stores/userResult.store';

const props = defineProps({
  id: Number,
  title: String,
  description: String,
  image_url: String,
  alreadyDone: Boolean
});

const isModalActive = ref(false);
const userResultManager = useUserResultStore();
const router = useRouter();

const startTest = async () => {
  await userResultManager.start(props.id as number);
  await userResultManager.getQuestions();
  return router.push({
    name: 'quiz',
    params: {
      testId: userResultManager.test_id
    }
  });
};
</script>

<template>
  <q-card class="my-card">
    <q-dialog v-model="isModalActive">
      <q-card class="text-black">
        <q-card-section>
          <div class="text-h6">Are you sure?</div>
        </q-card-section>

        <q-card-section class="q-pt-none">
          You will start a new test session.
        </q-card-section>

        <q-card-actions align="right">
          <q-btn flat label="Exit" color="red" v-close-popup />
          <q-btn flat label="Sure, let's start" color="primary" @click.native="startTest" v-close-popup />
        </q-card-actions>
      </q-card>
    </q-dialog>

      <q-img
        :src="image_url"
      />

      <q-card-section>
        <div class="text-overline text-orange-9">New</div>
        <div class="text-h5 q-mt-sm q-mb-xs text-black">{{ title }}</div>
        <div class="text-caption text-grey">
          {{ description }}
        </div>
      </q-card-section>

      <q-card-actions>
        <q-btn flat color="dark" label="Start" @click.native="() => isModalActive = true" />
        <!-- <q-btn flat color="primary" label="Book" /> -->

        <q-space />

        <!-- <q-btn
          v-if="alreadyDone"
          color="grey"
          round
          flat
          dense
          :icon="expanded ? 'keyboard_arrow_up' : 'keyboard_arrow_down'"
          @click="expanded = !expanded"
        /> -->
      </q-card-actions>

      <!-- <q-slide-transition>
        <div v-show="expanded">
          <q-separator />
          <q-card-section class="text-subitle2">
            {{ lorem }}
          </q-card-section>
        </div>
      </q-slide-transition> -->
    </q-card>
</template>

<style lang="scss" scoped>

</style>
