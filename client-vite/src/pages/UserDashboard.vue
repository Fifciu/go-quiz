<script setup lang="ts">
import { useUserStore } from '../stores/user.store';
import { onMounted, ref } from 'vue';
import TestBlock from '../components/TestBlock.vue';
import { client, TestPublic } from 'api-client';
import { Cookies } from 'quasar';
import { useRouter } from 'vue-router';

const loadingTests = ref(true);
const tests = ref<TestPublic[]>([]);

const router = useRouter();

onMounted(async () => {
  const { me } = useUserStore();
  await me();

  tests.value = await client.tests.getAll();
  loadingTests.value = false;
});

const logout = () => {
  Cookies.remove(import.meta.env.cookie_token_key);
  router.push({ name: 'home' });
};
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
        <q-btn flat round dense icon="exit_to_app" @click.native="logout" />
      </q-toolbar>
    </q-header>

    <q-page-container>
      <q-page padding>
        <transition name="fade">
          <div v-if="loadingTests" class="text-center">
            <q-spinner-hourglass color="white" size="4em" />
            <q-tooltip>Loading available tests...</q-tooltip>
          </div>
          <div class="q-pa-md row items-start q-gutter-md" v-else-if="tests.length">
            <TestBlock v-for="test in tests" :key="test.id" v-bind="test" />
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