<script setup lang="ts">
import { useTasksStore } from './stores/tasksStore'
import { computed } from 'vue';

const tasksStore = useTasksStore();
const title = import.meta.env.VITE_TITLE;

const filteredTasks = computed(() => {
  return tasksStore.tasks;
});

</script>

<template>

  <div>
    Hello World! This is the {{ title }} app title.
  </div>

  <div>
    <ul>
      <li v-for="task in filteredTasks" :key="task.id" class="border border-gray-100 p-5 rounded m-5">
        <div :class="task.completed ? 'line-through' : ''">
          {{ task.title }} 
          <input 
            type="checkbox" 
            :checked="task.completed" 
            @change="task.completed = !task.completed" 
          />
        </div>
      </li>
    </ul>
  </div>

</template>

<style scoped>

</style>
