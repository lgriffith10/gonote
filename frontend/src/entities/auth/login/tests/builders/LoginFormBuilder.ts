import type { LoginForm } from '../../schema'
import { faker } from '@faker-js/faker'

interface Blueprint {
  email: LoginForm['email']
  password: LoginForm['password']
}

class Blueprint implements Blueprint {
  constructor(blueprint: Required<LoginForm>) {
    this.email = blueprint.email
    this.password = blueprint.password
  }

  email: LoginForm['email']
  password: LoginForm['password']
}

export class LoginFormBuilder {
  private readonly blueprint: Blueprint

  constructor() {
    this.blueprint = new Blueprint({
      email: faker.internet.email(),
      password: faker.internet.password(),
    })
  }

  public static create(): LoginFormBuilder {
    return new LoginFormBuilder()
  }

  public withEmail(email: LoginForm['email']): this {
    this.blueprint.email = email
    return this
  }

  public withPassword(password: LoginForm['password']): this {
    this.blueprint.password = password
    return this
  }

  public get(): LoginForm {
    return { ...this.blueprint }
  }
}
