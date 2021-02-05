import React, { ReactElement } from 'react';
import { Text, View, StyleSheet } from 'react-native';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import { ListItem, Avatar, Divider, Image, Badge } from 'react-native-elements'
import { COLOR } from '~~/constants/theme';
import { Ionicons } from '@expo/vector-icons'; 

const list = [
  {
    name: 'hamachans',
    avatar_url: 'https://storage.cloud.google.com/presto-pay-dev.appspot.com/user_thumbnails/80d01b6c-566f-43fa-89e1-7b54cfcb6558',
    subtilte: '2020/12/02',
    text: '面白すぎわろた',
    numberOfLikes: 5,
  },
  {
    name: 'Atsuhide',
    avatar_url: 'https://storage.cloud.google.com/presto-pay-dev.appspot.com/user_thumbnails/80d01b6c-566f-43fa-89e1-7b54cfcb6558',
    subtilte: '2020/12/04',
    text: '最高です',
    numberOfLikes: 10000,
  }
]

const bookInfo = {
  title: '何者',
  image_url: 'https://storage.cloud.google.com/presto-pay-dev.appspot.com/user_thumbnails/80d01b6c-566f-43fa-89e1-7b54cfcb6558',
  author: '稲富',
}

const styles = StyleSheet.create({
  badgeStyle: {
    backgroundColor: COLOR.DARKGREY, 
    alignSelf: 'flex-start', 
    marginStart: 10, 
    marginTop: 10, 
    height: 30, 
    width: 150, 
    borderRadius:75, 
  },
  bookInfoStyle: {
    flexDirection:'row', 
    marginStart: 10, 
    marginTop: 10, 
    backgroundColor: COLOR.TEXT_WHITE,
  },
});

const BookImpression = function BookImpression(): ReactElement {
  return (
    <View>
      <HeaderWithBackButton
        title='感想'
        onPress={() => undefined}
      />
      <Badge 
        value={<Text style={{fontSize: 16}}>{list.length + '件'}</Text>}
        badgeStyle={styles.badgeStyle}
      />
      <View style={styles.bookInfoStyle}>
        <Image
          source={{uri:bookInfo.image_url}}
          style={{width: 50, height: 70,}}
        />
        <View style={{justifyContent: 'space-around', marginStart:20}}>
          <Text style={{fontSize: 16}}>{bookInfo.title}</Text>
          <Text style={{fontSize: 16, color: COLOR.GREY}}>{bookInfo.author}</Text>
        </View>
      </View>
      <View style={{marginTop: 10}}>
        {
          list.map((l, i) => (
            <View style={{backgroundColor: COLOR.TEXT_WHITE}}>
            <ListItem key={i} >
                <Avatar source={{uri: l.avatar_url}}/>
                <ListItem.Content>
                  <ListItem.Title>{l.name + 'が感想を投稿しました'}</ListItem.Title>
                  <ListItem.Subtitle>{l.subtilte}</ListItem.Subtitle>
                </ListItem.Content>
            </ListItem>
            <Text style={{fontSize: 16, marginStart: 15, marginEnd:15,}}>{l.text}</Text>
            <View style={{marginStart: 15, flexDirection: 'row',alignItems: 'center'}} >
              <Ionicons name="heart-outline" size={48} color="black"/>
              <Text style={{marginStart:10}}>{l.numberOfLikes}</Text>
            </View>
            <Divider style={{height: 2}}/>
            </View>
          ))
        }
      </View>
    </View>
  );
};

export default BookImpression;
