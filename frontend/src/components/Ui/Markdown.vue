<template>
  <div v-html="markdown" class="markdown"></div>
</template>

<script>
import {computed} from 'vue'
import DOMPurify from 'dompurify'
import marked from 'marked'

export default {
  name: 'Markdown',
  props: {
    body: String,
  },
  setup(props) {
    const markdown = computed(() => {
      return marked(DOMPurify.sanitize(props.body))
    })

    return {
      markdown
    }
  }
}
</script>
