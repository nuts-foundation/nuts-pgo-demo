<template>
  <div class="flex">
    <!-- Navigation -->
    <nav class="main-nav shadow-lg fixed flex flex-col justify-between top-0 h-full w-72 bg-white z-10">
      <div class="px-3 mt-6">
        <router-link
            to="/logout"
            active-class="menu-link-active"
            class="menu-link">
          <div class="w-5 h-5 mr-3">
            <svg viewBox="0 0 24 24" fill="currentColor">
              <g>
                <path d="M0,0h24v24H0V0z" fill="none"/>
              </g>
              <g>
                <path
                    d="M17,8l-1.41,1.41L17.17,11H9v2h8.17l-1.58,1.58L17,16l4-4L17,8z M5,5h7V3H5C3.9,3,3,3.9,3,5v14c0,1.1,0.9,2,2,2h7v-2H5V5z"/>
              </g>
            </svg>
          </div>

          Logout
        </router-link>
      </div>
    </nav>

    <main class="ml-72 mb-14 mt-8 px-12 w-full">
      <!-- Main content -->
      <div class="w-full m-auto max-w-screen-2xl">
        <router-view @statusUpdate="updateStatus"></router-view>
      </div>

      <status-bar :statusMessage="eventMessage"></status-bar>
    </main>
  </div>

</template>

<script>
import StatusBar from '../components/StatusBar.vue'

export default {
  components: { StatusBar },
  data () {
    return {
      eventMessage: ''
    }
  },
  methods: {
    updateStatus (status) {
      this.eventMessage = status
    }
  }
}
</script>

<style>
.menu-link {
  @apply rounded mb-2 py-3 px-3 inline-flex items-center text-gray-500 text-sm;
}

.menu-link:hover {
  @apply text-gray-800 bg-gray-100;
}

.menu-link-active, .menu-link-active:hover {
  @apply bg-gray-200 text-black;
}

.main-title {
  color: #1683af;
}
</style>
