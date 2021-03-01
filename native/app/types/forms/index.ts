export type SingUpForm = {
  email: string,
  password: string,
  passwordConfirmation: string,
  username: string,
  agreement: boolean,
}

export type SignInForm = {
  email: string,
  password: string
}

export type PasswordResetForm = {
  email: string
}

export type ProfileEditForm = {
  name: string,
  avatar: string,
  bio: string,
  gender: number,
}

export type RadioGroupForm = {
  label: string,
}

export type PasswordEditForm = {
  password: string,
  passwordConfirmation: string,
}
