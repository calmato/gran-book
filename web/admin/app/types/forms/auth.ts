export interface ISignInForm {
  email: string
  password: string
}

export interface IAuthEditEmailForm {
  email: string
}

export interface IAuthEditEmailValidate {
  email: Object
}

export const AuthEditEmailValidate = {
  email: {
    required: true,
    email: true,
  },
}
