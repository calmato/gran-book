export interface IAuthResponse {
  id: string
  username: string
  gender: number
  email: string
  phoneNumber: string
  role: number
  thumbnailUrl: string
  selfIntroduction: string
  lastName: string
  firstName: string
  lastNameKana: string
  firstNameKana: string
  postalCode: string
  prefecture: string
  city: string
  addressLine1: string
  addressLine2: string
  createdAt: string
  updatedAt: string
}

export interface IOwnOtherProfileResponse {
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
  products: [],
}
