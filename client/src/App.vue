<template>
  <v-app>
<!--    <v-alert>1</v-alert>-->

    <v-app-bar app color="primary" dark>
      <v-icon aria-hidden="false">
        mdi-menu
      </v-icon>
<!--          <v-navigation-drawer app>-->
<!--            Menu-->
<!--          </v-navigation-drawer>-->


    </v-app-bar>
    <v-main>
<!--      <v-content>-->
<!--        <v-btn depressed @click="this.info = 'Second'">Normal</v-btn>-->
<!--      </v-content>-->
      <ImageCard />

      <v-btn class="some" color="success">Success</v-btn>
    </v-main>
    <v-footer app>
      Copyright
    </v-footer>
  </v-app>
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
    }
  },
})
</script>
