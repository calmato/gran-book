import { ISelect, ITextField } from './utils'

// ---------------------------
// interface - form
// ---------------------------
export interface IAdminListForm {
  limit: number
  offset: number
  order: IAdminListFormOrder
}

export interface IAdminListFormOrder {
  by: string
  desc: boolean
}

export interface IAdminNewForm {
  params: IAdminNewParams
  options: IAdminNewOptions
}

// ---------------------------
// interface - params
// ---------------------------
export interface IAdminNewParams {
  username: string
  email: string
  password: string
  passwordConfirmation: string
  role: number
  lastName: string
  firstName: string
  lastNameKana: string
  firstNameKana: string
}

// ---------------------------
// interface - options
// ---------------------------
export interface IAdminNewOptions {
  username: ITextField
  email: ITextField
  password: ITextField
  passwordConfirmation: ITextField
  role: ISelect
  lastName: ITextField
  firstName: ITextField
  lastNameKana: ITextField
  firstNameKana: ITextField
}

// ---------------------------
// const - Options
// ---------------------------
export const AdminNewOptions: IAdminNewOptions = {
  username: {
    label: '表示名',
    rules: {
      required: true,
      alpha_dash: true,
      max: 32,
    },
  } as ITextField,
  email: {
    label: 'メールアドレス',
    rules: {
      required: true,
      email: true,
    },
  } as ITextField,
  password: {
    label: 'パスワード',
    rules: {
      required: true,
      password: true,
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
  role: {
    label: '権限',
    rules: {
      required: true,
    },
  } as ISelect,
  lastName: {
    label: '姓',
    rules: {
      required: true,
      max: 16,
    },
  } as ITextField,
  firstName: {
    label: '名',
    rules: {
      required: true,
      max: 16,
    },
  } as ITextField,
  lastNameKana: {
    label: '姓 (かな)',
    rules: {
      required: true,
      hiragana: true,
      max: 32,
    },
  } as ITextField,
  firstNameKana: {
    label: '名 (かな)',
    rules: {
      required: true,
      hiragana: true,
      max: 32,
    },
  } as ITextField,
}
