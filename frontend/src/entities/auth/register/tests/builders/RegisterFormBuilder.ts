import { faker } from '@faker-js/faker'
import type { RegisterForm } from '../../schema'

interface Blueprint {
  email: RegisterForm['email']
  password: RegisterForm['password']
  firstname: RegisterForm['firstname']
  lastname: RegisterForm['lastname']
}

class Blueprint implements Blueprint {
  constructor(blueprint: Required<RegisterForm>) {
    this.email = blueprint.email
    this.password = blueprint.password
    this.firstname = blueprint.firstname
    this.lastname = blueprint.lastname
  }

  email: RegisterForm['email']
  password: RegisterForm['password']
  firstname: RegisterForm['firstname']
  lastname: RegisterForm['lastname']
}

export class RegisterFormBuilder {
  private readonly blueprint: Blueprint

  constructor() {
    this.blueprint = new Blueprint({
      email: faker.internet.email(),
      password: faker.internet.password(),
      firstname: faker.person.firstName(),
      lastname: faker.person.lastName(),
    })
  }

  public static create(): RegisterFormBuilder {
    return new RegisterFormBuilder()
  }

  public withEmail(email: RegisterForm['email']): this {
    this.blueprint.email = email
    return this
  }

  public withPassword(password: RegisterForm['password']): this {
    this.blueprint.password = password
    return this
  }

  public withFirstname(firstname: RegisterForm['firstname']): this {
    this.blueprint.firstname = firstname
    return this
  }

  public withLastname(lastname: RegisterForm['lastname']): this {
    this.blueprint.lastname = lastname
    return this
  }

  public get(): RegisterForm {
    return { ...this.blueprint }
  }
}
