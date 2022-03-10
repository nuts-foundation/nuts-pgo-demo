<template>
  <div class="flex">
    <!-- Navigation -->
    <nav class="main-nav shadow-lg fixed flex flex-col justify-between top-0 h-full w-72 bg-white z-10">

      <div class="px-3 mt-6">
        <div class="grid grid-cols-1">

          <router-link
              :to="{name: 'pgo.consent'}"
              active-class="menu-link-active"
              class="menu-link">
            <div class="w-5 h-5 mr-3">
              <svg xmlns="http://www.w3.org/2000/svg" enable-background="new 0 0 24 24" viewBox="0 0 24 24"
                   fill="currentColor">
                <g>
                  <rect fill="none" height="24" width="24"/>
                </g>
                <g>
                  <g>
                    <path
                        d="M20,6h-4V4c0-1.1-0.9-2-2-2h-4C8.9,2,8,2.9,8,4v2H4C2.9,6,2,6.9,2,8v12c0,1.1,0.9,2,2,2h16c1.1,0,2-0.9,2-2V8 C22,6.9,21.1,6,20,6z M10,4h4v2h-4V4z M20,20H4V8h16V20z"/>
                    <polygon points="13,10 11,10 11,13 8,13 8,15 11,15 11,18 13,18 13,15 16,15 16,13 13,13"/>
                  </g>
                </g>
              </svg>
            </div>
            Zorgnetwerk&Toestemmingen
          </router-link>
        </div>
      </div>

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
  components: {StatusBar},
  data() {
    return {
      eventMessage: ''
    }
  },
  methods: {
    updateStatus(status) {
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
