<template>
  <v-card
      card-background="purple"
      card-color="#FF00FF"
      class="mx-auto my-8"
      max-width="800"
  >
      <v-text-field
          :rules="rules"
          hide-details="auto"
          label="Put URL link image here"
      ></v-text-field>
  </v-card>
</template>

<script>
export default {
  data: () => ({
    rules: [
      value => !!value || 'Required.',
      value => (value && value.length >= 3) || 'Min 3 characters',
    ],
  }),
}
</script>