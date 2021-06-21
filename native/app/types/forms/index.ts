import firebase from '~/lib/firebase';

export type SingUpForm = {
  email: string;
  password: string;
  passwordConfirmation: string;
  username: string;
  agreement: boolean;
};

export type SignInForm = {
  email: string;
  password: string;
};

export type PasswordResetForm = {
  email: string;
};

export type ProfileEditForm = {
  name: string;
  avatar: string | undefined;
  selfIntroduction: string;
  gender: number;
};

export type RadioGroupForm = {
  label: string;
};

export type PasswordEditForm = {
  password: string;
  passwordConfirmation: string;
};

export type AccountEditForm = {
  firstName: string;
  lastName: string;
  firstNameKana: string;
  lastNameKana: string;
  phoneNumber: string;
  postalCode: string;
  prefecture: string;
  city: string;
  addressLine1: string;
  addressLine2: string;
};

export interface MessageForm {
  text: string;
  createdAt: firebase.firestore.Timestamp;
  _id: string;
  name: string;
}

export interface TransferMessageForm {
  text: string;
  createdAt: any;
  _id: string;
  user: {
    _id: string;
    name: string;
  };
}

export interface ImpressionForm {
  impression: string;
  readOn: string;
}
