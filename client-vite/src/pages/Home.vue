<script setup>
import { computed, reactive } from 'vue';

const user = reactive({
  email: '',
  password: ''
});
const dirty = reactive({
  email: false,
  password: false
});

const emailValidator = email =>
  String(user.email)
    .toLowerCase()
    .match(
      /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
    );
</script>

<template>
  <div class="home__wrapper">
    <div class="home">
      <div class="home__heading">
        <h1 class="home__header my-5 ">GoQuiz</h1>
        <p class="home__subtitle">A simple quiz app created for purpose of the learning!</p>
      </div>
      <q-card class="text-black">
        <q-card-section>

          <q-form @submit="onSubmit" @reset="onReset" class="q-gutter-md">
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
            ]" bottom-slots v-model="user.email">
              <template v-slot:prepend>
                <q-icon name="lock" />
              </template>
            </q-input>


          </q-form>
        </q-card-section>

        <q-card-actions>
          <div>
            <q-btn label="Register" type="submit" color="primary" />
            <q-btn label="Login" type="submit" color="green" class="q-ml-sm" />
            <q-btn label="Reset" type="reset" color="primary" flat class="q-ml-sm" />
          </div>
        </q-card-actions>
      </q-card>
    </div>
  </div>
</template>

<style>
.home__wrapper {
  width: 100%;
  height: 100%;
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
}

.home__header {
  font-weight: bold;
  margin-bottom: 15px;
  text-align: center;
}

.home__subtitle {
  font-size: 1.25em;
}
</style>