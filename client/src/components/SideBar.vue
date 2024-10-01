<template>
<!--  <v-card>-->
<!--    <v-layout>-->
<!--  <v-app-bar app clipped-left>-->
<!--    <v-app-bar-nav-icon v-if="!permanent" @click.stop="drawer = !drawer" />-->
<!--    <v-toolbar-title v-text="title" />-->
<!--  </v-app-bar>-->
      <v-navigation-drawer
          expand-on-hover
          rail
      >
        <v-list>
          <v-list-item
              prepend-avatar="https://randomuser.me/api/portraits/men/71.jpg"
              subtitle="vvuri...@gmailcom"
              title="Yuri V."
          ></v-list-item>
        </v-list>

        <v-divider></v-divider>

        <v-list density="compact" nav>
          <v-list-item prepend-icon="mdi-plus-circle-outline" title="Add reference" value="reference"></v-list-item>
          <v-list-item prepend-icon="mdi-image-multiple-outline" title="List images" value="shared"></v-list-item>
          <v-list-item prepend-icon="mdi-image-search-outline" title="Search" value="none"></v-list-item>
          <v-list-item prepend-icon="mdi-account-box-outline" title="Contact me" value="none"></v-list-item>
        </v-list>
      </v-navigation-drawer>

<!--      <v-main style="height: 250px"></v-main>-->
<!--    </v-layout>-->
<!--  </v-card>-->
</template>