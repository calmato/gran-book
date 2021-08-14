export * from './book'

export enum OrderBy {
  ORDER_BY_ASC = 0,
  ORDER_BY_DESC = 1,
}

export enum Gender {
  GENDER_UNKNOWN = 0,
  GENDER_MAN = 1,
  GENDER_WOMAN = 2,
}

export enum Role {
  ROLE_USER = 0,
  ROLE_ADMIN = 1,
  ROLE_DEVELOPER = 2,
  ROLE_OPERATOR = 3,
}

export const LIST_DEFAULT_LIMIT = 100
export const LIST_DEFAULT_OFFSET = 0
export const LIST_DEFAULT_SEARCH_FIELD = ''
export const LIST_DEFAULT_ORDER_FIELD = ''
export const LIST_DEFAULT_ORDER_BY = OrderBy.ORDER_BY_ASC
