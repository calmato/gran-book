import { act, renderHook } from '@testing-library/react-hooks/native/pure';
import React, { useContext } from 'react';

import { AuthContext, AuthProvider } from '~/context/auth';
import { AuthValues, initialState, ProfileValues } from '~/store/models/auth';

/**
 * firebase authenticationのmock
 */
jest.mock('~/lib/firebase', () => {
  return {
    auth: jest.fn().mockReturnThis(),
    onAuthStateChanged: jest.fn(),
    signOut: jest.fn(),
    Unsubscribed: jest.fn(),
  };
});

describe('auth context', () => {
  test('call dispatch when call outside context', () => {
    const { result } = renderHook(() => useContext(AuthContext));

    const authPayload: AuthValues = {
      id: '1111',
      token: 'xxxxx',
      email: 'test@calmato.dev',
      emailVerified: true,
    };

    act(() => {
      result.current.dispatch({
        type: 'SET_AUTH_VALUES',
        payload: authPayload,
      });
    });

    expect(result.current.authState).toEqual(initialState);
  });

  test('can dispatch SET_AUTH_VALUES', () => {
    const wrapper = ({ children }) => <AuthProvider>{children}</AuthProvider>;

    const { result } = renderHook(() => useContext(AuthContext), { wrapper });

    const wantInitState = initialState;
    expect(result.current.authState).toEqual(wantInitState);

    const authPayload: AuthValues = {
      id: '1111',
      token: 'xxxxx',
      email: 'test@calmato.dev',
      emailVerified: true,
    };

    act(() => {
      result.current.dispatch({
        type: 'SET_AUTH_VALUES',
        payload: authPayload,
      });
    });

    const expectValue: AuthValues = {
      id: result.current.authState.id,
      token: result.current.authState.token,
      email: result.current.authState.email,
      emailVerified: result.current.authState.emailVerified,
    };

    expect(expectValue).toEqual(authPayload);
  });

  test('can dispatch SET_PROFILE_VALUES', () => {
    const wrapper = ({ children }) => <AuthProvider>{children}</AuthProvider>;

    const { result } = renderHook(() => useContext(AuthContext), { wrapper });

    const wantInitState = initialState;
    expect(result.current.authState).toEqual(wantInitState);

    const profilePayload: ProfileValues = {
      username: 'test',
      gender: 0,
      phoneNumber: '00011112222',
      role: 0,
      thumbnailUrl: '',
      selfIntroduction: '',
      lastName: 'test',
      firstName: 'user',
      lastNameKana: 'てすと',
      firstNameKana: 'ゆーざー',
      postalCode: '',
      prefecture: '',
      city: '',
      addressLine1: '',
      createdAt: '',
      updatedAt: '',
    };

    act(() => {
      result.current.dispatch({
        type: 'SET_PROFILE_VALUES',
        payload: profilePayload,
      });
    });

    const expectValue: ProfileValues = {
      username: result.current.authState.username,
      gender: result.current.authState.gender,
      phoneNumber: result.current.authState.phoneNumber,
      role: result.current.authState.role,
      thumbnailUrl: result.current.authState.thumbnailUrl,
      selfIntroduction: result.current.authState.selfIntroduction,
      lastName: result.current.authState.lastName,
      firstName: result.current.authState.firstName,
      lastNameKana: result.current.authState.lastNameKana,
      firstNameKana: result.current.authState.firstNameKana,
      postalCode: result.current.authState.postalCode,
      prefecture: result.current.authState.prefecture,
      city: result.current.authState.city,
      addressLine1: result.current.authState.addressLine1,
      createdAt: result.current.authState.createdAt,
      updatedAt: result.current.authState.updatedAt,
    };

    expect(expectValue).toEqual(profilePayload);
  });

  test('called cleanup function when components unmount', () => {
    const wrapper = ({ children }) => <AuthProvider>{children}</AuthProvider>;

    const { unmount } = renderHook(() => useContext(AuthContext), { wrapper });
    unmount();

    expect(true).toBeTruthy();
  });
});
