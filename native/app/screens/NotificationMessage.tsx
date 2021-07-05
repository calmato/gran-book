import { MaterialIcons } from '@expo/vector-icons';
import SegmentedControl from '@react-native-segmented-control/segmented-control';
import React, { useState, useEffect } from 'react';
import { View, StyleSheet, Text, FlatList } from 'react-native';
import { Header, Avatar } from 'react-native-elements';
import HeaderText from '~/components/atoms/HeaderText';
import { Auth } from '~/store/models';
import { getRoomInfoByUserId } from '~/store/usecases/chatServices';
import { RoomInfoResponse } from '~/types/response/chat';
import { COLOR } from '~~/constants/theme';

const styles = StyleSheet.create({
  selected: {
    color: COLOR.PRIMARY,
  },
  listContainer: {
    backgroundColor: COLOR.BACKGROUND_GREY,
    marginVertical: 5,
    height: '100%',
  },
  roomContainer: {
    paddingVertical: 10,
    paddingHorizontal: 10,
    backgroundColor: COLOR.BACKGROUND_WHITE,
    height: 70,
    alignItems: 'center',
    flexDirection: 'row',
  },
  roomInfo: {
    paddingLeft: 10,
    paddingRight: 20,
    width: '100%',
    flexDirection: 'column',
  },
  topInfo: {
    width: '100%',
    height: 30,
    alignItems: 'center',
    flexDirection: 'row',
  },
  bottomInfo: {
    width: '100%',
    alignItems: 'center',
    height: 40,
    flexDirection: 'row',
  },
  userNameStyle: {
    color: COLOR.TEXT_DEFAULT,
    width: '50%',
    fontWeight: 'bold',
    fontSize: 24,
  },
  createdAtStyle: {
    textAlign: 'right',
    width: '40%',
    color: COLOR.TEXT_GRAY,
    fontSize: 20,
  },
  latestMessageStyle: {
    fontSize: 24,
    width: '70%',
    color: COLOR.TEXT_GRAY,
  },
  forwardButton: {
    textAlign: 'right',
    color: 'black',
    width: '20%',
  },
});
interface Props {
  auth: Auth.Model;
  roomInfo?: RoomInfoResponse;
}

const intitialData :RoomInfoResponse = {
  limit: 1,
  offset: 1,
  total: 1,
  rooms:[{
    rooms: [
      {
        id: '',
        users: [
          {
            id: 'initialize',
            username: 'ユーザー',
            thumbnailUrl: 'https://iconbu.com/wp-content/uploads/2021/03/%E3%82%86%E3%82%8B%E3%81%84%E6%81%90%E7%AB%9C%E3%81%AE%E3%83%95%E3%83%AA%E3%83%BC%E3%82%A2%E3%82%A4%E3%82%B3%E3%83%B3.jpg',
          }
        ],
        latestMessage: {
          userId: '',
          text: 'メッセージがありません',
          image: '',
          createdAt: '2021/07/06/23:30'
        },
        createdAt: '2021/07/06/23:30',
        updatedAt: '2021/07/06/23:30',
      }
    ]
  }
  ]
};


const NotificationMessage = (props: Props) => {
  useEffect(() => {
    getRoomInfoByUserId(props.auth.id);
  }, []);;
  const roomItem = props.roomInfo?.rooms || intitialData ;
  const thumbnailUrl = roomItem[0].rooms[0].users[0].thumbnailUrl;
  const userName = roomItem[0].rooms[0].users[0].username;
  const createdAt = roomItem[0].rooms[0].latestMessage.createdAt;
  const latestMessage = roomItem[0].rooms[0].latestMessage.text;
  const notificationList = ['メッセージ', '取り引き', 'お知らせ'];
  const [selectedIndex, setIndex] = useState<number>(0);
  const renderRoom = () => {
    getRoomInfoByUserId(props.auth.id);
    return (
      <View style={styles.roomContainer}>
        <Avatar rounded size="medium" source={{ uri: thumbnailUrl}} />
        <View style={styles.roomInfo}>
          <View style={styles.topInfo}>
            <Text style={styles.userNameStyle}>{userName}</Text>
            <Text style={styles.createdAtStyle}>{createdAt}</Text>
          </View>
          <View style={styles.bottomInfo}>
            <Text style={styles.latestMessageStyle}>{latestMessage}</Text>
            <MaterialIcons style={styles.forwardButton} size={32} name="keyboard-arrow-right" />
          </View>
        </View>
      </View>
    );
  };

  return (
    <View>
      <Header centerComponent={<HeaderText title="通知" />} centerContainerStyle={{ height: 30 }} />
      <SegmentedControl
        activeFontStyle={styles.selected}
        values={notificationList}
        backgroundColor={COLOR.BACKGROUND_WHITE}
        selectedIndex={selectedIndex}
        onValueChange={(event) => setIndex(notificationList.indexOf(event))}
      />
      {selectedIndex === 0 ? (
        <View style={styles.listContainer}>
          <FlatList data={roomItem} renderItem={renderRoom}></FlatList>
        </View>
      ) : selectedIndex === 1 ? (
        <View>
          <Text>取引き</Text>
        </View>
      ) : (
        <View>
          <Text>お知らせ</Text>
        </View>
      )}
    </View>
  );
};

export default NotificationMessage;
