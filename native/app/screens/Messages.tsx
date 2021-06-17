
import React, { useState, useEffect } from 'react';
import { View, StyleSheet } from 'react-native';
import { StatusBar as ExpoStatusBar } from 'expo-status-bar';
import firebase from '~/lib/firebase';
import { getMessageDocRef } from '~/store/usecases/auth';
import { COLOR } from '~~/constants/theme';
import { MaterialIcons, Ionicons } from '@expo/vector-icons';
import { TransferMessageForm } from '~/types/forms';
import { Header } from 'react-native-elements';
import { GiftedChat, Send } from 'react-native-gifted-chat';
import { MaterialIcons, Ionicons, MaterialCommunityIcons } from '@expo/vector-icons';
import { StatusBar as ExpoStatusBar } from 'expo-status-bar';
import React, { useState, useMemo, useEffect } from 'react';
import {
  View,
  StyleSheet,
  TextInput,
  FlatList,
} from 'react-native';
import { Header } from 'react-native-elements';
import { ScrollView } from 'react-native-gesture-handler';
import { MessageItem } from '~/components/organisms/MessageItem';
import firebase from '~/lib/firebase';
import { getMessageDocRef } from '~/store/usecases/auth';
import { MessageForm } from '~/types/forms';
import { COLOR } from '~~/constants/theme';

export const MessagesScreen = () => {
  const [messages, setMessages] = useState<TransferMessageForm[]>([]);
  const docRef = getMessageDocRef();
  let initialStatus = true;
  const onSend = (newMessages: TransferMessageForm[]) => {
    newMessages.forEach(async (newMessage) => {
      (await docRef).add(newMessage);
    });
  };
  const getMessage =  () => {
    firebase.firestore().collection('messages').orderBy('createdAt', 'desc').limit(30)
      .onSnapshot( (snapshot) => {
        snapshot.docChanges().forEach((change) => {
          if (change?.type === 'added'){
            const messageInfo: TransferMessageForm = {
              _id: change.doc.data()._id,
              createdAt: change.doc.data().createdAt.toDate(),
              text: change.doc.data().text,
              user: change.doc.data().user,
            };
            if (initialStatus){
              setMessages((messages)=>[...messages, messageInfo]);
            } else {
              setMessages((messages)=> [messageInfo, ...messages]);
            }}
        });
        initialStatus = false;
      });
  };
  useEffect(() => {
    getMessage();
  });

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
        renderSend= {(props) => {
          return (
            <Send {...props} containerStyle={styles.sendContainer}>
              <Ionicons
                style={styles.sendButton}
                name= "send"
                size= {32}
              />
            </Send>
          );
        }}
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
  sendContainer: {
    justifyContent: 'center',
    alignItems: 'flex-end',
    flexDirection: 'row',
    height: 'auto',
    minHeight: '10%',
    marginRight: '3%',
    bottom: 0,
  },
  sendButton: {
    marginRight: '3%',
    marginBottom: '5%',
    color: COLOR.PRIMARY
  },
});

export default MessagesScreen;
