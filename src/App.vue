<template>
  <div class="flex flex-col min-h-screen">
    <Header></Header>
    <div class="flex-grow">
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
    return {};
  },
  mounted() {},
  methods: {},
  computed: {},
};
</script>
