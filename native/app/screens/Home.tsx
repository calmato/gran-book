import React, { ReactElement, useCallback, useEffect, useMemo, useState } from 'react';
import { StackNavigationProp } from '@react-navigation/stack';
import { StyleSheet, ScrollView } from 'react-native';
import { View } from 'react-native';
import { Header } from 'react-native-elements';
import HeaderText from '~/components/atoms/HeaderText';
import BookList from '~/components/molecules/BookList';
import SearchBar from '~/components/molecules/SearchBar';
import { searchBookByTitle } from '~/lib/rakuten-books';
import { HomeTabStackPramList } from '~/types/navigation';
import { ViewBooks } from '~/types/models/book';
import { IBook } from '~/types/response';
import RNPickerSelect from 'react-native-picker-select';
import { COLOR } from '~~/constants/theme';
import { MaterialIcons } from '@expo/vector-icons';

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
});

interface Props {
  navigation?: StackNavigationProp<HomeTabStackPramList, 'Home'>;
  actions: {
    getAllBook: () => Promise<void>
  };
  books: ViewBooks;
}

const pickerItems = [
  { label: '読んだ本', value: 'read' },
  { label: '読んでいる本', value: 'reading' },
  { label: '積読本', value: 'stack' },
  { label: '欲しい本', value: 'want' },
  { label: '手放したい本', value: 'release' },
];

const iconComponent = () => {
  return <MaterialIcons name="keyboard-arrow-down" size={16} color={COLOR.TEXT_TITLE} />;
};

const Home = function Home(props: Props): ReactElement {
  const navigation = props.navigation;
  const [books, setBooks] = useState(props.books.read);
  const [value, setValue] = useState('read');
  const [keyword, setKeyword] = useState('');

  const onSubmitEditingCallback = useCallback(() => {
    (async () => {
      if (keyword !== '') { // TODO: titleに変更する？
        const res = await searchBookByTitle(keyword);
        if (res) navigation?.navigate('SearchResult', { keyword, results: res.data });
      }
    })();
  }, [keyword, navigation]);

  const handleBookClick = useCallback((book: IBook) => {
    navigation?.navigate('BookShow', { book: book });
  }, [ navigation ]);

  const handlePickerSelect = useCallback((value) => {
    setValue(value);
  },[]);

  const cancelCallback = useCallback(() => {
    return setKeyword('');
  }, [setKeyword]);

  useEffect(() => {
    props.actions.getAllBook();
    console.log(props.books.read.length);
    console.log(props.books.reading.length);
    console.log(props.books.stack.length);
    console.log(props.books.want.length);
    console.log(props.books.release.length);
  }, [props.actions]);

  useEffect(() => {
    switch (value) {
      case 'read':
        setBooks(props.books.read);
        break;
      case 'reading':
        setBooks(props.books.reading);
        break;
      case 'want':
        setBooks(props.books.want);
        break;
      case 'release':
        setBooks(props.books.release);
        break;
      case 'stack':
        setBooks(props.books.stack);
        break;
      default:
        break;
    }
  }, [value, props.books]);

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
          onChangeText={(text) => setKeyword(text)}
          onSubmitEditing={onSubmitEditingCallback}
        />
        <RNPickerSelect
          style={{
            viewContainer: styles.pickerStyle,
            inputIOS: styles.inputIOS,
            inputAndroid: styles.inputAndroid,
          }}
          placeholder={{}}
          onValueChange={handlePickerSelect}
          value={value}
          items={pickerItems}
          Icon={iconComponent}
        />
        {
          books? <BookList books={books} handleClick={handleBookClick} />: null
        }
      </ScrollView>
    </View>
  );
};

export default Home;
