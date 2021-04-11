import 'jest';
import React from 'react';

import { shallow, configure } from 'enzyme';
import Adapter from 'enzyme-adapter-react-16';
import SearchBar from '~/components/molecules/SearchBar';
import { SearchBar as ElementsSearchBar } from 'react-native-elements';
import { ViewStyle } from 'react-native';

configure({ adapter: new Adapter() });

describe('<SearchBar />', () => {
  it('has default props', () => {
    const placeholder = 'キーワード';
    const inputContainerStyle: ViewStyle = {
      height: 14,
    };
    const wrapper = shallow(<SearchBar />);

    const bar = wrapper.find(ElementsSearchBar).get(0);

    expect(bar.props.placeholder).toEqual(placeholder);
    expect(bar.props.inputContainerStyle).toEqual(inputContainerStyle);
  });
});
