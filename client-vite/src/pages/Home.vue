<script lang="ts" setup>
import { reactive, ref, watch } from 'vue';
import { useQuasar } from 'quasar'
import { emailValidator } from '../utils/emailValidator';
import TransitionExpand from '../components/TransitionExpand.vue';

const user = reactive({
  email: '',
  fullname: '',
  password: ''
});

const uiForm = ref(null);

const $q = useQuasar()

const REGISTER_TAB = 'register';
const LOGIN_TAB = 'login';
const tab = ref(LOGIN_TAB);

const onSubmit = () => {
  if (accept.value !== true) {
    $q.notify({
      color: 'red-5',
      textColor: 'white',
      icon: 'warning',
      message: 'You need to accept the license and terms first'
    })
  }
  else {
    $q.notify({
      color: 'green-4',
      textColor: 'white',
      icon: 'cloud_done',
      message: 'Submitted'
    })
  }
}

const onReset = () => {
  user.email = '';
  user.fullname = '';
  user.password = '';
  (uiForm.value as any).resetValidation();
};

watch(tab, onReset);

</script>

<template>
  <div class="home__wrapper">
    <div class="home">
      <div class="home__logo__wrapper">
        <div class="home__logo-inner-wrapper">
          <div class="logo__igi__wrapper">
            <img src="/IGIGOResized.png" class="home__logo logo__igi" />
            <img src="/IGIWlos.png" class="home__logo logo__igi__wlos" />
          </div>
          <img src="/FilipGOResized.png" class="home__logo home__logo__fifi" />
        </div>
      </div>
      <div class="home__card">

        <q-card class="text-black">
          <q-tabs v-model="tab" inline-label class="text-white shadow-2 home__tabs">
            <q-tab :name="LOGIN_TAB" icon="mail" label="Login" />
            <q-tab :name="REGISTER_TAB" icon="alarm" label="Register" />
          </q-tabs>

          <q-form ref="uiForm" @submit="onSubmit" @reset="onReset" class="q-gutter-md">
            <q-card-section>
              <TransitionExpand>
                <q-input v-if="tab === REGISTER_TAB" label="Your full name *" hint="Full name" lazy-rules :rules="[
                  val => val && val.length > 0 || 'Please type something',
                  val => val && val.lenght >= 3 || 'Full name has to have at least 3 chars'
                ]" bottom-slots v-model="user.fullname">
                  <template v-slot:prepend>
                    <q-icon name="person" />
                  </template>
                </q-input>
              </TransitionExpand>

              <q-input label="Your email address *" hint="Email address" lazy-rules :rules="[
                val => val && val.length > 0 || 'Please type something',
                val => val && emailValidator(val) || 'Wrong email'
              ]" bottom-slots v-model="user.email">
                <template v-slot:prepend>
                  <q-icon name="contact_mail" />
                </template>
              </q-input>

              <q-input label="Your password *" hint="Password" lazy-rules :rules="[
                val => val && val.length > 0 || 'Please type something',
                val => val && val.length >= 8 || 'Password must has at least 8 characters'
              ]" bottom-slots v-model="user.password">
                <template v-slot:prepend>
                  <q-icon name="lock" />
                </template>
              </q-input>

              <q-btn label="Send" type="submit" color="primary" style="margin-left: 50%;transform:translateX(-50%);" />

            </q-card-section>
          </q-form>
        </q-card>
      </div>
    </div>
  </div>
</template>

<style lang="scss">
// #006cad blue
// #fcc2c0 pink

.home {
  display: flex;
  align-items: center;
  justify-content: center;
  flex-wrap: wrap;

&__logo-inner-wrapper {
    display: flex;

  }

  &__tabs {
    background: #1a1a1a;

    &:hover {
      background: #262626;
    }
  }

  &__wrapper {
    width: 100%;
    height: 100%;
    min-height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  &__subtitle {
    text-align: center;
    margin: 0;
    font-size: 1.25em;
  }

  &__logo {
    max-width: 150px;
    margin-bottom: -46px;

    &__fifi {
      transform: rotate(353deg);
      left: -28px;
      position: relative;
    }
  }

  &__logo__wrapper {
    width: 100%;
    text-align: center;
    display: flex;
    justify-content: center;
  }

  &__card {
    min-width: 300px;
  }
}

.logo {
  &__igi {

    &__wrapper {
      position: relative;
    bottom: -39px;
      left: -19px;

    }

    &__wlos {
      position: absolute;
      z-index:1;
      left: 0;
      bottom: 0;
      max-width:30px;
      pointer-events: none;
    }
  }

}
</style>
