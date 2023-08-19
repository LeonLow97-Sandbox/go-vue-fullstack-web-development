<template>
  <TheHeader />
  <TheBody />
  <TheFooter />
</template>

<script>
import TheHeader from './components/TheHeader.vue';
import TheBody from './components/TheBody.vue';
import TheFooter from './components/TheFooter.vue';

export default {
  name: 'App',
  components: {
    TheHeader,
    TheBody,
    TheFooter,
  },
};
</script>

<style></style>
