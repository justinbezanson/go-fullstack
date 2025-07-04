<script setup lang="ts">
    import { ref } from 'vue';
    
    const email = ref('');
    const password = ref('');
</script>

<template>
  <h1>Login</h1>
  <div> 
    <input type="text" placeholder="email" name="email" v-model="email" />
    <input type="password" placeholder="password" name="password" v-model="password" />
  </div>  
</template>

<style scoped>

</style>
