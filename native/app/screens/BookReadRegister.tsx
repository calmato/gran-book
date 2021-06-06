import { RouteProp } from '@react-navigation/native';
import { StackNavigationProp } from '@react-navigation/stack';
import React, { ReactElement, useState } from 'react';
import { StyleSheet, View, ScrollView, Text } from 'react-native';
import { Input, Button } from 'react-native-elements';
import BookNameAuthorRegister from '~/components/organisms/BookNameAuthorRegister';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import ReadDate from '~/components/organisms/ReadDate';
import { HomeTabStackPramList } from '~/types/navigation';

const styles = StyleSheet.create({
  container: {
    flex: 1,
  },
});

interface Props {
  route: RouteProp<HomeTabStackPramList, 'BookReadRegister'>;
  navigation: StackNavigationProp<HomeTabStackPramList, 'BookReadRegister'>;
}

const BookReadRegister = function BookReadRegister(props: Props): ReactElement {
  const book = props.route.params.book;

  const [impreessionData, setState] = useState({
    date: new Date(),
    impresstion: '',
    isDateUnknown: false,
  });

  return (
    <View style={styles.container}>
      <HeaderWithBackButton
        title="読んだ本登録"
        onPress={() => {
          props.navigation.goBack();
        }}
      />
      <ScrollView>
        <BookNameAuthorRegister
          title={book.title}
          imageUrl={book.thumbnailUrl}
          author={book.author}
        />
        <ReadDate
          date={impreessionData.date}
          handleSetDate={(date) => setState({ ...impreessionData, date: date })}
          isDateUnknown={impreessionData.isDateUnknown}
          handleIsDateUnknown={(isDateUnknown) =>
            setState({ ...impreessionData, isDateUnknown: isDateUnknown })
          }
        />
        <Text
          style={{
            fontSize: 16,
            marginStart: 20,
            marginTop: 20,
            marginBottom: 10,
            fontWeight: 'bold',
          }}>
          感想
        </Text>
        <Input
          onChangeText={(text) => setState({ ...impreessionData, impresstion: text })}
          value={impreessionData.impresstion}
          maxLength={1000}
          multiline={true}
        />
        <View style={{ alignItems: 'center', marginBottom: 20 }}>
          <Button onPress={undefined} title="本を登録する" />
        </View>
      </ScrollView>
    </View>
  );
};

export default BookReadRegister;
