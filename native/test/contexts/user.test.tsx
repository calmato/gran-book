import { act, renderHook } from '@testing-library/react-hooks/native/pure';
import React, { useContext } from 'react';
import { UserContext, UserProvider } from '~/context/user';
import { initialState, UserValues } from '~/store/models/user';

describe('user', () => {
  const wrapper = ({ children }) => <UserProvider>{children}</UserProvider>;

  test('default dispatch', () => {
    const { result } = renderHook(() => useContext(UserContext));

    expect(result.current.dispatch).toBeDefined();
    expect(result.current.dispatch({ type: 'SET_USER', payload: initialState })).toBe(undefined);
  });

  test('can dispatch SET_USER', () => {
    const { result } = renderHook(() => useContext(UserContext), { wrapper });

    const wantState = initialState;
    expect(result.current.userState).toEqual(wantState);

    const userPayload: UserValues = {
      id: '',
      username: 'testuser',
      thumbnailUrl: 'sample',
      selfIntroduction: 'test test',
      isFollow: false,
      isFollower: false,
      followCount: 0,
      followerCount: 0,
      reviewCount: 0,
      rating: 0,
      products: [],
    };

    act(() => {
      result.current.dispatch({
        type: 'SET_USER',
        payload: userPayload,
      });
    });

    expect(result.current.userState).toEqual(userPayload);
  });
});
