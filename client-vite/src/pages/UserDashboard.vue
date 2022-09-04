<script setup lang="ts">
import { useUserStore } from '../stores/user.store';
import { onMounted, ref } from 'vue';
import TestBlock from '../components/TestBlock.vue';
import { client, TestPublic } from 'api-client';

const loadingTests = ref(true);
const tests = ref<TestPublic[]>([]);

onMounted(async () => {
  const { me } = useUserStore();
  await me();

  tests.value = await client.tests.getAll();
  loadingTests.value = false;
});
</script>

<template>
  <q-layout view="lHh lpr lFf" container style="height: 100vh">
    <q-header elevated>
      <q-toolbar class="bg-dark">
        <q-avatar square style="align-self: flex-end;">
          <img src="/FilipGOResized.png" style="height: 53px; width: initial;">
        </q-avatar>

        <q-toolbar-title>
          Go<strong>Quiz</strong>
        </q-toolbar-title>
        <q-space></q-space>

      </q-toolbar>
    </q-header>

    <q-page-container>
      <q-page padding>
        <transition name="fade">
          <div v-if="loadingTests" class="text-center">
            <q-spinner-hourglass color="white" size="4em" />
            <q-tooltip>Loading available tests...</q-tooltip>
          </div>
          <div v-else-if="tests.length">
            <TestBlock />
          </div>
          <q-card class="my-card text-black" v-else>
            <q-card-section>
              <h2 class="text-h5">No tests found, sorry.</h2>
              <p>It looks like a server administrator didn't add any test yet.</p>
            </q-card-section>
          </q-card>
        </transition>
      </q-page>
    </q-page-container>
  </q-layout>
</template>

<style lang="scss">
</style>