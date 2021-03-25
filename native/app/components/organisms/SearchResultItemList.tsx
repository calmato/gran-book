import React, { ReactElement } from 'react';
import { StyleSheet, View } from 'react-native';

import { ISearchResponse, ISearchResultItem } from '~/types/response/search';
import SearchResultItem from '~/components/molecules/SearchResultItem';

const styles = StyleSheet.create({});

interface Props {
  results: ISearchResponse;
  onPress: (item: Partial<ISearchResultItem>) => void;
}

const SearchResultItemList = function SearchResultItemList(props: Props): ReactElement {
  const results = props.results;

  return (
    <View>
      {results.items.map((item) => (
        <SearchResultItem key={item.id} book={item} onPress={() => props.onPress(item)} />
      ))}
    </View>
  );
};

export default SearchResultItemList;
