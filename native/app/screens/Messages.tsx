import React, { useState, useMemo, useEffect, useCallback, } from 'react';
import { View, StyleSheet, SafeAreaView, TextInput, Text, KeyboardAvoidingView, FlatList, } from 'react-native';
import { StatusBar as ExpoStatusBar } from 'expo-status-bar';
import firebase from '~/lib/firebase';
import { getMessageDocRef } from '~/store/usecases/auth';
import { COLOR } from '~~/constants/theme';
import { MaterialIcons, Ionicons, MaterialCommunityIcons } from '@expo/vector-icons';
import { MessageForm, transferMessageForm } from '~/types/forms';
import { Header } from 'react-native-elements';
import { MessageItem } from '~/components/organisms/MessageItem';
import { GiftedChat } from 'react-native-gifted-chat';

export const MessagesScreen = () => {
  const [textData, setText] = useState ({
    text: '',
    createdAt: firebase.firestore.Timestamp.now(),
    _id: '',
  });
  const [messages, setMessages] = useState([]);
  const docRef = getMessageDocRef();

  const onSend = (messages = []) => {
    messages.forEach(async (message) => {
      (await docRef).add(message);
    });
  };
  const getMessage = async () => {
    const messages = [] as transferMessageForm[];
    await firebase.firestore().collection('messages').orderBy('createdAt')
      .onSnapshot((snapshot) => {
        snapshot.docChanges().forEach((change) => {
          if (change.type === 'added') {
            const messageInfo = {
              _id: change.doc.data()._id,
              createdAt: change.doc.data().createdAt.toDate(),
              text: change.doc.data().text,
              user: change.doc.data().user,
            };
            messages.unshift(messageInfo as transferMessageForm);
          }
          if (messages.length != 0) {
            setMessages(messages);
          }
        });
      });
  };
  useEffect(() => {
    getMessage();
  }, []);
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
          text: 'Name',
          style: styles.header
        }}
        centerContainerStyle={{
          height: 40
        }}
      />
      <GiftedChat
        messages= {messages}
        onSend={messages => onSend(messages)}
        user={{
          _id: 1,
          name: 'Kaito'
        }}
      />

    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: COLOR.BACKGROUND_GREY,
  },
  header: {
    color:COLOR.TEXT_TITLE,
    fontSize: 20,
    fontWeight: 'bold',
    flexDirection: 'column',
    justifyContent: 'flex-end',
  },
  messagesContainer: {
    marginBottom: 'auto',
    width: 'auto',
    padding: 10
  },

  inputText: {
    color: COLOR.TEXT_DEFAULT,
    borderWidth: 1,
    borderColor: COLOR.BACKGROUND_GREY,
    backgroundColor: COLOR.BACKGROUND_GREY,
    height: 'auto',
    minHeight: '20%',
    width: '70%',
    marginVertical: '5%',
    marginHorizontal: '5%',
    borderRadius: 10,
    paddingHorizontal: 10,
  },
  inputImage: {
    marginLeft: '3%',
    marginBottom: '5%',
    color: COLOR.TEXT_GRAY
  },
  sendButton: {
    marginRight: '3%',
    marginBottom: '5%',
    color: COLOR.PRIMARY
  },
  sendButtonDisabled: {
    marginRight: '3%',
    marginBottom: '5%',
    color: COLOR.LIGHT_GREY
  },
  chatFooter: {
    backgroundColor: COLOR.BACKGROUND_WHITE,
    height: 'auto',
    width: '100%',
    flexDirection: 'row',
    justifyContent: 'center',
    alignItems: 'flex-end',
    bottom: 0,
    minHeight: '10%',
    maxHeight: '20%'
  },
});

export default MessagesScreen;
