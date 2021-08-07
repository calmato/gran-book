import { RouteProp } from '@react-navigation/native';
import { StackNavigationProp } from '@react-navigation/stack';
import * as WebBrowser from 'expo-web-browser';
import React, { ReactElement, useCallback, useEffect, useState } from 'react';
import { StyleSheet, View } from 'react-native';
import { Overlay, Tab, TabView, Text } from 'react-native-elements';
import { ScrollView } from 'react-native-gesture-handler';
import BookImpression from '../components/organisms/BookImpression';
import BookInfo from '~/components/organisms/BookInfo';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import { convertToIBook } from '~/lib/converter';
import { addBookAsync, getAllImpressionByBookIdAsync, getBookByISBNAsync } from '~/store/usecases';
import { HomeTabStackPramList } from '~/types/navigation';
import { IBook, IImpressionResponse } from '~/types/response';
import { ISearchResultItem } from '~/types/response/external/rakuten-books';
import { COLOR, FONT_SIZE } from '~~/constants/theme';

const styles = StyleSheet.create({
  menuActiveFontStyle: {
    color: COLOR.PRIMARY,
  },
  tabTitle: {
    fontSize: FONT_SIZE.TAB_TITLE,
    color: COLOR.GREY,
  },
  indicator: {
    height: 3,
    backgroundColor: COLOR.PRIMARY,
  },
  tabView: {
    width: '100%',
    height: '100%',
  },
});

interface Props {
  route:
    | RouteProp<HomeTabStackPramList, 'SearchResultBookShow'>
    | RouteProp<HomeTabStackPramList, 'BookShow'>;
  navigation:
    | StackNavigationProp<HomeTabStackPramList, 'SearchResultBookShow'>
    | StackNavigationProp<HomeTabStackPramList, 'BookShow'>;
  actions: {
    registerOwnBook: (status: string, bookId: number) => Promise<void>;
  };
}

const BookShow = function BookShow(props: Props): ReactElement {
  const navigation = props.navigation;
  const routeParam = props.route.params;

  const [_wbResult, setWbResult] = useState<WebBrowser.WebBrowserResult>();
  const [showMessage, setShowMessage] = useState<boolean>(false);

  const [isRegister, setIsRegister] = useState<boolean>('id' in routeParam.book);
  const [book, setBook] = useState<IBook>(
    'id' in routeParam.book ? routeParam.book : convertToIBook(routeParam.book),
  );

  const [impressions, setImpressions] = useState<IImpressionResponse>();

  const [index, setIndex] = useState<number>(0);

  // TODO: エラーハンドリング
  const handleAddBookButton = useCallback(async () => {
    await addBookAsync(routeParam.book as ISearchResultItem)
      .then((res) => {
        setBook(res.data);
        setShowMessage(true);
        setIsRegister(true);
      })
      .catch((res) => console.log('登録に失敗しました.', res));
  }, [routeParam.book]);

  const handleBookStatusButton = useCallback(
    (status: string) => {
      if (status === 'read') {
        props.navigation.push('BookReadRegister', { book: book });
        return;
      }
      props.actions.registerOwnBook(status, book.id);
      setBook({
        ...book,
        bookshelf: {
          ...book.bookshelf,
          status: status,
        },
      });
    },
    [props.actions, props.navigation, book],
  );

  const _handleOpenRakutenPageButtonAsync = async (url: string) => {
    const r = await WebBrowser.openBrowserAsync(url);
    setWbResult(r);
  };

  // TODO エラーハンドリングの処理を分ける
  useEffect(() => {
    const f = async () => {
      try {
        const res = await getBookByISBNAsync(book.isbn);
        setBook(res.data);
        setIsRegister(true);
        const impRes = await getAllImpressionByBookIdAsync(book.id);
        setImpressions(impRes);
      } catch (err) {
        setIsRegister(false);
      }
    };
    f();
  }, []);

  return (
    <View>
      <Overlay
        backdropStyle={{
          opacity: 0.8,
        }}
        isVisible={showMessage}
        onBackdropPress={() => setShowMessage(false)}>
        <View
          style={{
            width: 'auto',
            height: 'auto',
            justifyContent: 'center',
            margin: 8,
          }}>
          <Text style={{ fontSize: FONT_SIZE.BOOK_INFO_TITLE, color: COLOR.TEXT_TITLE, margin: 4 }}>
            「{book.title}」
          </Text>
          <Text> を登録しました。</Text>
        </View>
      </Overlay>
      <HeaderWithBackButton onPress={() => navigation.goBack()} title={book.title} />

      <Tab value={index} onChange={setIndex} indicatorStyle={styles.indicator}>
        <Tab.Item title="情報" titleStyle={styles.tabTitle} />
        <Tab.Item title="感想" titleStyle={styles.tabTitle} />
      </Tab>

      <ScrollView style={{ marginBottom: 'auto' }}>
        <TabView value={index} onChange={setIndex}>
          <TabView.Item style={styles.tabView}>
            <BookInfo
              book={book}
              isRegister={isRegister}
              handleBookStatusButton={handleBookStatusButton}
              handleOpenRakutenPageButton={_handleOpenRakutenPageButtonAsync}
              handleAddButton={handleAddBookButton}
            />
          </TabView.Item>
          <TabView.Item style={styles.tabView}>
            {impressions && <BookImpression book={book} impressionResponse={impressions} />}
          </TabView.Item>
        </TabView>
      </ScrollView>
    </View>
  );
};

export default BookShow;
