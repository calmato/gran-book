import { StackNavigationProp } from '@react-navigation/stack';
import React, { ReactElement, useCallback, useEffect, useState } from 'react';
import { StyleSheet, ScrollView, RefreshControl, View } from 'react-native';

import { Header, Tab, TabView } from 'react-native-elements';
import HeaderText from '~/components/atoms/HeaderText';
import BookList from '~/components/molecules/BookList';
import SearchBar from '~/components/molecules/SearchBar';
import { ViewBooks } from '~/types/models/book';
import { HomeTabStackPramList } from '~/types/navigation';
import { IBook } from '~/types/response';
import { COLOR } from '~~/constants/theme';

const styles = StyleSheet.create({
  pickerStyle: {
    backgroundColor: COLOR.BACKGROUND_WHITE,
    height: 24,
    justifyContent: 'center',
    alignContent: 'center',
    paddingHorizontal: 12,
  },
  inputIOS: {
    fontSize: 12,
    fontWeight: 'bold',
    color: COLOR.TEXT_TITLE,
  },
  inputAndroid: {
    fontSize: 12,
    fontWeight: 'bold',
    color: COLOR.TEXT_TITLE,
  },
  tabTitle: {
    fontSize: 12,
    color: COLOR.GREY,
  },
  indicator: {
    height: 3,
    backgroundColor: COLOR.PRIMARY,
  },
  tabView: {
    width: '100%',
    height: 800,
  },
});

interface Props {
  navigation?: StackNavigationProp<HomeTabStackPramList, 'Home'>;
  actions: {
    getAllBook: () => Promise<void>;
  };
  books: ViewBooks;
}

const Home = function Home(props: Props): ReactElement {
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
          <Tab.Item title="呼んだ本" titleStyle={styles.tabTitle} />
          <Tab.Item title="読んでいる本" titleStyle={styles.tabTitle} />
          <Tab.Item title="積読本" titleStyle={styles.tabTitle} />
          <Tab.Item title="欲しい本" titleStyle={styles.tabTitle} />
          <Tab.Item title="手放したい本" titleStyle={styles.tabTitle} />
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
        style={{ marginBottom: 'auto', height: '100%' }}>
        <TabView value={index} onChange={setIndex}>
          <TabView.Item style={styles.tabView}>
            {props.books && <BookList books={props.books.read} handleClick={handleBookClick} />}
          </TabView.Item>

          <TabView.Item style={styles.tabView}>
            {props.books && <BookList books={props.books.reading} handleClick={handleBookClick} />}
          </TabView.Item>

          <TabView.Item style={styles.tabView}>
            {props.books && <BookList books={props.books.stack} handleClick={handleBookClick} />}
          </TabView.Item>

          <TabView.Item style={styles.tabView}>
            {props.books && <BookList books={props.books.want} handleClick={handleBookClick} />}
          </TabView.Item>

          <TabView.Item style={styles.tabView}>
            {props.books && <BookList books={props.books.release} handleClick={handleBookClick} />}
          </TabView.Item>
        </TabView>
      </ScrollView>
    </View>
  );
};

export default Home;
