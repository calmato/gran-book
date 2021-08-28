import { AuthThumbnailV1Response, AuthV1Response } from '~/types/api/auth_apiv1_response_pb'

export const get: { [key: string]: AuthV1Response.AsObject } = {
  '/v1/auth': {
    id: '00000000-0000-0000-00000000',
    username: 'test-user',
    email: 'test@calmato.com',
    phoneNumber: '000-0000-0000',
    role: 0,
    thumbnailUrl: 'https://calmato.com/images/01',
    selfIntroduction: 'よろしくお願いします',
    lastName: 'テスト',
    firstName: 'ユーザ',
    lastNameKana: 'てすと',
    firstNameKana: 'ゆーざ',
    createdAt: '2021-01-01 00:00:00',
    updatedAt: '2021-01-01 00:00:00',
  },
}

export const post: { [key: string]: AuthThumbnailV1Response.AsObject } = {
  '/v1/auth/thumbnail': {
    thumbnailUrl: 'https://calmato.com/images/01',
  },
}

export const patch: { [key: string]: AuthV1Response.AsObject } = {
  '/v1/auth/email': {
    id: '00000000-0000-0000-00000000',
    username: 'test-user',
    email: 'test@calmato.com',
    phoneNumber: '000-0000-0000',
    role: 0,
    thumbnailUrl: 'https://calmato.com/images/01',
    selfIntroduction: 'よろしくお願いします',
    lastName: 'テスト',
    firstName: 'ユーザ',
    lastNameKana: 'てすと',
    firstNameKana: 'ゆーざ',
    createdAt: '2021-01-01 00:00:00',
    updatedAt: '2021-01-01 00:00:00',
  },
  '/v1/auth/password': {
    id: '00000000-0000-0000-00000000',
    username: 'test-user',
    email: 'test@calmato.com',
    phoneNumber: '000-0000-0000',
    role: 0,
    thumbnailUrl: 'https://calmato.com/images/01',
    selfIntroduction: 'よろしくお願いします',
    lastName: 'テスト',
    firstName: 'ユーザ',
    lastNameKana: 'てすと',
    firstNameKana: 'ゆーざ',
    createdAt: '2021-01-01 00:00:00',
    updatedAt: '2021-01-01 00:00:00',
  },
  '/v1/auth/profile': {
    id: '00000000-0000-0000-00000000',
    username: 'test-user',
    email: 'test@calmato.com',
    phoneNumber: '000-0000-0000',
    role: 0,
    thumbnailUrl: 'https://calmato.com/images/01',
    selfIntroduction: 'よろしくお願いします',
    lastName: 'テスト',
    firstName: 'ユーザ',
    lastNameKana: 'てすと',
    firstNameKana: 'ゆーざ',
    createdAt: '2021-01-01 00:00:00',
    updatedAt: '2021-01-01 00:00:00',
  },
}
