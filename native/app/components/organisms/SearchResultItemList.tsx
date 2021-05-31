import React, { ReactElement } from 'react';
import { StyleSheet, View } from 'react-native';
import { ISearchResultItem } from '~/types/response/external/rakuten-books';
import SearchResultItem from '~/components/molecules/SearchResultItem';

const styles = StyleSheet.create({});

interface Props {
  resultItems: ISearchResultItem[];
  onPress: (item: Partial<ISearchResultItem>) => void;
}

const SearchResultItemList = function SearchResultItemList(props: Props): ReactElement {
  const items = props.resultItems;

  return (
    <View>
      {items.map((item) => (
        <SearchResultItem key={item.isbn} book={item} onPress={() => props.onPress(item)} />
      ))}
    </View>
  );
};

export default SearchResultItemList;
