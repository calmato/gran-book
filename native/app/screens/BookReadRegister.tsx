import { RouteProp } from '@react-navigation/native';
import { StackNavigationProp } from '@react-navigation/stack';
import React, { ReactElement, useCallback, useState } from 'react';
import { StyleSheet, View, ScrollView, Text } from 'react-native';
import { Button, Input } from 'react-native-elements';
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

  const [impressionData, setState] = useState({
    date: new Date(),
    impression: '',
    isDateUnknown: false,
  });

  const handleRegisterButtonClick = useCallback(() => {
    props.navigation.goBack();
  }, [props.navigation]);

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
          date={impressionData.date}
          handleSetDate={(date) => setState({ ...impressionData, date: date })}
          isDateUnknown={impressionData.isDateUnknown}
          handleIsDateUnknown={(isDateUnknown) =>
            setState({ ...impressionData, isDateUnknown: isDateUnknown })
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
          onChangeText={(text) => setState({ ...impressionData, impression: text })}
          value={impressionData.impression}
          maxLength={1000}
          multiline={true}
        />
        <View style={{ alignItems: 'center', marginBottom: 20 }}>
          <Button onPress={handleRegisterButtonClick} title="本を登録する" />
        </View>
      </ScrollView>
    </View>
  );
};

export default BookReadRegister;
