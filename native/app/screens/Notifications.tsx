import React, { useState } from 'react';
import { Header } from 'react-native-elements';
import { View, StyleSheet, Text } from 'react-native';
import HeaderText from '~/components/atoms/HeaderText';
import { COLOR } from '~~/constants/theme';
import SegmentedControl from '@react-native-segmented-control/segmented-control';

const styles = StyleSheet.create({
  selected: {
    color: COLOR.PRIMARY,
  },
});

const Notifications = () => {
  const notificationList = ['メッセージ', '取り引き', 'お知らせ'];
  const [selectedIndex, setIndex] = useState<number>(0);

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
        <View>
          <Text>チャットルーム一覧</Text>
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
