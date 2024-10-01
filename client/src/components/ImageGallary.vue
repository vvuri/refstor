<template>
  <v-card class="mx-4 my-2 pa-3">
    <v-row>
      <v-col
          v-for="n in 20"
          cols="6"
          sm="4"
          md="3"
          lg="2"
      >
        <v-card class="rounded-lg">
          <v-img
              :src="`https://i.pravatar.cc/400?img=${n * 3 + 5}`"
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
  </v-card>
</template>