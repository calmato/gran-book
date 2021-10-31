import { RouteProp } from '@react-navigation/native';
import { StackNavigationProp } from '@react-navigation/stack';
import { AxiosResponse } from 'axios';
import React from 'react';
import OtherProfile from '~/screens/OtherProfile';
import { useReduxDispatch } from '~/store/modules';
import { getOtherProfileAsync } from '~/store/usecases/user';
import { UserProfileV1Response } from '~/types/api/user_apiv1_response_pb';
import { BookshelfTabStackParamList } from '~/types/navigation';

interface Props {
  route: RouteProp<BookshelfTabStackParamList, 'OtherProfile'>;
  navigation: StackNavigationProp<BookshelfTabStackParamList, 'OtherProfile'>;
}

export default function ConnectedOwnProfile(props: Props): JSX.Element {
  const dispatch = useReduxDispatch();
  const actions = React.useMemo(
    () => ({
      getOtherProfile(id: string): Promise<AxiosResponse<UserProfileV1Response.AsObject>> {
        return dispatch(getOtherProfileAsync(id));
      },
    }),
    [dispatch],
  );

  return <OtherProfile route={props.route} navigation={props.navigation} actions={actions} />;
}
