<script setup lang="ts">
import { PlayerSettings, Presenter, PresenterInfo, PresenterState, Project, ProjectMetadata, SettingsMetadata, StageSettings, Vector2 } from '@motion-canvas/core';
import { ref, onMounted, watch, defineProps, provide, Ref, useTemplateRef } from 'vue';
import StageComponent from './StageComponent.vue';

enum PresenterStatus {
    LOADING,
    READY,
    ERROR
}

export interface PresenterContext {
    project: Project | null;
    presenter: Presenter | null;
    meta: ProjectMetadata | null;
    settings: PlayerSettings & StageSettings | null;
    state: PresenterState | null;
    status: Ref<PresenterStatus>;
    info: PresenterInfo | null;
}

const presenterContext: PresenterContext = {
    project: null,
    presenter: null,
    meta: null,
    settings: null,
    status: ref<PresenterStatus>(PresenterStatus.LOADING),
    state: null,
    info: null
}
//use provide/inject to pass the presenter context to child components
provide<PresenterContext | null>('presenterContext', presenterContext);

// Define props
const props = defineProps<{
    projectSource: string;
}>();

// Function to call when projectSource changes
async function handleProjectSourceChange(newValue: string, oldValue: string) {
    presenterContext.status.value = PresenterStatus.LOADING;
    await loadProject(newValue);
    if (presenterContext.status.value !== PresenterStatus.LOADING || !presenterContext.project) return;
    presenterContext.meta = presenterContext.project.meta;
    presenterContext.settings = presenterContext.project.meta.getFullRenderingSettings();
};

async function loadProject(projectSource: string) {
    if (!projectSource) {
        presenterContext.status.value = PresenterStatus.ERROR;
        return;
    }
    let project: Project;
    try {
        const promise = import(
        /* webpackIgnore: true */ /* @vite-ignore */ projectSource
        );
        const delay = new Promise(resolve => setTimeout(resolve, 200));
        await Promise.any([delay, promise]);
        project = (await promise).default;
    } catch (e) {
        console.error(e);
        presenterContext.status.value = PresenterStatus.ERROR;
        return;
    }

    project.logger.onLogged.subscribe(log => {
        const { level, message, stack, object, durationMs, ...rest } = log;
        const fn = console[level as 'error'] ?? console.log;
        fn(message, ...[object, durationMs, rest].filter(part => !!part));
        if (stack) {
            fn(stack);
        }
    });

    presenterContext.project = project;

    const presenter = new Presenter(presenterContext.project);
    project.plugins.forEach(plugin => plugin.presenter?.(presenter));
    const meta = project.meta;

    presenterContext.presenter = presenter;
    presenterContext.meta = meta;
    presenterContext.settings = meta.getFullRenderingSettings();

    presenter.onStateChanged.subscribe(state => {
        presenterContext.state = state;
    });

    presenter.onInfoChanged.subscribe(info => {
        presenterContext.info = info;
    });

    presenter.present({
        ...meta.getFullRenderingSettings(),
        name: project.name,
        slide: null,
        size: new Vector2(2560, 1440),
        resolutionScale: 1,
    });

    presenterContext.status.value = PresenterStatus.READY;
}

// Watch for changes in projectSource
watch(() => props.projectSource, handleProjectSourceChange);

onMounted(() => {
    handleProjectSourceChange(props.projectSource, '');
});
</script>

<template>
    <div>
        <div v-if="presenterContext.status.value === PresenterStatus.LOADING">Loading...</div>
        <div v-else-if="presenterContext.status.value === PresenterStatus.ERROR">Error loading project</div>
        <div v-else-if="presenterContext.status.value === PresenterStatus.READY">
            {{ presenterContext.project?.name }}
            <StageComponent />
        </div>
    </div>
</template>

<style scoped>
/* Scoped styles */
</style>