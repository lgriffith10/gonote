<template>
  <Card
    :class="cn('w-full relative overflow-hidden', props.class)"
    class="w-full relative overflow-hidden"
  >
    <CardHeader>
      <CardTitle>{{ title }}</CardTitle>
      <CardDescription>
        <slot name="description">
          {{ description }}
        </slot>
      </CardDescription>
    </CardHeader>

    <CardContent>
      <slot />
    </CardContent>
  </Card>
</template>

<script lang="ts" setup>
import Card from '@/components/ui/card/Card.vue'
import CardContent from '@/components/ui/card/CardContent.vue'
import CardHeader from '@/components/ui/card/CardHeader.vue'
import CardTitle from '@/components/ui/card/CardTitle.vue'
import CardDescription from '@/components/ui/card/CardDescription.vue'
import type { HTMLAttributes } from 'vue'
import { cn } from '@/lib/utils'

type Props = {
  title?: string
  description?: string
  class?: HTMLAttributes['class']
}

const props = defineProps<Props>()
</script>
