<template>
  <Loader :is-loading="isPending" />
  <form class="flex flex-col gap-y-4" data-testid="login-form">
    <FormField v-slot="{ componentField }" name="email" :validate-on-blur="!isFieldDirty">
      <FormItem>
        <FormLabel>Email</FormLabel>
        <FormControl>
          <Input
            v-bind="componentField"
            data-testid="login-form-email"
            type="text"
            placeholder="john@doe.com"
          />
        </FormControl>
        <FormMessage />
      </FormItem>
    </FormField>

    <FormField v-slot="{ componentField }" name="password" :validate-on-blur="!isFieldDirty">
      <FormItem>
        <FormLabel>Password</FormLabel>
        <FormControl>
          <Input
            v-bind="componentField"
            data-testid="login-form-password"
            type="password"
            placeholder="password"
          />
        </FormControl>
        <FormMessage />
      </FormItem>
    </FormField>

    <Button class="w-full" type="submit" data-testid="login-form-button" @click="onSubmit">
      Login
    </Button>
  </form>
</template>

<script lang="ts" setup>
import Button from '@/components/ui/button/Button.vue'
import FormLabel from '@/components/ui/form/FormLabel.vue'
import FormControl from '@/components/ui/form/FormControl.vue'
import Input from '@/components/ui/input/Input.vue'
import FormItem from '@/components/ui/form/FormItem.vue'
import FormMessage from '@/components/ui/form/FormMessage.vue'
import Loader from '@/components/ui/loader/Loader.vue'
import { FormField } from '@/components/ui/form'

import { useForm } from 'vee-validate'
import { useRouter } from 'vue-router'
import { useLoginMutation } from '../hooks'
import { useLocalStorage } from '@vueuse/core'
import { toTypedSchema } from '@vee-validate/zod'
import { toast } from 'vue-sonner'
import type { LoginResponse } from '../types'
import { LoginFormSchema } from '../schema'
import type { AxiosResponse } from 'axios'

const { mutate, isPending } = useLoginMutation()
const router = useRouter()

const formSchema = toTypedSchema(LoginFormSchema)

const { isFieldDirty, handleSubmit } = useForm({
  validationSchema: formSchema,
})

const onSubmit = handleSubmit((values) => {
  mutate(values, {
    onError: (error) => {
      toast.error(
        error.message.includes('401') ? 'Invalid email or password' : 'Something went wrong',
      )
    },
    onSuccess: (res: AxiosResponse<LoginResponse>) => {
      toast.success('Login successful')
      useLocalStorage('bearer-token', res.data.token)
      router.push({ name: 'Home' })
    },
  })
})
</script>
