<template>
  <div class="content flex flex-col min-h-screen bg-main-bg" :class="darkMode?'dark':'light'">
    <Header @theme="changeTheme"></Header>
    <div class="flex-grow p-5">
      <router-view />
    </div>
    <Footer></Footer>
  </div>
</template>

<script>
import Header from "./components/Header.vue";
import Footer from "./components/Footer.vue";
export default {
  name: "app",
  components: {
    Header,
    Footer
  },
  data() {
    return {
      darkMode: false,
    };
  },
  mounted() {},
  methods: {
    changeTheme(darkMode) {
      this.darkMode = darkMode
      console.log('A', darkMode)
    }
  },
  computed: {},
};
</script>
