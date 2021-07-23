import { MaterialIcons } from '@expo/vector-icons';
import React, { useState, useEffect } from 'react';
import { View, StyleSheet, Text } from 'react-native';
import { Header, Avatar, ListItem, Tab } from 'react-native-elements';
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
  roomInfo: RoomInfo;
}

const initialData: RoomInfo = {
  rooms: [
    {
      createdAt: '2021/07/06/23:30',
      id: '1',
      users: [
        {
          id: 'initialize',
          userName: 'ユーザー',
          thumbnailUrl:
            'https://iconbu.com/wp-content/uploads/2021/03/%E3%82%86%E3%82%8B%E3%81%84%E6%81%90%E7%AB%9C%E3%81%AE%E3%83%95%E3%83%AA%E3%83%BC%E3%82%A2%E3%82%A4%E3%82%B3%E3%83%B3.jpg',
        },
      ],
      latestMessage: {
        id: '1',
        userId: '',
        text: 'メッセージがありません',
        image: '',
        createdAt: '2021/07/06/23:30',
      },
      updatedAt: '2021/07/06/23:30',
    },
  ],
};

const NotificationMessage = (props: Props) => {
  const [data, setData] = useState<any>(initialData);
  useEffect(() => {
    const f = async () => {
      const res = await getRoomInfoByUserId();
      console.log(res);
      if (res.rooms[0].latestMessage.text !== '' || res.rooms[0].latestMessage.image !== '') {
        setData(res);
      }
    };
    f();
  }, []);
  const [selectedIndex, setIndex] = useState<number>(0);

  return (
    <View>
      <Header centerComponent={<HeaderText title="通知" />} centerContainerStyle={{ height: 30 }} />
      <Tab value={selectedIndex} onChange={setIndex}>
        <Tab.Item title="メッセージ" titleStyle={styles.selected} />
        <Tab.Item title="取り引き" titleStyle={styles.selected} />
        <Tab.Item title="お知らせ" titleStyle={styles.selected} />
      </Tab>
      {selectedIndex === 0 ? (
        <View style={styles.listContainer}>
          {data.rooms.map((dataInfo) => (
            <ListItem key={dataInfo.id} style={styles.roomContainer}>
              {dataInfo.users[0].thumbnailUrl !== '' ? (
                <Avatar source={{ uri: dataInfo.users[0].thumbnailUrl }} />
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
                  {dataInfo.latestMessage.text !== ''
                    ? dataInfo.latestMessage.text
                    : dataInfo.latestMessage.image}
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
