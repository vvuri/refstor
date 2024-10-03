<template>
  <v-app>
    <side-bar @selectOption="onSelectOption"/>
    <v-main>
      <add-link />

<!--      <image-card />-->
      <image-gllary :isNext="isNext"/>

<!--      <v-btn class="some" color="success">Success</v-btn>-->
    </v-main>
    <v-footer app>
      Copyright {{ isNext }}
    </v-footer>
  </v-app>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import ImageCard from "@/components/ImageCard.vue";
import AddLink from "@/components/AddLink.vue";
import SideBar from "@/components/SideBar.vue"
import ImageGallary from "@/components/ImageGallary.vue";
import {ref} from "vue"


const isNext = ref(0)
const onSelectOption = (isDelta) => {
  isNext.value = isNext.value + Number(isDelta)
  console.log("===>>>", isDelta, isNext.value)
}

export default defineComponent({
  name: 'App',

  components: {
    ImageCard,
    AddLink,
    SideBar,
    ImageGallary,
  },

  data () {
    return {
    }
  },
})
</script>
