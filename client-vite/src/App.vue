<template>
  <router-view v-slot="{ Component }">
    <transition name="fade">
      <component :is="Component" />
    </transition>
  </router-view>
</template>

<style>
* {
  margin: 0;
  padding: 0;
}

body {
  color: #fff;
  font-family: 'Ubuntu', sans-serif;
  background: rgb(252, 194, 192);
  background: -moz-linear-gradient(90deg, rgba(252, 194, 192, 1) 0%, rgba(0, 108, 173, 1) 100%);
  background: -webkit-linear-gradient(90deg, rgba(252, 194, 192, 1) 0%, rgba(0, 108, 173, 1) 100%);
  background: linear-gradient(90deg, rgba(252, 194, 192, 1) 0%, rgba(0, 108, 173, 1) 100%);
  filter: progid:DXImageTransform.Microsoft.gradient(startColorstr="#fcc2c0", endColorstr="#006cad", GradientType=1);
}

.q-body--prevent-scroll {
  position: static !important;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.5s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>