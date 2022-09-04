import { defineStore } from 'pinia';
import { client } from 'api-client';

export const useUserStore = defineStore('user', {
  state: () => ({
    id: -1,
    fullname: '',
    email: ''
  }),
  actions: {
    async me(force = false) {
      if (!force && this.id > 0) {
        return 0;
      }
      const user = await client.users.me();
      this.id = user.id;
      this.fullname = user.fullname;
      this.email = user.email;
    }
  }
});
