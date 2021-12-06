import { act, renderHook } from '@testing-library/react-hooks/native/pure';
import React, { useContext } from 'react';
import { createApplicationInitialState, Status, UiContext, UiProvider } from '~/context/ui';

describe('ui context', () => {
  const wrapper = ({ children }) => <UiProvider>{children}</UiProvider>;

  test('default function', () => {
    const { result } = renderHook(() => useContext(UiContext));

    act(() => {
      result.current.setApplicationState(Status.FIRST_OPEN);
    });

    expect(result.current.applicationState).toBe(Status.LOADING);
  });

  test('init value', () => {
    const { result } = renderHook(() => useContext(UiContext), { wrapper });

    const wantInitState = createApplicationInitialState();
    expect(result.current.applicationState).toEqual(wantInitState);
  });

  test.each(Object.entries(Status))('update state to %s', (expected) => {
    const { result } = renderHook(() => useContext(UiContext), { wrapper });

    act(() => {
      result.current.setApplicationState(expected as Status);
    });

    expect(result.current.applicationState).toBe(expected);
  });
});
