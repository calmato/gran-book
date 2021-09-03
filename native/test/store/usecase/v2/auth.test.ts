import MockAdapter from 'axios-mock-adapter';
import internalInstance from '~/lib/axios/internal';
import {
  getProfile,
  signInWithEmailAndPassword,
  signOut,
  signUpWithEmail,
} from '~/store/usecases/v2/auth';
import { AuthV1Response } from '~/types/api/auth_apiv1_response_pb';

import { SingUpForm } from '~/types/forms';

window.addEventListener = jest.fn();

/**
 * firebase authenticationのmock
 */
jest.mock('~/lib/firebase', () => {
  return {
    auth: jest.fn().mockReturnThis(),
    currentUser: jest.fn().mockReturnThis(),
    signInWithEmailAndPassword: jest.fn(() => {
      return Promise.resolve({
        user: {
          uid: '1234567890',
          email: 'test@calmato.dev',
          emailVerified: true,
        },
      });
    }),
    sendEmailVerification: jest.fn().mockResolvedValue,
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

  test('can sing up with email', async () => {
    mockAxios.onPost(`/${API_VERSION}/auth`).reply(201, {});

    const payload: SingUpForm = {
      username: 'test calmato',
      email: 'test@calmato.dev',
      password: '12345678',
      passwordConfirmation: '12345678',
      agreement: true,
    };

    await expect(signUpWithEmail(payload)).resolves.not.toThrow();
  });

  test('cant sign out service', async () => {
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
});
