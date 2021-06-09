import { RouteProp } from '@react-navigation/native';
import { StackNavigationProp } from '@react-navigation/stack';
import React, { ReactElement, useCallback, useEffect, useState } from 'react';
import { StyleSheet, View, Text, Alert } from 'react-native';
import { ScrollView } from 'react-native-gesture-handler';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import SearchResultItemList from '~/components/organisms/SearchResultItemList';
import { searchBookByTitle } from '~/lib/rakuten-books';
import { generateErrorMessage } from '~/lib/util/ErrorUtil';
import { HomeTabStackPramList } from '~/types/navigation';
import { ISearchResponse, ISearchResultItem } from '~/types/response/external/rakuten-books';
import { COLOR } from '~~/constants/theme';

const styles = StyleSheet.create({
  totalItemsTextStyle: {
    color: COLOR.TEXT_DEFAULT,
    marginLeft: 10,
    padding: 4,
  },
});

interface Props {
  route: RouteProp<HomeTabStackPramList, 'SearchResult'>;
  navigation: StackNavigationProp<HomeTabStackPramList, 'SearchResult'>;
}

const SearchResult = function SearchResult(props: Props): ReactElement {
  const { keyword } = props.route.params;
  const navigation = props.navigation;

  const [result, setResult] = useState<ISearchResponse>();
  const [books, setBooks] = useState<ISearchResultItem[]>([]);
  const [page, setPage] = useState<number>(1);

  const selectBook = useCallback(
    (item) => {
      return navigation.navigate('SearchResultBookShow', { book: item });
    },
    [navigation],
  );

  const createAlert = (code: number) => {
    Alert.alert('検索結果の取得に失敗しました', `${generateErrorMessage(code)}`, [
      {
        text: 'OK',
      },
    ]);
  };

  // TODO: パフォーマンス面で要リファクタ
  useEffect(() => {
    let unmounted = false;
    const f = async () => {
      try {
        const res = await searchBookByTitle(keyword, page);
        if (!unmounted) {
          if (page === 1) {
            setResult(res.data);
            setBooks(res.data.Items);
          } else if (result && page <= result?.pageCount) {
            setBooks([...books, ...res.data.Items]);
          }
        }
      } catch (e) {
        createAlert(e.status);
        return;
      }
    };
    f();
    return () => {
      unmounted = true;
    };
  }, [page]);

  return (
    <View>
      <HeaderWithBackButton onPress={() => navigation.goBack()} title={keyword} />
      <Text style={styles.totalItemsTextStyle}>検索結果：{result?.count}件</Text>
      <ScrollView
        style={{ marginBottom: 'auto' }}
        onMomentumScrollEnd={() => {
          setPage(page + 1);
        }}>
        <SearchResultItemList resultItems={books} onPress={selectBook} />
      </ScrollView>
    </View>
  );
};

export default SearchResult;
