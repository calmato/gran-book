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
import { IBook, IBookResponse } from '~/types/response';

interface Props {
  navigation?: StackNavigationProp<HomeTabStackPramList, 'Home'>;
  actions: {
    getAllBook: () => Promise<void>
  }
  books?: IBook[],
}

const Home = function Home(props: Props): ReactElement {
  const navigation = props.navigation;
  const [keyword, setValue] = useState('');
  // const [books, setBooks] = useState<IBookResponse>();
  const books = props.books;

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
        style={{ marginBottom: 'auto' }}>
        <SearchBar
          onCancel={cancelCallback}
          keyword={keyword}
          onChangeText={(text) => setValue(text)}
          onSubmitEditing={onSubmitEditingCallback}
        />
        {
          books? <BookList books={books} />: null
        }
      </ScrollView>
    </View>
  );
};

export default Home;
