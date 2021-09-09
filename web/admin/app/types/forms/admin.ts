import { IFileInput, ISelect, ITextField } from './utils'

// ---------------------------
// interface - form
// ---------------------------
export interface IAdminSearchForm {
  params: IAdminSearchParams
  options: IAdminSearchOptions
}

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

export interface IAdminEditProfileForm {
  params: IAdminEditProfileParams
  options: IAdminEditProfileOptions
}

export interface IAdminEditContactForm {
  params: IAdminEditContactParams
  options: IAdminEditContactOptions
}

export interface IAdminEditSecurityForm {
  params: IAdminEditSecurityParams
  options: IAdminEditSecurityOptions
}

// ---------------------------
// interface - params
// ---------------------------
export interface IAdminSearchParams {
  value: string
}

export interface IAdminNewParams {
  email: string
  password: string
  passwordConfirmation: string
  role: number
  lastName: string
  firstName: string
  lastNameKana: string
  firstNameKana: string
}

export interface IAdminEditProfileParams {
  lastName: string
  firstName: string
  lastNameKana: string
  firstNameKana: string
  thumbnail: File | null
  thumbnailUrl: string | ArrayBuffer | null
  role: number
}

export interface IAdminEditContactParams {
  email: string
  phoneNumber: string
}

export interface IAdminEditSecurityParams {
  password: string
  passwordConfirmation: string
}

// ---------------------------
// interface - options
// ---------------------------
export interface IAdminSearchOptions {
  value: ITextField
}

export interface IAdminNewOptions {
  email: ITextField
  password: ITextField
  passwordConfirmation: ITextField
  role: ISelect
  lastName: ITextField
  firstName: ITextField
  lastNameKana: ITextField
  firstNameKana: ITextField
}

export interface IAdminEditProfileOptions {
  lastName: ITextField
  firstName: ITextField
  lastNameKana: ITextField
  firstNameKana: ITextField
  thumbnail: IFileInput
  role: ISelect
}

export interface IAdminEditContactOptions {
  email: ITextField
  phoneNumber: ITextField
}

export interface IAdminEditSecurityOptions {
  password: ITextField
  passwordConfirmation: ITextField
}

// ---------------------------
// const - Options
// ---------------------------
export const AdminSearchOptions: IAdminSearchOptions = {
  value: {
    label: '検索値',
    rules: {},
  } as ITextField,
}

export const AdminNewOptions: IAdminNewOptions = {
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

export const AdminEditProfileOptions: IAdminEditProfileOptions = {
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
  thumbnail: {
    label: 'サムネイル',
    rules: {
      image: true,
      size: 3 * 1024, // KB
    },
  } as IFileInput,
  role: {
    label: '権限',
    rules: {
      required: true,
    },
  } as ISelect,
}

export const AdminEditContactOptions: IAdminEditContactOptions = {
  email: {
    label: 'メールアドレス',
    rules: {
      required: true,
      email: true,
    },
  } as ITextField,
  phoneNumber: {
    label: '電話番号',
    rules: {
      required: true,
      max: 16,
    },
  } as ITextField,
}

export const AdminEditSecurityOptions: IAdminEditSecurityOptions = {
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
}
