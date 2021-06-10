import React, { ReactElement } from 'react';
import { Platform, StyleSheet } from 'react-native';
import { SearchBar as ElementsSearchBar } from 'react-native-elements';
import SearchIcon from '~/components/atoms/SearchIcon';
import { COLOR } from '~~/constants/theme';

const styles = StyleSheet.create({
  inputContainerStyle: {
    height: 14,
  },
});

interface Props {
  keyword: string;
  placeholder?: string;
  onChangeText: (t: string) => void;
  onSubmitEditing: () => void;
  onCancel?: () => void;
}

const SearchBar = function SearchBar(props: Props): ReactElement {
  const platform = Platform.OS === 'ios' ? 'ios' : 'android';

  return (
    <ElementsSearchBar
      placeholder={props.placeholder}
      inputContainerStyle={styles.inputContainerStyle}
      cancelButtonProps={{
        color: COLOR.TEXT_DEFAULT,
      }}
      platform={platform}
      searchIcon={<SearchIcon />}
      clearIcon={false}
      value={props.keyword}
      onChangeText={props.onChangeText}
      onCancel={props.onCancel}
      onSubmitEditing={props.onSubmitEditing}
      multiline={false}
      returnKeyType={'search'}
    />
  );
};

SearchBar.defaultProps = {
  placeholder: 'キーワード',
};

export default SearchBar;
