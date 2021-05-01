import React, { ReactElement, useState } from 'react';
import { StyleSheet, View } from 'react-native';
import { Button, Image, Overlay, Text } from 'react-native-elements';
import FlexBoxBookCategory from '~/components/organisms/FlexBoxBookCategory';
import { COLOR } from '~~/constants/theme';
import ButtonGroupBookFooter from '~/components/organisms/ButtonGroupBookFooter';
import { ScrollView } from 'react-native-gesture-handler';
import { RouteProp } from '@react-navigation/native';
import { StackNavigationProp } from '@react-navigation/stack';
import { HomeTabStackPramList } from '~/types/navigation';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import * as WebBrowser from 'expo-web-browser';
import { addBookAsync } from '~/store/usecases';
import { fullWidth2halfWidth } from '~/lib/util';

const styles = StyleSheet.create({
  container: {
    alignItems: 'center',
  },
  imageContainer: {
    marginVertical: 24,
    width: 200,
    height: 280,
    resizeMode: 'contain',
  },
  titleContainer: {
    paddingTop: 10,
    paddingHorizontal: 16,
    fontSize: 16,
    alignSelf: 'stretch',
    color: COLOR.GREY,
    fontWeight: '500',
    backgroundColor: COLOR.BACKGROUND_WHITE,
  },
  authorContainer: {
    paddingTop: 10,
    paddingHorizontal: 16,
    paddingBottom: 10,
    fontSize: 12,
    alignSelf: 'stretch',
    backgroundColor: COLOR.BACKGROUND_WHITE,
  },
  detailContainer: {
    paddingHorizontal: 16,
    paddingBottom: 10,
    fontSize: 14,
    lineHeight: 16,
    color: COLOR.GREY,
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
  const [_wbResult, setWbResult] = useState<WebBrowser.WebBrowserResult>();
  const [showMessage, setShowMessage] = useState<boolean>(false);
  const [isRegister, _setIsRegister] = useState<boolean>(true);

  // TODO: エラーハンドリング
  const handleAddBookButton = async () => {
    return await addBookAsync(book)
      .then((res) => {
        setShowMessage(true);
      })
      .catch((res) => console.log('登録に失敗しました.', res));
  };

  const _handleOpenRakutenPageButtonAsync = async (url: string) => {
    const r = await WebBrowser.openBrowserAsync(url);
    setWbResult(r);
  };

  return (
    <View>
      <Overlay
        backdropStyle={{
          opacity: 0.8,
        }}
        isVisible={showMessage}
        onBackdropPress={() => setShowMessage(false)}
      >
        <View style={{
          width: 'auto',
          height: 'auto',
          justifyContent: 'center',
          margin: 8,
        }}>
          <Text style={{fontSize: 16, color: COLOR.TEXT_TITLE, margin: 4}}>「{book.title}」</Text>
          <Text> を登録しました。</Text>
        </View>
      </Overlay>
      <ScrollView contentContainerStyle={styles.container} style={{ marginBottom: 'auto' }}>
        <HeaderWithBackButton onPress={() => navigation.goBack()} title={book.title} />
        <View
          style={{
            alignSelf: 'stretch',
            justifyContent: 'space-around',
            alignItems: 'center',
          }}>
          <Image
            source={
              book.largeImageUrl
                ? { uri: book.largeImageUrl }
                : require('assets/logo.png')
            }
            style={styles.imageContainer}
            transition={true}
          />
        </View>
        <Text style={styles.titleContainer}>{book.title}</Text>
        <Text style={styles.authorContainer}>
          {book.author ? book.author : '著者情報がありません'}
        </Text>
        {
          book.itemCaption !== '' ?
            <Text style={styles.detailContainer}>
              {fullWidth2halfWidth(book.itemCaption)}
            </Text> : null
        }
        <FlexBoxBookCategory
          category={
            book.size
              ? book.size
              : 'カテゴリ情報がありません'
          }
        />
        {
          isRegister ?
            <ButtonGroupBookFooter
              status='読んだ本'
              handleNavigateToReadBoook={() => undefined}
              handleNavigateToReadingBoook={() => undefined}
              handleNavigateToTsundoku={() => undefined}
              handleNavigateToSellBoook={() => undefined}
              handleNavigateToWishList={() => undefined}
            /> :
            <Button
              title="本を登録する"
              onPress={() => handleAddBookButton()}
            />
        }
        <Button
          onPress={() => _handleOpenRakutenPageButtonAsync(book.itemUrl)}
          title="楽天で見る"
          containerStyle={{ marginTop: 10, marginBottom: 10 }}
          buttonStyle={{ backgroundColor: COLOR.PRIMARY_DARK }}
        />
      </ScrollView>
    </View>
  );
};

export default BookShow;
