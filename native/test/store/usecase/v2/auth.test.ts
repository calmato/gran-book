import { signInWithEmailAndPassword } from '~/store/usecases/v2/auth';

window.addEventListener = jest.fn();

jest.mock('~/lib/firebase', () => {
  return {
    auth: jest.fn().mockReturnThis(),
    signInWithEmailAndPassword: jest.fn(() => {
      return Promise.resolve({
        user: {
          uid: '1234567890',
          email: 'test@calmato.dev',
          emailVerified: true,
        },
      });
    }),
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
});
