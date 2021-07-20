<template>
  <div class="content flex flex-col min-h-screen bg-main-bg">
    <Header></Header>
    <div class="flex-grow p-5">
      <router-view/>
    </div>
    <Footer></Footer>
    <div id="modal"/>
  </div>
</template>

<script>
import Header from './components/Header.vue';
import Footer from './components/Footer.vue';

export default {
  name: "app",
  components: {
    Header,
    Footer
  },
  data() {
    return {};
  },
  mounted() {
  },
  methods: {},
  computed: {},
};
</script>
