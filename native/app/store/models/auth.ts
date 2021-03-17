// Model
export interface Model {
  readonly id: string;
  readonly token: string;
  readonly email: string;
  readonly emailVerified: boolean;
  readonly username: string;
  readonly gender: number;
  readonly phoneNumber: string;
  readonly role: number;
  readonly thumbnailUrl: string;
  readonly selfIntroduction: string;
  readonly lastName: string;
  readonly firstName: string;
  readonly lastNameKana: string;
  readonly firstNameKana: string;
  readonly postalCode: string;
  readonly prefecture: string;
  readonly city: string;
  readonly addressLine1: string;
  readonly addressLine2: string;
  readonly createdAt: string;
  readonly updatedAt: string;
  readonly followCount: number;
  readonly followerCount: number;
  readonly reviewCount: number;
  readonly rating: number;
  readonly products: Array<{
      id:number, 
      name: string, 
      thumbnailUrl: string, 
      authors: Array<string>
    }>;
}

export const initialState: Model = {
  id: '',
  token: '',
  email: '',
  emailVerified: false,
  username: '',
  gender: 0,
  phoneNumber: '',
  role: 0,
  thumbnailUrl: '',
  selfIntroduction: '',
  lastName: '',
  firstName: '',
  lastNameKana: '',
  firstNameKana: '',
  postalCode: '',
  prefecture: '',
  city: '',
  addressLine1: '',
  addressLine2: '',
  createdAt: '',
  updatedAt: '',
  followCount: 0,
  followerCount: 0,
  reviewCount: 0,
  rating: 0,
  products: [],
};

// Input
export interface AuthValues {
  id: string;
  token: string;
  email?: string;
  emailVerified?: boolean;
}

export interface ProfileValues {
  username: string;
  gender: number;
  phoneNumber: string;
  role: number;
  thumbnailUrl: string;
  selfIntroduction: string;
  lastName: string;
  firstName: string;
  lastNameKana: string;
  firstNameKana: string;
  postalCode: string;
  prefecture: string;
  city: string;
  addressLine1: string;
  addressLine2?: string;
  createdAt: string;
  updatedAt: string;
}

export interface OwnOtherProfileValues {
  id: string;
  username: string;
  thumbnailUrl: string;
  selfIntroduction: string;
  isFollow: boolean;
  isFollower: boolean;
  followCount: 0,
  followerCount: 0,
  reviewCount: 0,
  rating: 0,
  products: [{
    id: number,
    name: string,
    thumbnailUrl: string,
    authors: [string]
  }],
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

export function setProfile(auth: Model, values: ProfileValues): Model {
  return {
    ...auth,
    ...values,
  };
}

export function setOwnProfile(auth: Model, values: OwnOtherProfileValues): Model {
  return {
    ...auth,
    ...values,
  };
}
