import { RouteProp } from '@react-navigation/native';
import { StackNavigationProp } from '@react-navigation/stack';
import { AxiosResponse } from 'axios';
import React from 'react';
import { useSelector } from 'react-redux';
import OtherProfile from '~/screens/OtherProfile';
import { Auth, User } from '~/store/models';
import { useReduxDispatch } from '~/store/modules';
import { authSelector, userSelector } from '~/store/selectors';
import { getOtherProfileAsync } from '~/store/usecases/user';
import { BookshelfTabStackParamList } from '~/types/navigation';

interface Props {
  route: RouteProp<BookshelfTabStackParamList, 'OtherProfile'>;
  navigation: StackNavigationProp<BookshelfTabStackParamList, 'OtherProfile'>;
}

export default function ConnectedOwnProfile(props: Props): JSX.Element {
  // MEMO OtherProfile画面で下にスワイプしたら更新できるようにするので、Actionsとして関数渡さなきゃいけない

  return <OtherProfile route={props.route} navigation={props.navigation} actions={undefined} />;
}
