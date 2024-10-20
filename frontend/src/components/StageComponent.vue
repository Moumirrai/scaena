<script setup lang="ts">
import { onMounted, useTemplateRef, inject, onUnmounted, ref } from 'vue';
import { PresenterContext } from './PresenterComponent.vue';

const presenterContext = inject<PresenterContext | null>('presenterContext');

const stageRef = useTemplateRef<HTMLDivElement>("stage");

const isFullscreen = ref(false);

onMounted(() => {
    if (!presenterContext?.presenter) return;
    stageRef.value?.appendChild(presenterContext.presenter.stage.finalBuffer);
    document.addEventListener('fullscreenchange', handleFullscreenChange);
    document.addEventListener('keydown', (e) => {
        if (e.key === ' ') {
            onSpacebarPress();
        }
    });
    //presenterContext.presenter.stage.finalBuffer.requestFullscreen();
});

onUnmounted(() => {
    document.removeEventListener('fullscreenchange', handleFullscreenChange);
    document.removeEventListener('keydown', (e) => {
        if (e.key === ' ') {
            onSpacebarPress();
        }
    });
});

//create listener for spacebar
const onSpacebarPress = () => {
    if (isFullscreen.value === true) {
        presenterContext?.presenter?.resume();
    }
}

const handleFullscreenChange = () => {
    isFullscreen.value = !!document.fullscreenElement;
};
</script>

<template>
    <div ref="stage">
        <button @click="presenterContext?.presenter?.stage.finalBuffer.requestFullscreen()">Next</button>
    </div>
    {{ isFullscreen }}

</template>

<style scoped>
/* Scoped styles */
</style>