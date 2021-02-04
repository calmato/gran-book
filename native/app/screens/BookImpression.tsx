import React, { ReactElement } from 'react';
import { StyleSheet, View } from 'react-native';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import { ListItem, Avatar } from 'react-native-elements'

const list = [
  {
    name: 'hamachans',
    avatar_url: 'https://storage.cloud.google.com/presto-pay-dev.appspot.com/user_thumbnails/80d01b6c-566f-43fa-89e1-7b54cfcb6558',
    subtilte: '2020/12/02'
  },
  {
    name: 'Atsuhide',
    avatar_url: 'https://storage.cloud.google.com/presto-pay-dev.appspot.com/user_thumbnails/80d01b6c-566f-43fa-89e1-7b54cfcb6558',
    subtilte: '2020/12/04'
  }
]

const BookImpression = function BookImpression(): ReactElement {
  return (
    <View>
      <HeaderWithBackButton
        title='感想'
        onPress={() => undefined}
      />
      <View>
        {
          list.map((l, i) => (
            <ListItem key={i} bottomDivider>
              <Avatar source={{uri: l.avatar_url}} />
              <ListItem.Content>
                <ListItem.Title>{l.name + 'が感想を投稿しました'}</ListItem.Title>
                <ListItem.Subtitle>{l.subtilte}</ListItem.Subtitle>
              </ListItem.Content>
            </ListItem>
          ))
        }
      </View>
    </View>
  );
};

export default BookImpression;
