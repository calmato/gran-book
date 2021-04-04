import React, { useState } from 'react';
import { View, StyleSheet, SafeAreaView, TextInput, Button, Alert } from 'react-native';
import { StatusBar as ExpoStatusBar } from 'expo-status-bar';
import firebase from 'firebase';
import { getMessageDocRef } from '~/lib/firebase';
import { Message } from '../types/message/index';
import { COLOR } from '~~/constants/theme';

export const MessagesScreen = () => {
  const [text, setText] = useState<string>('');

  const sendMessage = async (value: string) => {
    if (value != '') {
      const docRef = await getMessageDocRef();
      const newMessage = {
        text: value,
        createdAt: firebase.firestore.Timestamp.now(),
        userId: ''
      } as Message;
      await docRef.set(newMessage);
      setText('');
    } else {
      Alert.alert('エラー', 'メッセージを入力してください！');
    }
  };

  return (
    <SafeAreaView style={styles.container}>
      <ExpoStatusBar style="light" />
      <View style={styles.inputTextContainer}>
        <TextInput
          style={styles.inputText}
          onChangeText={(value) => {
            setText(value);
          }}
          value={text}
          placeholder="メッセージを入力してください"
          placeholderTextColor='#939393'
          autoCapitalize="none"
          autoCorrect={false}
          returnKeyType="done"
        />
        <Button
          title="send"
          onPress={() => {
            sendMessage(text);
          }} />
      </View>
    </SafeAreaView >
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: COLOR.BACKGROUND_GREY,
    alignItems: 'center',
    justifyContent: 'center'
  },
  inputTextContainer: {
    width: '100%',
    flexDirection: 'row',
    alignItems: 'center'
  },
  inputText: {
    color: COLOR.TEXT_DEFAULT,
    borderWidth: 1,
    borderColor: COLOR.BACKGROUND_GREY,
    height: 32,
    flex: 1,
    borderRadius: 5,
    paddingHorizontal: 10

  }
});
