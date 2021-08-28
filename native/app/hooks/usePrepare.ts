import { useEffect, useContext, useMemo } from 'react';
import { UiContext as Context } from '~/lib/context';
import * as UiContext from '~/lib/context/ui';
import { useReduxDispatch } from '~/store/modules';
import { authenticationAsync, getAllBookAsync, getAuthAsync } from '~/store/usecases';

/**
 * アプリケーション起動時の準備を全て行うカスタムhooks
 * @returns アプリケーション起動時のコンテキストの状態を返します
 */
export function usePrepare() {
  const uiContext = useContext(UiContext.Context);
  const dispatch = useReduxDispatch();
  const { setApplicationState } = useContext(Context);

  const wait = async (ms: number) => new Promise((resolve) => setTimeout(resolve, ms));
  const waitTime = 1600;

  const actions = useMemo(
    () => ({
      authentication() {
        return dispatch(authenticationAsync());
      },
      getAuth() {
        return dispatch(getAuthAsync());
      },
      getAllBook() {
        return dispatch(getAllBookAsync());
      },
    }),
    [dispatch],
  );

  useEffect(() => {
    const f = async () => {
      actions
        .authentication()
        .then(() => {
          return actions.getAuth();
        })
        .then(() => {
          return actions.getAllBook();
        })
        .then(async () => {
          await wait(waitTime);
          setApplicationState(UiContext.Status.AUTHORIZED);
        })
        .catch(async () => {
          await wait(waitTime);
          setApplicationState(UiContext.Status.UN_AUTHORIZED);
        });
    };
    f();
  }, []);

  return uiContext;
}
