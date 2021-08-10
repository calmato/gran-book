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
    color: COLOR.PRIMARY_DARK,
  },
  indicatorStyle: {
    backgroundColor: COLOR.PRIMARY,
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
    height: 100,
    flexDirection: 'column',
  },
  topInfo: {
    width: '100%',
    height: 30,
    flexDirection: 'row',
    marginBottom: 10,
  },
  bottomInfo: {
    width: '100%',
    alignItems: 'center',
    justifyContent: 'space-between',
    height: 30,
    flexDirection: 'row',
  },
  userNameStyle: {
    color: COLOR.TEXT_DEFAULT,
    fontWeight: 'bold',
    width: '30%',
    fontSize: 24,
  },
  createdAtStyle: {
    width: '70%',
    color: COLOR.TEXT_GRAY,
    fontSize: 20,
  },
  latestMessageStyle: {
    width: '80%',
    fontSize: 24,
    color: COLOR.TEXT_GRAY,
  },
  forwardButton: {
    color: 'black',
  },
});
interface Props {
  auth: Auth.Model;
  roomInfo: RoomInfo;
}

const initialData: RoomInfo = {
  rooms: [
    {
      createdAt: '',
      id: '',
      users: [
        {
          id: 'initialize',
          username: '',
          thumbnailUrl: '',
        },
      ],
      latestMessage: {
        id: '',
        userId: '',
        text: '',
        image: '',
        createdAt: '',
      },
      updatedAt: '',
    },
  ],
};

const NotificationMessage = (props: Props) => {
  const [data, setData] = useState<any>(initialData);
  useEffect(() => {
    const f = async () => {
      const res = await getRoomInfoByUserId(props.auth.id);
      console.log(res);
      if (res.rooms[0].latestMessage.text !== '' || res.rooms[0].latestMessage.image !== '') {
        setData(res);
      } else {
        setData(initialData);
      }
    };
    f();
  }, []);
  const [index, setIndex] = useState<number>(0);
  return (
    <View>
      <Header centerComponent={<HeaderText title="通知" />} centerContainerStyle={{ height: 30 }} />
      <Tab value={index} onChange={setIndex} indicatorStyle={styles.indicatorStyle}>
        <Tab.Item title="メッセージ" titleStyle={styles.selected} />
        <Tab.Item title="取り引き" titleStyle={styles.selected} />
        <Tab.Item title="お知らせ" titleStyle={styles.selected} />
      </Tab>
      {index === 0 ? (
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
                <View style={styles.topInfo}>
                  <Text style={styles.userNameStyle}>{dataInfo.users[0].username}</Text>
                  <Text style={styles.createdAtStyle}>{dataInfo.latestMessage.createdAt}</Text>
                </View>
                <View style={styles.bottomInfo}>
                  <Text style={styles.latestMessageStyle}>
                    {dataInfo.latestMessage.text !== ''
                      ? dataInfo.latestMessage.text
                      : dataInfo.latestMessage.image !== ''
                        ? '画像が送信されました'
                        : ''}
                  </Text>
                  <MaterialIcons
                    style={styles.forwardButton}
                    size={32}
                    name="keyboard-arrow-right"
                  />
                </View>
              </ListItem.Content>
            </ListItem>
          ))}
        </View>
      ) : index === 1 ? (
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
