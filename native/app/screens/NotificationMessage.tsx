import { MaterialIcons } from '@expo/vector-icons';
import SegmentedControl from '@react-native-segmented-control/segmented-control';
import React, { useState , useEffect } from 'react';
import { View, StyleSheet, Text } from 'react-native';
import { Header, Avatar, ListItem } from 'react-native-elements';
import HeaderText from '~/components/atoms/HeaderText';
import { Auth } from '~/store/models';
import { getRoomInfoByUserId } from '~/store/usecases/chatServices';
import { RoomInfo } from '~/types/response/chat';
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
  roomInfo?: RoomInfo;
}

const initialData :RoomInfo = {
  rooms: [
    {
      id: '1',
      users: [
        {
          id: 'initialize',
          userName: 'ユーザー',
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
};


const NotificationMessage = (props: Props) => {
  const [data, setData] = useState<any>(initialData);
  useEffect (() => {
    const f = async () => {
      const res = await getRoomInfoByUserId(props.auth.id);
      console.log(res);
      setData(res);
    };
    f();
  } , []
  );
  const notificationList = ['メッセージ', '取り引き', 'お知らせ'];
  const [selectedIndex, setIndex] = useState<number>(0);
  const renderRoom = () => {
    return (
      <View style={styles.roomContainer}>
        <Avatar rounded size="medium" source={{ uri: data.rooms.id}} />
        <View style={styles.roomInfo}>
          <View style={styles.topInfo}>
            <Text style={styles.userNameStyle}>{data.rooms.id}</Text>
            <Text style={styles.createdAtStyle}>{data.rooms.id}</Text>
          </View>
          <View style={styles.bottomInfo}>
            <Text style={styles.latestMessageStyle}>{data.rooms.id}</Text>
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
          {initialData.rooms.map((dataInfo) => (
            <ListItem key={dataInfo.id} style={styles.roomContainer}>
              {dataInfo.users[0].thumbnailUrl !== '' ? (
                <Avatar source={{ uri: dataInfo.users[0].thumbnailUrl}} />
              ) : (
                <Avatar rounded>
                  <MaterialIcons name="person-outline" size={36} color={COLOR.GREY} />
                </Avatar>
              )}
              <ListItem.Content>
                <ListItem.Title style={styles.userNameStyle}>
                  {dataInfo.users[0].userName}
                </ListItem.Title>
                <ListItem.Subtitle>
                  {dataInfo.latestMessage.text}
                </ListItem.Subtitle>
              </ListItem.Content>
            </ListItem>
          ))}
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
