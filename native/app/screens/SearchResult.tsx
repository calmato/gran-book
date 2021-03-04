import { RouteProp } from '@react-navigation/native';
import { StackNavigationProp } from '@react-navigation/stack';
import React, { ReactElement, useCallback } from 'react';
import { StyleSheet, View, Text } from 'react-native';
import { ScrollView } from 'react-native-gesture-handler';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import SearchResultItemList from '~/components/organisms/SearchResultItemList';
import { HomeTabStackPramList } from '~/types/navigation';
import { COLOR } from '~~/constants/theme';

const styles = StyleSheet.create({
  totalItemsTextStyle: {
    color: COLOR.TEXT_DEFAULT,
    marginLeft: 10,
    padding: 4,
  }
});

interface Props {
  route: RouteProp<HomeTabStackPramList, 'SearchResult'>,
  navigation: StackNavigationProp<HomeTabStackPramList, 'SearchResult'>
}

const SearchResult = function SearchResult(props: Props): ReactElement {
  const {keyword, results} = props.route.params;
  const navigation = props.navigation;

  const selectBook = useCallback((item) => {
    return navigation.navigate('SearchResultBookShow', {book: item});
  }, [navigation]);

  return (
    <View>
      <HeaderWithBackButton
        onPress={() => navigation.goBack()}
        title={keyword}
      />
      <Text style={styles.totalItemsTextStyle}>
        検索結果：{results.totalItems}件
      </Text>
      <ScrollView
        style={{ marginBottom: 'auto' }}
      >
        <SearchResultItemList
          results={results}
          onPress={selectBook}
        />
      </ScrollView>
    </View>
  );
};

export default SearchResult;
