<template>
  <v-row>
    <v-col
        v-for="n in 12"
        cols="3"
    >
      <v-card>
        <v-img
            :src="`https://i.pravatar.cc/400?img=${n * 5 + 10}`"
            aspect-ratio="1"
            class="bg-grey-lighten-2"
            cover
        >
          <template v-slot:placeholder>
            <v-row
                align="center"
                class="fill-height ma-0"
                justify="center"
            >
              <v-progress-circular
                  color="grey-lighten-5"
                  size="50"
                  width="2"
                  indeterminate
              ></v-progress-circular>
            </v-row>
          </template>
        </v-img>
      </v-card>

    </v-col>
  </v-row>
</template>