import { IFileInput, ISelect, ITextField } from './utils'

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

export interface IAdminEditForm {
  params: IAdminEditParams
  options: IAdminEditOptions
}

// ---------------------------
// interface - params
// ---------------------------
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

export interface IAdminEditParams {
  email: string
  phoneNumber: string
  role: number
  lastName: string
  firstName: string
  lastNameKana: string
  firstNameKana: string
  thumbnail?: File
  thumbnailUrl: string
}

// ---------------------------
// interface - options
// ---------------------------
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

export interface IAdminEditOptions {
  email: ITextField
  phoneNumber: ITextField
  role: ISelect
  lastName: ITextField
  firstName: ITextField
  lastNameKana: ITextField
  firstNameKana: ITextField
  thumbnail: IFileInput
}

// ---------------------------
// const - Options
// ---------------------------
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

export const AdminEditOptions: IAdminEditOptions = {
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
  thumbnail: {
    label: 'サムネイル',
    rules: {
      image: true,
      size: 3, // 3MB
    },
  } as IFileInput,
}
