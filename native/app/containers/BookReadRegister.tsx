import { RouteProp } from '@react-navigation/native';
import { StackNavigationProp } from '@react-navigation/stack';
import React, { ReactElement } from 'react';
import BookReadRegister from '~/screens/BookReadRegister';
import { useReduxDispatch } from '~/store/modules';
import { registerReadBookImpressionAsync } from '~/store/usecases';
import { ImpressionForm } from '~/types/forms';
import { BookshelfTabStackParamList } from '~/types/navigation';

interface Props {
  route: RouteProp<BookshelfTabStackParamList, 'BookReadRegister'>;
  navigation: StackNavigationProp<BookshelfTabStackParamList, 'BookReadRegister'>;
}

const ConnectedBookReadRegister = function ConnectedBookReadRegister(props: Props): ReactElement {
  const dispatch = useReduxDispatch();

  const actions = React.useMemo(
    () => ({
      registerReadBookImpression(bookId: number, impression: ImpressionForm) {
        return dispatch(registerReadBookImpressionAsync(bookId, impression));
      },
    }),
    [dispatch],
  );

  return <BookReadRegister {...props} actions={actions} />;
};

export default ConnectedBookReadRegister;
