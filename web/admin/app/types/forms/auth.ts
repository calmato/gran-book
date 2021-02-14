import { ITextField } from './utils'

// ---------------------------
// interface - form
// ---------------------------
export interface ISignInForm {
  email: string
  password: string
}

export interface IAuthEditEmailForm {
  params: IAuthEditEmailParams
  options: IAuthEditEmailOptions
}

export interface IAuthEditPasswordForm {
  password: string
  passwordConfirmation: string
}

// ---------------------------
// interface - params
// ---------------------------
export interface IAuthEditEmailParams {
  email: string
}

// ---------------------------
// interface - options
// ---------------------------
export interface IAuthEditEmailOptions {
  email: ITextField
}

// ---------------------------
// const - Options
// ---------------------------
export const AuthEditEmailOptions = {
  email: {
    label: 'メールアドレス',
    rules: {
      required: true,
      email: true,
    },
  } as ITextField,
}
