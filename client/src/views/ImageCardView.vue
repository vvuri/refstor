<template>
  <hello-world />
</template>

<script lang="ts">
import { defineComponent } from 'vue';

// Components
import ImageCard from '../components/ImageCard.vue';

export default defineComponent({
  name: 'ImageCard',

  components: {
    ImageCard,
  },
});
</script>
