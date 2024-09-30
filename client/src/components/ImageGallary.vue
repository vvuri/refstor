<template>
  <v-row>
    <v-col
        v-for="n in 12"
        cols="3"
    >
      <v-card>
       Img-{{n}}
      </v-card>

    </v-col>
  </v-row>
</template>