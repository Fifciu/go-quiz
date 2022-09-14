<script lang="ts" setup>
import { onMounted, reactive, ref, watch } from 'vue';
import { useQuasar } from 'quasar'
import { emailValidator } from '../utils/emailValidator';
import TransitionExpand from '../components/TransitionExpand.vue';
import { client, UsersSignUpDraft } from 'api-client';
import { AxiosError } from 'axios';
import HomeLogoWrapper from '../components/HomeLogoWrapper.vue';
import { useRoute, useRouter } from 'vue-router';

const user = reactive<UsersSignUpDraft>({
  email: '',
  fullname: '',
  password: ''
});

const uiForm = ref(null);

const $q = useQuasar();
const route = useRoute();
const router = useRouter();

const REGISTER_TAB = 'register';
const LOGIN_TAB = 'login';
const tab = ref(LOGIN_TAB);

const onSubmit = async () => {
  if (tab.value === REGISTER_TAB) {
    try {
      await client.users.signUp(user);

      $q.notify({
        color: 'green-4',
        textColor: 'white',
        icon: 'cloud_done',
        message: 'Succesfully created account!'
      });
      router.push({ name: 'user-dashboard '});

    } catch (err: AxiosError) {
      $q.notify({
        color: 'red-5',
        textColor: 'white',
        icon: 'warning',
        message: err.response.data.message
      })
    }
  } else if (tab.value === LOGIN_TAB) {
    try {
      await client.users.signIn({
        email: user.email,
        password: user.password
      });
      $q.notify({
        color: 'green-4',
        textColor: 'white',
        icon: 'cloud_done',
        message: 'Succesfully signed in!'
      });
      router.push({ name: 'user-dashboard' });
    } catch (err: AxiosError) {
      console.log(err)
      $q.notify({
        color: 'red-5',
        textColor: 'white',
        icon: 'warning',
        message: err.response.data.message
      })
    }
  }
}

const onReset = () => {
  user.email = '';
  user.fullname = '';
  user.password = '';
  (uiForm.value as any).resetValidation();
};

watch(tab, onReset);

onMounted(async () => {
  if (route.query.error == '401') {
    $q.notify({
      color: 'red-5',
      textColor: 'white',
      icon: 'warning',
      message: 'Unauthorized. Please sign in to access this route.'
    })
  }
});

</script>

<template>
  <div class="home__wrapper">
    <div class="home">
      <HomeLogoWrapper />
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
                  val => val?.length > 0 || 'Please type something',
                  val => val?.length >= 3 || 'Full name has to be at least 3 characters long'
                ]" bottom-slots v-model="user.fullname">
                  <template v-slot:prepend>
                    <q-icon name="person" />
                  </template>
                </q-input>
              </TransitionExpand>

              <q-input label="Your email address *" hint="Email address" lazy-rules :rules="[
                val => val && val.length > 0 || 'Please type something',
                val => val && emailValidator(val) || 'Wrong email address'
              ]" bottom-slots v-model="user.email">
                <template v-slot:prepend>
                  <q-icon name="contact_mail" />
                </template>
              </q-input>

              <q-input type="password" label="Your password *" hint="Password" lazy-rules :rules="[
                val => val && val.length > 0 || 'Please type something',
                val => val && val.length >= 8 || 'Password has to be at least 8 characters long'
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

  &__card {
    min-width: 300px;
  }
}
</style>
