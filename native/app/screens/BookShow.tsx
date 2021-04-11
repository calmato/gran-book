import React, { ReactElement } from 'react';
import { StyleSheet, View } from 'react-native';
import { Button, Image, Text } from 'react-native-elements';
import FlexBoxBookCategory from '~/components/organisms/FlexBoxBookCategory';
import { COLOR } from '~~/constants/theme';
import ButtonGroupBookFooter from '~/components/organisms/ButtonGroupBookFooter';
import { ScrollView } from 'react-native-gesture-handler';
import { MaterialCommunityIcons } from '@expo/vector-icons';
import { RouteProp } from '@react-navigation/native';
import { StackNavigationProp } from '@react-navigation/stack';
import { HomeTabStackPramList } from '~/types/navigation';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';

const styles = StyleSheet.create({
  container: {
    alignItems: 'center',
  },
  imageContainer: {
    marginTop: 30,
    marginBottom: 30,
    width: 200,
    height: 280,
    resizeMode: 'contain',
  },
  titleContainer: {
    paddingTop: 10,
    paddingLeft: 30,
    paddingRight: 30,
    fontSize: 16,
    alignSelf: 'stretch',
    color: COLOR.GREY,
    backgroundColor: COLOR.BACKGROUND_WHITE,
  },
  authorContainer: {
    paddingTop: 10,
    paddingLeft: 30,
    paddingRight: 30,
    paddingBottom: 10,
    fontSize: 16,
    alignSelf: 'stretch',
    backgroundColor: COLOR.BACKGROUND_WHITE,
  },
  detailContainer: {
    marginTop: 20,
    paddingTop: 10,
    paddingLeft: 20,
    paddingRight: 20,
    paddingBottom: 10,
    fontSize: 16,
    alignSelf: 'stretch',
    backgroundColor: COLOR.BACKGROUND_WHITE,
  },
});

interface Props {
  route: RouteProp<HomeTabStackPramList, 'SearchResultBookShow'>;
  navigation: StackNavigationProp<HomeTabStackPramList, 'SearchResultBookShow'>;
}

const BookShow = function BookShow(props: Props): ReactElement {
  const navigation = props.navigation;
  const { book } = props.route.params;

  return (
    <View>
      <HeaderWithBackButton onPress={() => navigation.goBack()} title={book.title} />
      <ScrollView contentContainerStyle={styles.container} style={{ marginBottom: 'auto' }}>
        <View
          style={{
            flexDirection: 'row',
            alignSelf: 'stretch',
            justifyContent: 'space-around',
            alignItems: 'center',
          }}>
          <MaterialCommunityIcons name="chevron-left-circle" size={36} color={COLOR.TEXT_GRAY} />
          <Image
            // TODO 本の情報を代入
            source={
              book.largeImageUrl
                ? { uri: book.largeImageUrl }
                : require('assets/logo.png')
            }
            style={styles.imageContainer}
            transition={true}
          />
          <MaterialCommunityIcons name="chevron-right-circle" size={36} color={COLOR.TEXT_GRAY} />
        </View>
        <Text style={styles.titleContainer}>題名: {book.title}</Text>
        <Text style={styles.authorContainer}>
          著者:{' '}
          {book.author ? book.author : '著者情報がありません'}
        </Text>
        <Text style={styles.detailContainer}>本の詳細</Text>
        <Text>{book.contents}</Text>
        <FlexBoxBookCategory
          category={
            book.booksGenreId
              ? book.booksGenreId
              : 'カテゴリ情報がありません'
          }
        />
        <ButtonGroupBookFooter
          handleNavigateToReadBoook={() => undefined}
          handleNavigateToReadingBoook={() => undefined}
          handleNavigateToTsundoku={() => undefined}
          handleNavigateToSellBoook={() => undefined}
          handleNavigateToWishList={() => undefined}
        />
        <Button
          onPress={() => undefined}
          title="楽天で見る"
          containerStyle={{ marginTop: 10, marginBottom: 10 }}
          buttonStyle={{ backgroundColor: COLOR.PRIMARY_DARK }}
        />
        <Button
          onPress={() => undefined}
          title="本を買う"
          containerStyle={{ marginTop: 10, marginBottom: 10 }}
        />
      </ScrollView>
    </View>
  );
};

export default BookShow;
