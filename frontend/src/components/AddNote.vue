<script setup lang="ts">
import { ref, type FormHTMLAttributes } from 'vue';
import type { Note } from '../types/Note';

const formRef = ref<FormHTMLAttributes>();
const form = ref<Note>({
    id: '',
    title: '',
    content: '',
    isDone: false
})

function submit(e: SubmitEvent) {
    e.preventDefault();

    fetch("https://studious-palm-tree-pwv9j54v6rp36655-8080.app.github.dev/notes", {
        method: 'post', body: JSON.stringify(form.value)
    }).then((res) => {
        console.log(res)
    })


}
</script>

<template>
    <form ref="formRef" @submit="submit">
        <label for=" title">Title</label>
        <input type="text" name="title" v-model="form.title">
        <label for="title">Content</label>
        <textarea type="text" name="content" v-model="form.content" />
        <button type="submit">Submit</button>
    </form>

</template>
