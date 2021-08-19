<template>
<!--  <img :src="dataurl" alt="blockie"/>-->
  <Jazzicon :address="opts.seed" :diameter="opts.size" class="inline-block align-middle leading-none mt-0.5" />
</template>

<script>
import Jazzicon from 'vue3-jazzicon/src/components'
// import blockies from "blockies"
// blockies({})

export default {
  name: "Blockie",
  components: {
    Jazzicon
  },
  props: {
    opts: Object,
  },
  // computed: {
  //   dataurl() {
  //     const icon = blockies(this.opts);
  //     return icon.toDataURL();
  //   }
  // },
}
</script>
