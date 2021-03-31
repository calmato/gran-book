import 'jest';
import { emailValidation } from '~/lib/validation';

describe('emailValidation', () => {
  it('return true when valid email', () => {
    const validEmail = 'sample@test.com';
    const got = emailValidation(validEmail);
    expect(got).toBeTruthy;
  });

  it('return false when invalid email', () => {
    const invalidEmail = 'test';
    const got = emailValidation(invalidEmail);
    expect(got).toBeFalsy;
  });
});
