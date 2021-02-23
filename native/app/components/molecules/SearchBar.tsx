import React, { ReactElement } from 'react';
import { Platform, StyleSheet } from 'react-native';
import { SearchBar as ElementsSearchBar } from 'react-native-elements';
import SearchIcon from '../atoms/SearchIcon';

const styles = StyleSheet.create({
  inputContainerStyle: {
    height: 14,
  }
});

interface Props {
  placeholder?: string
}

const SearchBar = function SearchBar(props: Props): ReactElement {

  const platform = Platform.OS === 'ios'? 'ios' : 'android';

  return (
    <ElementsSearchBar
      placeholder={props.placeholder}
      inputContainerStyle={styles.inputContainerStyle}
      platform={platform}
      searchIcon={<SearchIcon />}
    />
  );
};

SearchBar.defaultProps={
  placeholder: 'キーワード'
};

export default SearchBar;
