import React, { ReactElement } from 'react';
import { Text, View, StyleSheet } from 'react-native';
import { ListItem, Avatar, Divider, Image, Badge } from 'react-native-elements';
import { COLOR } from '~~/constants/theme';
import { Ionicons } from '@expo/vector-icons'; 

const list = [
  {
    name: 'hamachans',
    avatar_url: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTF0rqSrJnxHQSHdHNXEEUYO4sRucTmGd3BtA&usqp=CAU',
    subtilte: '2020/12/02',
    text: '面白すぎわろた',
    numberOfLikes: 5,
  },
  {
    name: 'Atsuhide',
    avatar_url: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTF0rqSrJnxHQSHdHNXEEUYO4sRucTmGd3BtA&usqp=CAU',
    subtilte: '2020/12/04',
    text: '最高です',
    numberOfLikes: 10000,
  }
];

const bookInfo = {
  title: '何者',
  image_url: 'https://thechara.xsrv.jp/wp-content/uploads/2020/06/200622%E3%80%90NARUTO%E3%80%91KV_02.jpg',
  author: '稲富',
};

const styles = StyleSheet.create({
  badgeStyle: {
    backgroundColor: COLOR.BACKGROUND_WHITE, 
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
      <Badge 
        value={<Text style={{fontSize: 16}}>{`${list.length}件`}</Text>}
        badgeStyle={styles.badgeStyle}
      />
      <View style={styles.bookInfoStyle}>
        <Image
          source={{uri:bookInfo.image_url}}
          style={{width: 50, height: 70,}}
        />
        <View style={{justifyContent: 'space-around', marginStart: 20}}>
          <Text style={{fontSize: 16}}>{bookInfo.title}</Text>
          <Text style={{fontSize: 16, color: COLOR.GREY}}>{bookInfo.author}</Text>
        </View>
      </View>
      <View style={{marginTop: 10}}>
        {
          list.map((l, index) => (
            <View style={{backgroundColor: COLOR.TEXT_WHITE}} key={index}>
              <ListItem key={index} >
                <Avatar source={{uri: l.avatar_url}} rounded/>
                <ListItem.Content>
                  <ListItem.Title>{l.name + 'が感想を投稿しました'}</ListItem.Title>
                  <ListItem.Subtitle>{l.subtilte}</ListItem.Subtitle>
                </ListItem.Content>
              </ListItem>
              <Text style={{fontSize: 16, marginStart: 15, marginEnd:15,}}>{l.text}</Text>
              <View style={{marginStart: 15, flexDirection: 'row',alignItems: 'center'}} >
                <Ionicons name="heart-outline" size={36} color={COLOR.GREY}/>
                <Text style={{marginStart: 10}}>{l.numberOfLikes}</Text>
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
