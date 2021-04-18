import React, { useState, useMemo } from 'react';
import { View, StyleSheet, SafeAreaView, TextInput, Text, KeyboardAvoidingView } from 'react-native';
import { StatusBar as ExpoStatusBar } from 'expo-status-bar';
import firebase from '~/lib/firebase';
import { getMessageDocRef } from '~/store/usecases/auth';
import { COLOR } from '~~/constants/theme';
import { Ionicons } from '@expo/vector-icons';
import { MessageForm } from '~/types/forms';
import { Header } from 'react-native-elements';

export const MessagesScreen = () => {
  const [formData, setText] = useState<MessageForm>({
    newText: '',
    createdAt: firebase.firestore.Timestamp.now(),
    userId: '',
  });

  const canSubmit: boolean = useMemo((): boolean => {
    return formData.newText.length > 0;
  }, [formData.newText]);

  const sendMessage = async (value: string) => {
    if (value !== '') {
      const docRef = await getMessageDocRef();
      const newMessage: MessageForm = {
        newText: value,
        createdAt: firebase.firestore.Timestamp.now(),
        userId: '',
      };
      await docRef.add(newMessage);
    }
  };

  return (
    <SafeAreaView style={styles.container}>
      <ExpoStatusBar style="light" />
      <Header>

      </Header>
      <View style={styles.chatFooter}>
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
          name="send"
          size={24}
          color={COLOR.PRIMARY}
          disabled={!canSubmit}
          onPress={() => {
            sendMessage(formData.newText);
          }}
        />
      </View>
    </SafeAreaView>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: COLOR.BACKGROUND_GREY,
    alignItems: 'center',
    justifyContent: 'center',
  },
  inputText: {
    color: COLOR.TEXT_DEFAULT,
    borderWidth: 1,
    borderColor: COLOR.BACKGROUND_GREY,
    backgroundColor: COLOR.BACKGROUND_GREY,
    height: 32,
    flex: 1,
    borderRadius: 5,
    paddingHorizontal: 10,
  },

  inputImage:{

  },

  chatFooter: {
    flex: 1,
    backgroundColor: COLOR.BACKGROUND_WHITE,
    height: 96,
    width: '100%',
    flexDirection: 'row',
    alignItems: 'center',
    position: 'absolute',
    bottom: 0
  },

});

export default MessagesScreen;
