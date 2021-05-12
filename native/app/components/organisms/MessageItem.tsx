import React from 'react';
import { View, Text } from 'react-native';
import { MessageForm } from '~/types/forms/index';
import { COLOR } from '~~/constants/theme';

type Props ={
  userId: string | undefined;
  item: MessageForm;
};

export const MessageItem: React.FC<Props> = ({ item, userId }: Props) => {
  return (
    <View style={
      userId == item.userId
        ? {
          alignSelf: 'flex-end',
          backgroundColor: COLOR.BACKGROUND_WHITE,
          padding: 5,
          borderRadius: 5,
          borderBottomRightRadius: 0,
          marginBottom: 5
        }
        : {
          alignSelf: 'flex-start',
          backgroundColor: COLOR.BACKGROUND_WHITE,
          padding: 5,
          borderRadius: 5,
          borderBottomLeftRadius: 0,
          marginBottom: 5
        }
    }
    >
      <Text style={userId == item.userId ? { color:COLOR.TEXT_DEFAULT} : {}}>
        {item.newText}
      </Text>
    </View>
  );
};
