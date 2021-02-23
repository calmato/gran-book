// Model
export interface Model {
  readonly id: string;
  readonly token: string;
  readonly email: string;
  readonly emailVerified: boolean;
}

const initialState: Model = {
  id: '',
  token: '',
  email: '',
  emailVerified: false,
};

// Input
export interface AuthValues {
  id: string;
  token: string;
  email?: string;
  emailVerified?: boolean;
}

// Function
export function factory(): Model {
  return initialState;
}

export function setAuth(auth: Model, values: AuthValues): Model {
  return {
    ...auth,
    ...values,
  };
}
