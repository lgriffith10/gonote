<template>
  <Transition>
    <div
      v-if="isLoading"
      class="absolute top-0 left-0 w-full h-full flex justify-center items-center bg-primary/10 z-10"
    >
      <VueSpinnerDots size="40" />
    </div>
  </Transition>
</template>

<script lang="ts" setup>
type Props = {
  isLoading?: boolean
}

const props = defineProps<Props>()

import { VueSpinnerDots } from 'vue3-spinners'
</script>
