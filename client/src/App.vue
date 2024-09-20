<template>
  <v-app>
    <v-alert>1</v-alert>
    <v-navigation-drawer app>
      <!-- -->
      Menu
    </v-navigation-drawer>
    <v-app-bar app color="primary" dark>
      <v-icon aria-hidden="false">
        mdi-account
      </v-icon>
      <v-spacer></v-spacer>
    </v-app-bar>
    <v-main>
      My collection
      <v-card elevation="2"></v-card>
      <ImageCard></ImageCard>
    </v-main>
  </v-app>
  <v-footer app>
    <!-- -->
    Copyright
  </v-footer>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import ImageCard from "@/components/ImageCard.vue";

export default defineComponent({
  name: 'App',

  components: {
    ImageCard,
  },

  data () {
    return {
      //
    }
  },
})
</script>
