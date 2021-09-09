import { AxiosError } from 'axios';
import MockAdapter from 'axios-mock-adapter';
import internalInstance from '~/lib/axios/internal';
import {
  editPassword,
  getProfile,
  sendPasswordResetEmail,
  signInWithEmailAndPassword,
  signOut,
  signUpWithEmail,
} from '~/store/usecases/v2/auth';

import { AuthV1Response } from '~/types/api/auth_apiv1_response_pb';
import { PasswordEditForm, SingUpForm } from '~/types/forms';

window.addEventListener = jest.fn();

/**
 * firebase authenticationのmock
 */
jest.mock('~/lib/firebase', () => {
  return {
    auth: jest.fn().mockReturnThis(),
    currentUser: jest.fn().mockReturnThis(),
    signInWithEmailAndPassword: jest
      .fn()
      .mockImplementationOnce(() => {
        return Promise.resolve({
          user: {
            uid: '1234567890',
            email: 'test@calmato.dev',
            emailVerified: true,
          },
        });
      })
      .mockImplementationOnce(() => {
        return Promise.reject({});
      }),
    sendEmailVerification: jest.fn().mockResolvedValue(undefined),
    sendPasswordResetEmail: jest
      .fn()
      .mockImplementationOnce(() => {
        return Promise.resolve(undefined);
      })
      .mockImplementationOnce(() => {
        return Promise.reject({});
      }),
    signOut: jest.fn(),
  };
});

/**
 * axiosのmock
 */
const mockAxios = new MockAdapter(internalInstance);
const API_VERSION = 'v1';

/**
 * @react-native-community/async-storageのmock
 */
jest.mock('@react-native-community/async-storage', () => {
  return {
    setItem: jest.fn(),
    getItem: jest.fn(),
    removeItem: jest.fn(),
  };
});

describe('auth', () => {
  test('can sign in with email and password', async () => {
    const payload = {
      email: 'test@calmato.dev',
      password: '12345678',
    };

    const user = await signInWithEmailAndPassword(payload);
    expect(user?.email).toBe(payload.email);
  });

  test('signInWithEmailAndPassword return promise reject when sign in failed', () => {
    const payload = {
      email: 'test@calmato.dev',
      password: '12345678',
    };

    expect(signInWithEmailAndPassword(payload)).rejects.not.toThrow();
  });

  test('can sing up with email', async () => {
    mockAxios.onPost(`/${API_VERSION}/auth`).reply(201, {});

    const payload: SingUpForm = {
      username: 'test calmato',
      email: 'test@calmato.dev',
      password: '12345678',
      passwordConfirmation: '12345678',
      agreement: true,
    };

    expect(signUpWithEmail(payload)).resolves.toBe(undefined);
  });

  test('signUpWithEmail return promise reject when sign up failed', async () => {
    mockAxios.onPost(`/${API_VERSION}/auth`).reply(400, {});

    const inValidPayload: SingUpForm = {
      username: 'test calmato',
      email: '',
      password: '12345678',
      passwordConfirmation: '12345678',
      agreement: true,
    };

    return signUpWithEmail(inValidPayload).catch((e) => {
      expect(e.response.status).toBe(400);
    });
  });

  test('can sign out service', async () => {
    await expect(signOut()).resolves.not.toThrow();
  });

  test('can get profile', async () => {
    const mockReturnValue: AuthV1Response.AsObject = {
      id: '1234567890',
      username: 'test calmato',
      email: 'test@calmato.dev',
      gender: '男性',
      phoneNumber: '09011112222',
      thumbnailUrl: '',
      selfIntroduction: '自己紹介です。',
      lastName: 'test',
      firstName: 'user',
      lastNameKana: 'てすと',
      firstNameKana: 'ゆーざー',
      postalCode: '1111111',
      prefecture: '東京都',
      city: '米花町',
      addressLine1: '1-1-1',
      addressLine2: '2F',
      createdAt: '2021/08/31',
      updatedAt: '2021/09/03',
    };

    mockAxios.onGet(`/${API_VERSION}/auth`).reply<AuthV1Response.AsObject>(201, mockReturnValue);

    const profile = await getProfile();
    expect(profile).toBeTruthy();
    expect(profile?.username).toBe(mockReturnValue.username);
    expect(profile?.gender).toBe(mockReturnValue.gender);
    expect(profile?.phoneNumber).toBe(mockReturnValue.phoneNumber);
    expect(profile?.thumbnailUrl).toBe(mockReturnValue.thumbnailUrl);
    expect(profile?.selfIntroduction).toBe(mockReturnValue.selfIntroduction);
    expect(profile?.lastName).toBe(mockReturnValue.lastName);
    expect(profile?.firstName).toBe(mockReturnValue.firstName);
    expect(profile?.lastNameKana).toBe(mockReturnValue.lastNameKana);
    expect(profile?.firstNameKana).toBe(mockReturnValue.firstNameKana);
    expect(profile?.postalCode).toBe(mockReturnValue.postalCode);
    expect(profile?.prefecture).toBe(mockReturnValue.prefecture);
    expect(profile?.city).toBe(mockReturnValue.city);
    expect(profile?.addressLine1).toBe(mockReturnValue.addressLine1);
    expect(profile?.addressLine2).toBe(mockReturnValue.addressLine2);
    expect(profile?.createdAt).toBe(mockReturnValue.createdAt);
    expect(profile?.updatedAt).toBe(mockReturnValue.updatedAt);
  });

  test('getProfile return promise reject when api response status is 500', async () => {
    mockAxios.onGet(`/${API_VERSION}/auth`).reply(500);
    expect(getProfile()).rejects.not.toThrow();
  });

  test('can send password reset email ', () => {
    const payload = { email: 'test@calmato.dev' };
    expect(sendPasswordResetEmail(payload)).resolves.toBe(undefined);
  });

  test('sendPasswordResetEmail return promise reject when invalidPayload', async () => {
    const invalidPayload = { email: 'test2@calmato.dev' };
    await expect(sendPasswordResetEmail(invalidPayload)).rejects.toEqual({});
  });

  test('can edit password', async () => {
    mockAxios.onPatch(`/${API_VERSION}/auth/password`).reply(200);
    const payload: PasswordEditForm = {
      password: '1234567890',
      passwordConfirmation: '1234567890',
    };
    await expect(editPassword(payload)).resolves.toBe(undefined);
  });

  test('editPassword return promise reject when invalidPayload', async () => {
    mockAxios.onPatch(`/${API_VERSION}/auth/password`).reply(400, {});
    const invalidPayload: PasswordEditForm = {
      password: '1234567890',
      passwordConfirmation: '12345678',
    };
    return editPassword(invalidPayload).catch((e: AxiosError) => {
      expect(e.response?.status).toBe(400);
    });
  });
});
