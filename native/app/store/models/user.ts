export interface Model {
  readonly id: string;
  readonly username: string;
  readonly thumbnailUrl: string;
  readonly selfIntroduction: string;
  readonly isFollow: boolean;
  readonly isFollower: boolean;
  readonly followCount: number;
  readonly followerCount: number;
  readonly reviewCount: number;
  readonly rating: number;
  readonly products: Array<product>;
}

export interface product {
  readonly id: number;
  readonly name: string;
  readonly thumbnailUrl: string;
  readonly authors: Array<string>;
}

export const initialState: Model = {
  id: '',
  username: '',
  thumbnailUrl: '',
  selfIntroduction: '',
  isFollow: false,
  isFollower: false,
  followCount: 0,
  followerCount: 0,
  reviewCount: 0,
  rating: 0,
  products: [],
};

export interface UserValues {
  id: string;
  username: string;
  thumbnailUrl: string;
  selfIntroduction: string;
  isFollow: boolean;
  isFollower: boolean;
  followCount: number;
  followerCount: number;
  reviewCount: number;
  rating: number;
  products: Array<product>;
}

export function factory(): Model {
  return initialState;
}

export function setUser(user: Model, values: UserValues): Model {
  return {
    ...user,
    ...values,
  };
}
