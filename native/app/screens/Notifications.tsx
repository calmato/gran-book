import { MaterialIcons } from '@expo/vector-icons';
import SegmentedControl from '@react-native-segmented-control/segmented-control';
import React, { useState } from 'react';
import { View, StyleSheet, Text, FlatList } from 'react-native';
import { Header, Avatar } from 'react-native-elements';
import HeaderText from '~/components/atoms/HeaderText';
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
  userName: {
    color: COLOR.TEXT_DEFAULT,
    width: '50%',
    fontWeight: 'bold',
    fontSize: 24,
  },
  createdAt: {
    textAlign: 'right',
    width: '40%',
    color: COLOR.TEXT_GRAY,
    fontSize: 20,
  },
  latestMessage: {
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

const Notifications = () => {
  const notificationList = ['メッセージ', '取り引き', 'お知らせ'];
  const [selectedIndex, setIndex] = useState<number>(0);
  const testData: RoomInfo[] = [
    {
      rooms: [
        {
          id: '1',
          users: [
            {
              id: '1',
              username: '濵田',
              thumbnailUrl:
                'https://iconbu.com/wp-content/uploads/2021/03/%E3%82%86%E3%82%8B%E3%81%84%E6%81%90%E7%AB%9C%E3%81%AE%E3%83%95%E3%83%AA%E3%83%BC%E3%82%A2%E3%82%A4%E3%82%B3%E3%83%B3.jpg',
            },
          ],
          latestMassage: {
            userId: '1',
            text: 'こんにちは',
            image: '',
            createdAt: '2021/6/24 21:30',
          },
          createdAt: '2021/6/24 21:00',
          updatedAt: '2021/6/24 21:30',
        },
      ],
    },
    {
      rooms: [
        {
          id: '2',
          users: [
            {
              id: '2',
              username: '西川',
              thumbnailUrl:
                'https://iconbu.com/wp-content/uploads/2020/01/%E3%83%9A%E3%83%B3%E3%82%AE%E3%83%B3%E3%81%AE%E3%82%A2%E3%82%A4%E3%82%B3%E3%83%B3.jpg',
            },
          ],
          latestMassage: {
            userId: '2',
            text: 'こんばんは',
            image: '',
            createdAt: '2021/6/23 22:30',
          },
          createdAt: '2021/6/23 20:00',
          updatedAt: '2021/6/23 22:30',
        },
      ],
    },
  ];

  const renderRoom = ({ item }: { item: RoomInfo }) => {
    console.log(item);
    return (
      <View style={styles.roomContainer}>
        <Avatar rounded size="medium" source={{ uri: item.rooms[0].users[0].thumbnailUrl }} />
        <View style={styles.roomInfo}>
          <View style={styles.topInfo}>
            <Text style={styles.userName}>{item.rooms[0].users[0].username}</Text>
            <Text style={styles.createdAt}>{item.rooms[0].updatedAt}</Text>
          </View>
          <View style={styles.bottomInfo}>
            <Text style={styles.latestMessage}>{item.rooms[0].latestMassage.text}</Text>
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
          <FlatList data={testData} renderItem={renderRoom}></FlatList>
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

export default Notifications;
