<template>
  <div class="flex flex-1 items-center justify-center">
    <GnCard title="Register" class="max-w-lg">
      <template #description>
        Already have an account ?
        <RouterLink :to="{ name: 'Login' }">Login</RouterLink>
      </template>

      <RegisterForm />
    </GnCard>
  </div>
</template>

<script lang="ts" setup>
import GnCard from '@/components/ui/card/GnCard.vue'
import RegisterForm from '@/entities/auth/register/components/RegisterForm.vue'
</script>
