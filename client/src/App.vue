<template>
  <v-app>
    <side-bar @selectOption='onSelectOption'/>
    <v-main>
      <add-link />

      <image-gallery :isNext='isNext'/>

    </v-main>
    <v-footer app>
      Copyright
    </v-footer>
  </v-app>
</template>

<script setup>
// import { defineComponent } from 'vue'
import AddLink from "@/components/AddLink.vue";
import SideBar from "@/components/SideBar.vue"
import ImageGallery from "@/components/ImageGallery.vue";
import {ref} from "vue"


const isNext = ref(0)
const onSelectOption = (isDelta) => {
  isNext.value = isNext.value + isDelta
  console.log("===>>>", isDelta, isNext.value)
}

// export default defineComponent({
//   name: 'App',
//
//   components: {
//     AddLink,
//     SideBar,
//     ImageGallery,
//   },
//
//   data () {
//     return {
//     }
//   },
// })
</script>
