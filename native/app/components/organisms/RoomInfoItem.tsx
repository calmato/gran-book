import { MaterialIcons } from '@expo/vector-icons';
import React, { ReactElement } from 'react';
import { View, StyleSheet, Text } from 'react-native';
import { Avatar } from 'react-native-elements';
import { Auth } from '~/store/models';
import { RoomInfoResponse } from '~/types/response/chat';
import { COLOR } from '~~/constants/theme';

const styles = StyleSheet.create({

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
  updatedAtStyle: {
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
  roomInfo: RoomInfoResponse;

}

const RoomInfoItem = function RoomInfoItem (props:Props):ReactElement  {
  const thumbnailUrl = props.roomInfo.info[0].rooms[0].users[0].thumbnailUrl;
  const userName = props.roomInfo.info[0].rooms[0].users[0].username;
  const updatedAt = props.roomInfo.info[0].rooms[0].updatedAt;
  const latestMessage = props.roomInfo.info[0].rooms[0].latestMessage.text;
  return (
    <View style={styles.roomContainer}>
      <Avatar rounded size="medium" source={{ uri: thumbnailUrl}} />
      <View style={styles.roomInfo}>
        <View style={styles.topInfo}>
          <Text style={styles.userNameStyle}>{userName}</Text>
          <Text style={styles.updatedAtStyle}>{updatedAt}</Text>
        </View>
        <View style={styles.bottomInfo}>
          <Text style={styles.latestMessageStyle}>{latestMessage}</Text>
          <MaterialIcons style={styles.forwardButton} size={32} name="keyboard-arrow-right" />
        </View>
      </View>
    </View>
  );
};
export default RoomInfoItem;
