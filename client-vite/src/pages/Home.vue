<script setup>
import { reactive, ref } from 'vue';
import { useQuasar } from 'quasar'

const user = reactive({
  email: '',
  password: ''
});

const emailValidator = email =>
  String(email)
    .toLowerCase()
    .match(
      /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
    );

const $q = useQuasar()

const name = ref(null)
const age = ref(null)
const accept = ref(false)

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
  name.value = null
  age.value = null
  accept.value = false
}
</script>

<template>
  <div class="home__wrapper">
    <div class="home">
      <div class="home__logo__wrapper">
        <img src="/goicon.png" class="home__logo" />
      </div>
      <div class="home__card">
        <q-card class="text-black">
          <q-form @submit="onSubmit" @reset="onReset" class="q-gutter-md">
            <q-card-section>
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


            </q-card-section>

            <q-card-actions>
              <div>
                <q-btn label="Register" type="submit" color="primary" />
                <q-btn label="Login" type="submit" color="green" class="q-ml-sm" />
                <q-btn label="Reset" type="reset" color="primary" flat class="q-ml-sm" />
              </div>
            </q-card-actions>
          </q-form>

        </q-card>
      </div>
    </div>
  </div>
</template>

<style lang="scss">
.home {
  display: flex;
  align-items: center;
  justify-content: center;
  flex-wrap: wrap;

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
    max-width: 200px;
    margin-bottom: -6px;
  }

  &__logo__wrapper {
    width: 100%;
    text-align: center;
  }

  &__card {
    min-width: 300px;
  }
}
</style>
