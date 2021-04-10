import React, { useState } from 'react';
import { View, StyleSheet, SafeAreaView, TextInput, Button, Alert } from 'react-native';
import { StatusBar as ExpoStatusBar } from 'expo-status-bar';
import firebase from '~/lib/firebase';
import { getMessageDocRef } from '~/store/usecases/auth';
import { COLOR } from '~~/constants/theme';
import { Ionicons } from '@expo/vector-icons';
import { MessageForm } from '~/types/forms';
import { useMemo } from 'react';

export const MessagesScreen = () => {
  const [formData, setText] = useState<MessageForm>({
    newText: '',
    createdAt: firebase.firestore.Timestamp.now(),
    userId: '',
  });

  const postalCheck: boolean = useMemo((): boolean => {
    return formData.newText.length > 0;
  }, [formData.newText]);

  const sendMessage = async (value: string) => {
    if (value != '') {
      const docRef = await getMessageDocRef();
      const newMessage = {
        newText: value,
        createdAt: firebase.firestore.Timestamp.now(),
        userId: ''
      } as MessageForm;
      await docRef.set(newMessage);
    }
  };

  return (
    <SafeAreaView style={styles.container}>
      <ExpoStatusBar style="light" />
      <View style={styles.inputTextContainer}>
        <TextInput
          style={styles.inputText}
          onChangeText={(value) => {
            setText({ ...formData, newText: value });
          }}
          value={formData.newText}
          placeholder="メッセージを入力してください"
          placeholderTextColor={COLOR.TEXT_GRAY}
          autoCapitalize="none"
          autoCorrect={false}
          returnKeyType="done"
        />
        <Ionicons
          name="send" size={24} color={COLOR.PRIMARY}
          disabled={!postalCheck}
          onPress={() => {
            sendMessage(formData.newText);
          }}
        />
      </View>
    </SafeAreaView >
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: COLOR.BACKGROUND_WHITE,
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
    backgroundColor: COLOR.BACKGROUND_GREY,
    height: 32,
    flex: 1,
    borderRadius: 5,
    paddingHorizontal: 10

  }
});

