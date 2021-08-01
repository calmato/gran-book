import { RouteProp } from '@react-navigation/native';
import { StackNavigationProp } from '@react-navigation/stack';
import dayjs from 'dayjs';
import React, { ReactElement, useState } from 'react';
import { StyleSheet, View, ScrollView, Text, TextInput } from 'react-native';
import { Button } from 'react-native-elements';
import BookNameAuthorRegister from '~/components/organisms/BookNameAuthorRegister';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import ReadDate from '~/components/organisms/ReadDate';
import { ImpressionForm } from '~/types/forms';
import { HomeTabStackPramList } from '~/types/navigation';
import { COLOR } from '~~/constants/theme';

const styles = StyleSheet.create({
  container: {
    flex: 1,
  },
  impressionFormLabel: {
    fontSize: 15,
    paddingLeft: 16,
    marginTop: 10,
    marginBottom: 8,
    fontWeight: 'bold',
    color: COLOR.GREY,
  },
  impressionForm: {
    backgroundColor: COLOR.BACKGROUND_WHITE,
    padding: 0,
    fontSize: 18,
    paddingHorizontal: 16,
    height: 160,
    marginBottom: 16,
  },
});

interface Props {
  route: RouteProp<HomeTabStackPramList, 'BookReadRegister'>;
  navigation: StackNavigationProp<HomeTabStackPramList, 'BookReadRegister'>;
  actions: {
    registerReadBookImpression: (bookId: number, impression: ImpressionForm) => Promise<void>;
  };
}

const BookReadRegister = function BookReadRegister(props: Props): ReactElement {
  const book = props.route.params.book;

  const [impressionData, setState] = useState({
    date: new Date(),
    impression: '',
    isDateUnknown: false,
  });

  const handleRegisterButtonClick = () => {
    props.actions.registerReadBookImpression(book.id, {
      impression: impressionData.impression,
      readOn: dayjs(impressionData.date).format('YYYY-MM-DD'),
    });
    props.navigation.navigate('Home');
  };

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
        <Text style={styles.impressionFormLabel}>感想</Text>
        <TextInput
          style={styles.impressionForm}
          placeholder="ここに感想を入力"
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
