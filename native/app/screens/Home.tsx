import React, { ReactElement, useCallback, useEffect, useState } from 'react';
import { StackNavigationProp } from '@react-navigation/stack';
import { ScrollView } from 'react-native';
import { View } from 'react-native';
import { Header } from 'react-native-elements';
import HeaderText from '~/components/atoms/HeaderText';
import BookList from '~/components/molecules/BookList';
import SearchBar from '~/components/molecules/SearchBar';
import { searchBookByTitle } from '~/lib/rakuten-books';
import { HomeTabStackPramList } from '~/types/navigation';
import { ViewBooks } from '~/types/models/book';

interface Props {
  navigation?: StackNavigationProp<HomeTabStackPramList, 'Home'>;
  actions: {
    getAllBook: () => Promise<void>
  }
  books: ViewBooks
}

const Home = function Home(props: Props): ReactElement {
  const navigation = props.navigation;
  const books = props.books;
  const [keyword, setValue] = useState('');

  const onSubmitEditingCallback = useCallback(() => {
    (async () => {
      if (keyword !== '') { // TODO: titleに変更する？
        const res = await searchBookByTitle(keyword);
        if (res) navigation?.navigate('SearchResult', { keyword, results: res.data });
      }
    })();
  }, [keyword, navigation]);

  const cancelCallback = useCallback(() => {
    return setValue('');
  }, [setValue]);

  useEffect(() => {
    props.actions?.getAllBook();
  }, [props.actions]);

  return (
    <View>
      <Header centerComponent={<HeaderText title="Gran Book" />} />
      <ScrollView
        stickyHeaderIndices={[0]}
        keyboardShouldPersistTaps="handled"
        style={{ marginBottom: 'auto', height: '100%' }}>
        <SearchBar
          onCancel={cancelCallback}
          keyword={keyword}
          onChangeText={(text) => setValue(text)}
          onSubmitEditing={onSubmitEditingCallback}
        />
        {
          books? <BookList books={books.read} />: null
        }
      </ScrollView>
    </View>
  );
};

export default Home;
