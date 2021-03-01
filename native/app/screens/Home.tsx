import React, { ReactElement, useCallback, useEffect, useState } from 'react';
import { ScrollView } from 'react-native';
import { View } from 'react-native';
import { Header } from 'react-native-elements';
import HeaderText from '~/components/atoms/HeaderText';
import BookList from '~/components/molecules/BookList';
import SearchBar from '~/components/molecules/SearchBar';
import { searchBook } from '~/lib/GoogleBooksAPI';
import { ISearchResponse } from '~/types/response/search';

const Home = function Home(): ReactElement {

  const [ keyword, setValue] = useState('');
  const [results, setResult] = useState<ISearchResponse>({
    kind: '',
    totalItems: 0,
    items: []
  });

  const onSubmitEditingCallback = useCallback(() => {
    (async () => {
      if(keyword !== '') {
        const res = await searchBook(keyword);
        if(res) setResult(res);
      }
    })();
  }, [keyword]
  );

  const cancelCallback = useCallback(
    () => {
      return setValue('');
    }, [setValue]
  );

  return (
    <View>
      <ScrollView
        stickyHeaderIndices={[0]}
        keyboardShouldPersistTaps="handled"
      >
        <Header centerComponent={<HeaderText title="Gran Book"/>} />
        <SearchBar
          onCancel={cancelCallback}
          keyword={keyword}
          onChangeText={(text) => setValue(text)}
          onSubmitEditing={onSubmitEditingCallback}
        />
        <BookList />
      </ScrollView>
    </View>
  );
};

export default Home;
