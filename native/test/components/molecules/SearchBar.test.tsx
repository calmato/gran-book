import 'jest';

import { shallow, configure } from 'enzyme';
import Adapter from 'enzyme-adapter-react-16';
import React from 'react';
import { ViewStyle } from 'react-native';
import { SearchBar as ElementsSearchBar } from 'react-native-elements';
import SearchBar from '~/components/molecules/SearchBar';

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
