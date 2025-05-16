<template>
  <div class="flex flex-1 items-center justify-center">
    <GnCard title="Login" class="max-w-lg">
      <template #description>
        Don't have an account ?
        <RouterLink :to="{ name: 'Register' }">Register</RouterLink></template
      >

      <LoginForm />
    </GnCard>
  </div>
</template>

<script lang="ts" setup>
import LoginForm from '@/entities/auth/login/components/LoginForm.vue'
import GnCard from '@/components/ui/card/GnCard.vue'
</script>
