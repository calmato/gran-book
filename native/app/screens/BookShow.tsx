import { RouteProp } from '@react-navigation/native';
import { StackNavigationProp } from '@react-navigation/stack';
import * as WebBrowser from 'expo-web-browser';
import React, { ReactElement, useCallback, useContext, useEffect, useState } from 'react';
import { StyleSheet, View } from 'react-native';
import { Overlay, Tab, TabView, Text } from 'react-native-elements';
import { ScrollView } from 'react-native-gesture-handler';
import BookImpression from '~/components/organisms/BookImpression';
import BookInfo from '~/components/organisms/BookInfo';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import { AuthContext } from '~/context/auth';
import { convertToIBook } from '~/lib/converter';
import { addBook, getAllImpressionByBookId, getBookByISBN } from '~/store/usecases/v2/book';
import { BookshelfV1Response } from '~/types/api/bookshelf_apiv1_response_pb';
import { BookReviewListV1Response } from '~/types/api/review_apiv1_response_pb';
import { BookshelfTabStackParamList } from '~/types/navigation';
import { ISearchResultItem } from '~/types/response/external/rakuten-books';
import { COLOR, FONT_SIZE } from '~~/constants/theme';
import { BookReviewV1Response } from '~~/tmp/proto/gateway/native/review_apiv1_response_pb';

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
    | RouteProp<BookshelfTabStackParamList, 'SearchResultBookShow'>
    | RouteProp<BookshelfTabStackParamList, 'BookShow'>;
  navigation:
    | StackNavigationProp<BookshelfTabStackParamList, 'SearchResultBookShow'>
    | StackNavigationProp<BookshelfTabStackParamList, 'BookShow'>;
  actions: {
    registerBook: (
      bookId: number,
      status: 'reading' | 'read' | 'stack' | 'release' | 'want',
    ) => Promise<void>;
  };
}

const BookShow = function BookShow(props: Props): ReactElement {
  const navigation = props.navigation;
  const routeParam = props.route.params;

  const [_wbResult, setWbResult] = useState<WebBrowser.WebBrowserResult>();
  const [showMessage, setShowMessage] = useState<boolean>(false);

  const [isRegister, setIsRegister] = useState<boolean>('id' in routeParam.book);
  const [book, setBook] = useState<BookshelfV1Response.AsObject>(
    'id' in routeParam.book ? routeParam.book : convertToIBook(routeParam.book),
  );

  const [impressions, setImpressions] = useState<BookReviewListV1Response.AsObject>({
    total: 0,
    offset: 0,
    limit: 0,
    reviewsList: [] as BookReviewV1Response.AsObject[],
  });

  const [index, setIndex] = useState<number>(0);

  const { authState } = useContext(AuthContext);

  // TODO: エラーハンドリング
  const handleAddBookButton = useCallback(async () => {
    const res = await addBook({ book: routeParam.book as ISearchResultItem }, authState.token);
    setBook(res);
    setShowMessage(true);
    setIsRegister(true);
  }, [authState.token, routeParam.book]);

  const handleBookStatusButton = useCallback(
    (status: 'reading' | 'read' | 'stack' | 'release' | 'want') => {
      if (status === 'read') {
        props.navigation.push('BookReadRegister', { book: book });
        return;
      }

      props.actions.registerBook(book.id, status);

      setBook({
        ...book,
        bookshelf: {
          ...book.bookshelf,
          status,
        },
      });
    },
    [props.actions, props.navigation, book],
  );

  const _handleOpenRakutenPageButtonAsync = async (url: string) => {
    const r = await WebBrowser.openBrowserAsync(url);
    setWbResult(r);
  };

  useEffect(() => {
    const f = async () => {
      try {
        const res = await getBookByISBN({ isbn: book.isbn }, authState.token);
        setBook(res);
        setIsRegister(true);
        const impRes = await getAllImpressionByBookId({ bookId: book.id }, authState.token);
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
