import { ITextField, ITextArea, IFileInput } from './utils'

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
  options: IAuthEditPasswordOptions
}

export interface IAuthEditProfileForm {
  params: IAuthEditProfileParams
  options: IAuthEditProfileOptions
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

export interface IAuthEditProfileParams {
  username: string
  thumbnail?: File
  thumbnailUrl: string
  selfIntroduction: string
  lastName: string
  firstName: string
  lastNameKana: string
  firstNameKana: string
  phoneNumber: string
}

// ---------------------------
// interface - options
// ---------------------------
export interface IAuthEditEmailOptions {
  email: ITextField
}

export interface IAuthEditPasswordOptions {
  password: ITextField
  passwordConfirmation: ITextField
}

export interface IAuthEditProfileOptions {
  username: ITextField
  thumbnail: IFileInput
  selfIntroduction: ITextArea
  lastName: ITextField
  firstName: ITextField
  lastNameKana: ITextField
  firstNameKana: ITextField
  phoneNumber: ITextField
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

export const AuthEditPasswordOptions: IAuthEditPasswordOptions = {
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

export const AuthEditProfileOptions: IAuthEditProfileOptions = {
  username: {
    label: '表示名',
    rules: {
      required: true,
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
  selfIntroduction: {
    label: '自己紹介',
    rules: {
      max: 256,
    },
  } as ITextField,
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
  phoneNumber: {
    label: '電話番号',
    rules: {
      required: true,
      max: 16,
    },
  } as ITextField,
}
