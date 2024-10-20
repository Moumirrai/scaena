<script setup lang="ts">
import { onMounted, useTemplateRef, inject } from 'vue';
import { PresenterContext } from './PresenterComponent.vue';

const presenterContext = inject<PresenterContext | null>('presenterContext');

const presenter = presenterContext?.presenter;

const stageRef = useTemplateRef<HTMLDivElement>("stage");

onMounted(() => {
    if (!presenterContext?.presenter) return;
    stageRef.value?.appendChild(presenterContext.presenter.stage.finalBuffer);
    presenterContext.presenter.stage.finalBuffer.requestFullscreen();
});
</script>

<template>
    <button @click="presenter?.resume()">Next</button>
</template>

<style scoped>
/* Scoped styles */
</style>