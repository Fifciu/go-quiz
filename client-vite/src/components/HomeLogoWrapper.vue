<template>
  <div class="home__logo__wrapper">
    <div class="home__logo-inner-wrapper">
      <div class="logo__igi__wrapper">
        <img src="/IGIGOResized.png" class="home__logo" />
        <img src="/IGIWlos.png" class="logo__igi__wlos" />
      </div>
      <img src="/FilipGOResized.png" class="home__logo home__logo__fifi" />
    </div>
  </div>
</template>

<style lang="scss">
.home__logo-inner-wrapper {
  display: flex;
}

.home__logo {
  max-width: 150px;
  margin-bottom: -46px;

  &__fifi {
    transform: rotate(353deg);
    left: -28px;
    position: relative;
  }
}

.home__logo__wrapper {
  width: 100%;
  text-align: center;
  display: flex;
  justify-content: center;
}

.logo {
  &__igi {
    &__wrapper {
      position: relative;
      bottom: -39px;
      left: -19px;
    }

    &__wlos {
      margin-bottom: -46px;
      position: absolute;
      z-index: 1;
      left: 0;
      bottom: 0;
      max-width: 30px;
      pointer-events: none;
    }
  }
}
</style>