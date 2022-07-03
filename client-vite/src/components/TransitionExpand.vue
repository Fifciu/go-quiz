<script setup>
const onAfterEnter = (element) => {
  // eslint-disable-next-line no-param-reassign
  element.style.height = `auto`;
};
const onEnter = (element) => {
  const { width } = getComputedStyle(element);
  /* eslint-disable no-param-reassign */
  element.style.width = width;
  element.style.position = `absolute`;
  element.style.visibility = `hidden`;
  element.style.height = `auto`;
  element.style.transition = `none`;
  /* eslint-enable */
  const { height, paddingTop, paddingBottom } = getComputedStyle(element);
  /* eslint-disable no-param-reassign */
  element.style.width = null;
  element.style.position = null;
  element.style.visibility = null;
  element.style.height = 0;
  element.style.paddingTop = 0;
  element.style.paddingBottom = 0;
  /* eslint-enable */
  // Force repaint to make sure the
  // animation is triggered correctly.
  // eslint-disable-next-line no-unused-expressions
  getComputedStyle(element).height;
  element.style.transition = '';
  requestAnimationFrame(() => {
    // eslint-disable-next-line no-param-reassign
    element.style.height = height;
    element.style.paddingTop = paddingTop;
    element.style.paddingBottom = paddingBottom;
    /* eslint-enable */
  });
};
const onLeave = (element) => {
  const { height } = getComputedStyle(element);
  // eslint-disable-next-line no-param-reassign
  element.style.height = height;
  // Force repaint to make sure the
  // animation is triggered correctly.
  // eslint-disable-next-line no-unused-expressions
  getComputedStyle(element).height;
  requestAnimationFrame(() => {
    // eslint-disable-next-line no-param-reassign
    element.style.height = 0;
    element.style.paddingTop = 0;
    element.style.paddingBottom = 0;
    /* eslint-enable */
  });
}
</script>

<template>
  <Transition name="expand" @afterEnter="onAfterEnter" @enter="onEnter" @leave="onLeave">
    <slot></slot>
  </Transition>
</template>

<style scoped>
* {
  will-change: height;
  transform: translateZ(0);
  backface-visibility: hidden;
  perspective: 1000px;
}
</style>

<style>
.expand-enter-active,
.expand-leave-active {
  transition: height .4s ease-in-out;
  overflow: hidden;
}

.expand-enter,
.expand-leave-to {
  height: 0;
}
</style>