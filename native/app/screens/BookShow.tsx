import SegmentedControl from '@react-native-segmented-control/segmented-control';
import { RouteProp } from '@react-navigation/native';
import { StackNavigationProp } from '@react-navigation/stack';
import * as WebBrowser from 'expo-web-browser';
import React, { ReactElement, useCallback, useEffect, useState } from 'react';
import { View } from 'react-native';
import { Overlay, Text } from 'react-native-elements';
import BookImpression from '../components/organisms/BookImpression';
import BookInfo from '~/components/organisms/BookInfo';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import { convertToIBook } from '~/lib/converter';
import { addBookAsync, getBookByISBNAsync } from '~/store/usecases';
import { HomeTabStackPramList } from '~/types/navigation';
import { IBook } from '~/types/response';
import { ISearchResultItem } from '~/types/response/external/rakuten-books';
import { COLOR } from '~~/constants/theme';

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

  const selectMenuList = ['情報', '感想'];
  const [selectedIndex, setIndex] = useState<number>(0);

  // TODO: エラーハンドリング
  const handleAddBookButton = async () => {
    return await addBookAsync(routeParam.book as ISearchResultItem)
      .then((res) => {
        setBook(res.data);
        setShowMessage(true);
        setIsRegister(true);
      })
      .catch((res) => console.log('登録に失敗しました.', res));
  };

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

  useEffect(() => {
    const f = async () => {
      try {
        const res = await getBookByISBNAsync(book.isbn);
        setBook(res.data);
        setIsRegister(true);
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
          <Text style={{ fontSize: 16, color: COLOR.TEXT_TITLE, margin: 4 }}>「{book.title}」</Text>
          <Text> を登録しました。</Text>
        </View>
      </Overlay>
      <HeaderWithBackButton onPress={() => navigation.goBack()} title={book.title} />
      <SegmentedControl
        values={selectMenuList}
        selectedIndex={selectedIndex}
        onValueChange={(event) => setIndex(selectMenuList.indexOf(event))}
      />
      {selectedIndex === 0 ? (
        <BookInfo
          book={book}
          isRegister={isRegister}
          handleBookStatusButton={handleBookStatusButton}
          handleOpenRakutenPageButton={_handleOpenRakutenPageButtonAsync}
          handleAddButton={handleAddBookButton}
        />
      ) : (
        <BookImpression />
      )}
    </View>
  );
};

export default BookShow;
