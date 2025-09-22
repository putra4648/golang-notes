<script setup lang="ts">
import { ref } from 'vue';

const todos = ref([]);
</script>

<template>
    <h1>my notes</h1>

    <div v-if="!todos.length">Empty</div>
</template>