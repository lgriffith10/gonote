<template>
  <main>
    <Button>Test</Button>
  </main>
</template>

<script setup lang="ts">
import Button from '@/components/ui/button/Button.vue'
import { onMounted } from 'vue'
import { upfetch } from '@/lib/upfetch'

onMounted(async () => {
  await upfetch('/api/test', {})
})
</script>
