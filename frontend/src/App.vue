<template>
  <main>
    <Toaster />
    <component :is="layout">
      <RouterView />
    </component>
    <VueQueryDevtools />
  </main>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { RouterView } from 'vue-router'
import { useRoute } from 'vue-router'
import { VueQueryDevtools } from '@tanstack/vue-query-devtools'
import { Toaster } from '@/components/ui/sonner'

const route = useRoute()

const layout = computed(() => {
  return route.meta.layout ?? 'BaseLayout'
})
</script>
