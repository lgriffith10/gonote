<template>
  <div class="flex flex-1 items-center justify-center">
    <Card class="max-w-lg w-full relative overflow-hidden">
      <CardHeader>
        <CardTitle> Login </CardTitle>
        <CardDescription>Login to access to the app</CardDescription>
      </CardHeader>

      <CardContent>
        <LoginForm />
      </CardContent>
    </Card>
  </div>
</template>

<script lang="ts" setup>
import Card from '@/components/ui/card/Card.vue'
import CardContent from '@/components/ui/card/CardContent.vue'
import CardHeader from '@/components/ui/card/CardHeader.vue'
import CardTitle from '@/components/ui/card/CardTitle.vue'
import CardDescription from '@/components/ui/card/CardDescription.vue'
import LoginForm from '@/entities/auth/login/form/LoginForm.vue'
</script>
