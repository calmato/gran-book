import MockAdapter from 'axios-mock-adapter';
import internalInstance from '~/lib/axios/internal';
import { signInWithEmailAndPassword, signOut, signUpWithEmail } from '~/store/usecases/v2/auth';
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

mockAxios.onPost(`/${API_VERSION}/auth`).reply(201, {});

/**
 *
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
});
