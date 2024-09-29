<template>
  <v-app>
<!--    <v-app-bar app color="primary" dark>-->
<!--      <v-icon aria-hidden="false">-->
<!--        mdi-menu-->
<!--      </v-icon>-->
<!--          <v-navigation-drawer app>-->
<!--            Menu-->
<!--          </v-navigation-drawer>-->
<!--    </v-app-bar>-->
    <side-bar />
    <v-main>
      <add-link />

      <image-card />

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
import AddLink from "@/components/AddLink.vue";
import SideBar from "@/components/SideBar.vue"

export default defineComponent({
  name: 'App',

  components: {
    ImageCard,
    AddLink,
    SideBar,
  },

  data () {
    return {
    }
  },
})
</script>
