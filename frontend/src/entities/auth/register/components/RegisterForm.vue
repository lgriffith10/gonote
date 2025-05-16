<template>
  <Loader v-if="isPending" />
  <form class="flex flex-col gap-y-4" data-testid="register-form" @submit.prevent>
    {{ isError }}
    <FormField v-slot="{ componentField }" name="email" :validate-on-blur="!isFieldDirty">
      <FormItem>
        <FormLabel>Email</FormLabel>
        <FormControl>
          <Input
            v-bind="componentField"
            data-testid="register-form-email"
            type="email"
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
          <Input v-bind="componentField" data-testid="register-form-password" type="password" />
        </FormControl>
        <FormMessage />
      </FormItem>
    </FormField>

    <FormField v-slot="{ componentField }" name="repeatPassword" :validate-on-blur="!isFieldDirty">
      <FormItem>
        <FormLabel>Repeat password</FormLabel>
        <FormControl>
          <Input
            v-bind="componentField"
            data-testid="register-form-repeat-password"
            type="password"
          />
        </FormControl>
        <FormMessage />
      </FormItem>
    </FormField>

    <div class="flex gap-x-4">
      <FormField v-slot="{ componentField }" name="firstname" :validate-on-blur="!isFieldDirty">
        <FormItem>
          <FormLabel>Firstname</FormLabel>
          <FormControl>
            <Input v-bind="componentField" data-testid="register-form-firstname" type="text" />
          </FormControl>
          <FormMessage />
        </FormItem>
      </FormField>

      <FormField v-slot="{ componentField }" name="lastname" :validate-on-blur="!isFieldDirty">
        <FormItem>
          <FormLabel>Lastname</FormLabel>
          <FormControl>
            <Input v-bind="componentField" data-testid="register-form-lastname" type="text" />
          </FormControl>
          <FormMessage />
        </FormItem>
      </FormField>
    </div>
    <Button class="w-full" type="submit" data-testid="register-form-button" @click="onSubmit">
      Register
    </Button>
  </form>
</template>

<script lang="ts" setup>
import Loader from '@/components/ui/loader/Loader.vue'
import { toTypedSchema } from '@vee-validate/zod'
import { RegisterFormSchema } from '../schema'
import { useForm } from 'vee-validate'
import FormItem from '@/components/ui/form/FormItem.vue'
import FormLabel from '@/components/ui/form/FormLabel.vue'
import FormControl from '@/components/ui/form/FormControl.vue'
import { FormField } from '@/components/ui/form'
import Input from '@/components/ui/input/Input.vue'
import Button from '@/components/ui/button/Button.vue'
import FormMessage from '@/components/ui/form/FormMessage.vue'
import { useRegisterMutation } from '../hooks'
import { toast } from 'vue-sonner'

const { mutate, isPending, isError } = useRegisterMutation()

const formSchema = toTypedSchema(RegisterFormSchema)
const { isFieldDirty, handleSubmit } = useForm({
  validationSchema: formSchema,
})

const onSubmit = handleSubmit((values) => {
  mutate(values, {
    onError: (error) => {
      toast.error(error.message)
    },
    onSuccess: () => {
      toast.success('Registration successful! Please check your email to verify your account.')
    },
  })
})
</script>
