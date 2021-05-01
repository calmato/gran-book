import React, { useState, useMemo } from 'react';
import { View, StyleSheet, SafeAreaView, TextInput, Text, KeyboardAvoidingView } from 'react-native';
import { StatusBar as ExpoStatusBar } from 'expo-status-bar';
import firebase from '~/lib/firebase';
import { getMessageDocRef } from '~/store/usecases/auth';
import { COLOR } from '~~/constants/theme';
import { MaterialIcons, Ionicons, MaterialCommunityIcons } from '@expo/vector-icons';
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
    <View style={styles.container}>
      <ExpoStatusBar style="light" />
      <Header
        leftComponent={
          <MaterialIcons
            name="keyboard-arrow-left"
            size={24}
          />
        }
        centerComponent={{
          text: '濵田',
          style: styles.header
        }}
        centerContainerStyle={{
          height: 50
        }}
      />

      <View style={styles.chatFooter}>
        <MaterialCommunityIcons
          style={styles.inputImage}
          name="image-plus"
          size={32}
        />
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
          style={styles.sendButton}
          name="send"
          size={32}
          disabled={!canSubmit}
          onPress={() => {
            sendMessage(formData.newText);
          }}
        />
      </View>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: COLOR.BACKGROUND_GREY,
  },
  header:{
    color:COLOR.TEXT_TITLE,
    fontSize: 20,
    fontWeight: 'bold',
    flexDirection: 'column',
    justifyContent: 'flex-end',
  },
  inputText: {
    color: COLOR.TEXT_DEFAULT,
    borderWidth: 1,
    borderColor: COLOR.BACKGROUND_GREY,
    backgroundColor: COLOR.BACKGROUND_GREY,
    height: '50%',
    marginLeft: '5%',
    marginRight: '5%',
    flex: 1,
    borderRadius: 10,
    paddingHorizontal: 10,
  },
  inputImage:{
    marginLeft: '3%',
    color: COLOR.TEXT_GRAY
  },
  sendButton: {
    marginRight: '3%',
    color: COLOR.PRIMARY
  },
  chatFooter: {
    flex: 1,
    backgroundColor: COLOR.BACKGROUND_WHITE,
    height: '10%',
    width: '100%',
    flexDirection: 'row',
    alignItems: 'center',
    position: 'absolute',
    bottom: 0
  },
});

export default MessagesScreen;
