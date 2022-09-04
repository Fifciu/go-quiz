<script setup lang="ts">
defineProps({
  id: Number,
  title: String,
  description: String,
  image_url: String,
  alreadyDone: Boolean
});
</script>

<template>
  <q-card class="my-card">
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
        <q-btn flat color="dark" label="Start" />
        <!-- <q-btn flat color="primary" label="Book" /> -->

        <q-space />

        <q-btn
          v-if="alreadyDone"
          color="grey"
          round
          flat
          dense
          :icon="expanded ? 'keyboard_arrow_up' : 'keyboard_arrow_down'"
          @click="expanded = !expanded"
        />
      </q-card-actions>

      <q-slide-transition>
        <div v-show="expanded">
          <q-separator />
          <q-card-section class="text-subitle2">
            {{ lorem }}
          </q-card-section>
        </div>
      </q-slide-transition>
    </q-card>
</template>

<style lang="scss" scoped>

</style>