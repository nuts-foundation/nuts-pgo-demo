<template>
  <div class="flex justify-center">

    <div class="mt-12 bg-white shadow-sm border rounded-md w-96 p-8 flex flex-col">
      <h1>Create PGO Account</h1>

      <form class="w-full mt-4" @submit.stop.prevent="createAccount">
        <div class="space-y-4">
          <div>
            <label for="username_input" class="block text-sm font-medium text-gray-700">Username</label>
            <input id="username_input"
                   v-model="account.username"
                   type="text"
                   placeholder="Username"
                   class="flex-1 py-2 px-4 block border border-gray-300 rounded-md"
            />
          </div>

          <div>
            <label for="password_input" class="block text-sm font-medium text-gray-700">Password</label>
            <input
                id="password_input"
                v-model="account.password"
                type="password"
                placeholder="Password"
                class="flex-1 py-2 px-4 block border border-gray-300 rounded-md"
            />
          </div>
          <p v-if="!!error" class="p-2 text-center bg-red-100 rounded-md">{{ error }}</p>
          <button class="w-full btn btn-primary">
            Create Account
          </button>
          <button class="w-full btn"
                  @click="$router.push({name: 'login'})">
            I already have an account
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      error: '',
      account: {
        username: '',
        password: ''
      }
    }
  },
  watch: {
    // Remove error when typing
    'account.username'() {
      this.error = ''
    },
    'account.password'() {
      this.error = ''
    }
  },
  methods: {
    redirectAfterLogin() {
      this.$router.push('/pgo/')
    },
    createAccount() {
      console.log(this.account)
      this.$api.createAccount({body: this.account})
          .then(responseData => {
            // after the account is created successfully, login:
            this.$api.createSession({body: this.account})
                .then(responseData => {
                  console.log('success!')
                  localStorage.setItem('session', responseData.token)
                  this.redirectAfterLogin()
                })
          })
          .catch(response => {
            console.error('failure', response)
            if (response === 'invalid account information') {
              this.error = 'account information'
            } else {
              this.error = response.statusText
            }
          })
    }
  },
  mounted() {
  }
}
</script>
