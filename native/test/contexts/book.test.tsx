import { renderHook } from '@testing-library/react-hooks/native/pure';
import React, { useContext } from 'react';

import { AuthProvider } from '~/context/auth';
import { BookContext, BookProvider } from '~/context/book';
import { initialState } from '~/store/models/book';

/**
 * firebase authenticationのmock
 */
jest.mock('~/lib/firebase', () => {
  return {
    auth: jest.fn().mockReturnThis(),
    onAuthStateChanged: jest.fn(),
    signOut: jest.fn(),
  };
});

describe('book context', () => {
  const wrapper = ({ children }) => (
    <AuthProvider>
      <BookProvider>{children}</BookProvider>;
    </AuthProvider>
  );

  test('can fetchBooks', () => {
    const { result } = renderHook(() => useContext(BookContext), { wrapper });

    const wantInitState = initialState;
    expect(result.current.bookState).toEqual(wantInitState);
  });
});
