<template>
  <img :src="dataurl" alt="blockie"/>
</template>

<script>
import blockies from "blockies";
blockies({});
export default {
  name: "Blockie",
  props: {
    opts: Object,
  },
  computed: {
    dataurl() {
      const icon = blockies(this.opts);
      return icon.toDataURL();
    }
  },
}
</script>
