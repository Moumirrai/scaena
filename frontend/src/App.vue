<script setup lang="ts">
import { ref, onMounted } from 'vue';
import PresenterComponent from './components/PresenterComponent.vue';

const source = ref('https://motioncanvas.io/examples/presentation.js');
const tempSource = ref(source.value);

onMounted(() => {
  console.log('Mounted App');
});

async function updateSource() {
  console.log('Updating source to:', tempSource.value);
  source.value = tempSource.value;
}
</script>

<template>
  <div>
    <input v-model="tempSource" type="text" placeholder="Enter new source URL" />
    <button @click="updateSource">Update Source</button>
  </div>
  <PresenterComponent :projectSource="source" />
</template>

<style scoped>
</style>