<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <nav>
    <RouterLink to="/">Home</RouterLink>
    <span> | </span>
    <RouterLink to="/about">About</RouterLink>
  </nav>

  <RouterView />
</template>

<style scoped>
nav {
  padding: 1rem;
  text-align: center;
}

nav a {
  color: var(--color-text, #333);
  text-decoration: none;
  font-weight: 500;
}

nav a.router-link-exact-active {
  color: hsla(160, 100%, 37%, 1);
}
</style>
