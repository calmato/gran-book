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
import { getAllBookAsync, getAllBookByUserId, registerReadBookAsync } from '~/store/usecases';
import { IBookResponse } from '~/types/response';

interface Props {
  navigation?: StackNavigationProp<HomeTabStackPramList, 'Home'>;
}

const Home = function Home(props: Props): ReactElement {
  const navigation = props.navigation;
  const [keyword, setValue] = useState('');
  const [books, setBooks] = useState<IBookResponse>();

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
    const f = async () => await getAllBookByUserId('e7ce84b7-dc23-440b-8e7f-025402195a92')
      .then((res) => { console.log(res.data); setBooks(res.data);})
      .catch((err) => console.log(err));
    f();
  }, []);

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
          books? <BookList books={books?.books} />: null
        }
      </ScrollView>
    </View>
  );
};

export default Home;
