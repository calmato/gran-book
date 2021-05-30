import React from 'react';
import { View, Text, StyleSheet } from 'react-native';
import { MessageForm } from '~/types/forms/index';
import { COLOR } from '~~/constants/theme';

const styles = StyleSheet.create({
  messageStyle: {
    backgroundColor: COLOR.MESSAGE_BACKGROUND,
    padding: 5,
    borderRadius: 5,
    borderBottomRightRadius: 0,
    marginBottom: 5,
  },
});

type Props = {
  userId: string | undefined;
  item: MessageForm;
};

export const MessageItem: React.FC<Props> = ({ item, userId }: Props) => {
  return (
    <View
      style={
        userId == item.userId
          ? { ...styles.messageStyle, alignSelf: 'flex-end' }
          : { ...styles.messageStyle, alignSelf: 'flex-start' }
      }>
      <Text
        style={userId == item.userId ? { color: COLOR.TEXT_WHITE } : { color: COLOR.TEXT_WHITE }}>
        {item.newText}
      </Text>
    </View>
  );
};
