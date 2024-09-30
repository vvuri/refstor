<template>
  <v-app>
    <side-bar />
    <v-main>
      <add-link />

<!--      <image-card />-->
      <image-gllary />

<!--      <v-btn class="some" color="success">Success</v-btn>-->
    </v-main>
    <v-footer app>
      Copyright
    </v-footer>
  </v-app>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import ImageCard from "@/components/ImageCard.vue";
import AddLink from "@/components/AddLink.vue";
import SideBar from "@/components/SideBar.vue"
import ImageGllary from "@/components/ImageGallary.vue";


export default defineComponent({
  name: 'App',

  components: {
    ImageCard,
    AddLink,
    SideBar,
    ImageGllary,
  },

  data () {
    return {
    }
  },
})
</script>
