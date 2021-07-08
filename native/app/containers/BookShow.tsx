import { RouteProp } from '@react-navigation/native';
import { StackNavigationProp } from '@react-navigation/stack';
import React, { ReactElement } from 'react';
import BookShow from '~/screens/BookShow';
import { useReduxDispatch } from '~/store/modules';
import { registerOwnBookAsync } from '~/store/usecases';
import { HomeTabStackPramList } from '~/types/navigation';

interface Props {
  route:
    | RouteProp<HomeTabStackPramList, 'SearchResultBookShow'>
    | RouteProp<HomeTabStackPramList, 'BookShow'>;
  navigation:
    | StackNavigationProp<HomeTabStackPramList, 'SearchResultBookShow'>
    | StackNavigationProp<HomeTabStackPramList, 'BookShow'>;
}

export default function ConnectedBookShow(props: Props): ReactElement {
  const dispatch = useReduxDispatch();

  const actions = React.useMemo(
    () => ({
      registerOwnBook(status: string, bookId: number) {
        return dispatch(registerOwnBookAsync(status, bookId));
      },
    }),
    [dispatch],
  );

  return <BookShow route={props.route} navigation={props.navigation} actions={actions} />;
}
