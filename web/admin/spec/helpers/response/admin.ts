export const get = {
  '/v1/admin?limit=20&offset=0': {
    users: [
      {
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
    ],
    limit: 20,
    offset: 0,
    total: 2,
    order: {
      by: 'id',
      direction: 'asc',
    },
  },
  '/v1/admin?limit=20&offset=0&by=email&direction=asc': {
    users: [
      {
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
    ],
    limit: 20,
    offset: 0,
    total: 1,
    order: {
      by: 'email',
      direction: 'asc',
    },
  },
}

export const post = {
  '/v1/admin': {
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
