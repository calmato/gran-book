import { StackNavigationProp } from '@react-navigation/stack';
import React, { ReactElement, useCallback, useEffect, useState } from 'react';
import { StyleSheet, ScrollView, RefreshControl, View } from 'react-native';

import { Header, Tab, TabView } from 'react-native-elements';
import HeaderText from '~/components/atoms/HeaderText';
import BookList from '~/components/molecules/BookList';
import SearchBar from '~/components/molecules/SearchBar';
import { ViewBooks } from '~/types/models/book';
import { BookshelfTabStackPramList } from '~/types/navigation';
import { IBook } from '~/types/response';
import { COLOR } from '~~/constants/theme';

const styles = StyleSheet.create({
  tabTitle: {
    fontSize: 14,
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

const tabList = [
  { title: '読んだ本', alias: 'read' },
  { title: '読んでいる本', alias: 'reading' },
  { title: '積読本', alias: 'stack' },
  { title: '欲しい本', alias: 'want' },
  { title: '手放したい本', alias: 'release' },
];

interface Props {
  navigation?: StackNavigationProp<BookshelfTabStackPramList, 'Bookshelf'>;
  actions: {
    getAllBook: () => Promise<void>;
  };
  books: ViewBooks;
}

const Bookshelf = function Bookshelf(props: Props): ReactElement {
  const navigation = props.navigation;
  const [keyword, setKeyword] = useState<string>('');
  const [refreshing, setRefreshing] = useState<boolean>(false);
  const [index, setIndex] = useState<number>(0);

  const onSubmitEditingCallback = useCallback(() => {
    (async () => {
      if (keyword !== '') {
        navigation?.navigate('SearchResult', { keyword });
      }
    })();
  }, [keyword, navigation]);

  const handleBookClick = useCallback(
    (book: IBook) => {
      navigation?.navigate('BookShow', { book: book });
    },
    [navigation],
  );

  const cancelCallback = useCallback(() => {
    return setKeyword('');
  }, [setKeyword]);

  useEffect(() => {
    props.actions.getAllBook();
  }, [props.actions]);

  return (
    <View>
      <Header centerComponent={<HeaderText title="Gran Book" />} />
      <SearchBar
        onCancel={cancelCallback}
        keyword={keyword}
        onChangeText={(text) => setKeyword(text)}
        onSubmitEditing={onSubmitEditingCallback}
      />

      <ScrollView horizontal={true} bounces={false} showsHorizontalScrollIndicator={false}>
        <Tab value={index} onChange={setIndex} indicatorStyle={styles.indicator}>
          {tabList.map((item, idx) => {
            return <Tab.Item key={idx} title={item.title} titleStyle={styles.tabTitle} />;
          })}
        </Tab>
      </ScrollView>

      <ScrollView
        refreshControl={
          <RefreshControl
            refreshing={refreshing}
            onRefresh={() => {
              setRefreshing(true);
              props.actions.getAllBook();
              setRefreshing(false);
            }}
          />
        }
        keyboardShouldPersistTaps="handled"
        style={{ marginBottom: 'auto' }}>
        <TabView value={index} onChange={setIndex}>
          {props.books &&
            tabList.map((item, idx) => {
              return (
                <TabView.Item key={idx} style={styles.tabView}>
                  <BookList books={props.books[item.alias]} handleClick={handleBookClick} />
                </TabView.Item>
              );
            })}
        </TabView>
      </ScrollView>
    </View>
  );
};

export default Bookshelf;
