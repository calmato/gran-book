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
  params: IAuthEditPasswordParams
  options: IAutheditPasswordOptions
}

// ---------------------------
// interface - params
// ---------------------------
export interface IAuthEditEmailParams {
  email: string
}

export interface IAuthEditPasswordParams {
  password: string
  passwordConfirmation: string
}

// ---------------------------
// interface - options
// ---------------------------
export interface IAuthEditEmailOptions {
  email: ITextField
}

export interface IAutheditPasswordOptions {
  password: ITextField
  passwordConfirmation: ITextField
}

// ---------------------------
// const - Options
// ---------------------------
export const AuthEditEmailOptions: IAuthEditEmailOptions = {
  email: {
    label: 'メールアドレス',
    rules: {
      required: true,
      email: true,
    },
  } as ITextField,
}

export const AuthEditPasswordOptions: IAutheditPasswordOptions = {
  password: {
    label: 'パスワード',
    rules: {
      required: true,
      alpha_dash: true,
      min: 6,
      max: 32,
    },
  } as ITextField,
  passwordConfirmation: {
    label: 'パスワード(確認用)',
    rules: {
      required: true,
      confirmed: 'パスワード',
    },
  } as ITextField,
}
