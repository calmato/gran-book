import 'jest';
import { Ionicons } from '@expo/vector-icons';
import { shallow, configure } from 'enzyme';
import Adapter from 'enzyme-adapter-react-16';
import React from 'react';

import { Text, View, ViewStyle } from 'react-native';
import CheckBox from '~/components/molecules/CheckBox';
import { COLOR } from '~~/constants/theme';

configure({ adapter: new Adapter() });

describe('<CheckBox />', () => {
  it('has default props', () => {
    const title = 'test';
    const checked = true;
    const viewStyle: ViewStyle = {
      flexDirection: 'row',
      alignItems: 'center',
    };

    const wrapper = shallow(
      <CheckBox title={title} checked={checked} onPress={() => console.log('test')} />,
    );

    const view = wrapper.find(View).get(0);
    const icon = wrapper.find(Ionicons).get(0);
    const text = wrapper.find(Text).get(0);

    expect(view.props.style).toEqual(viewStyle);

    expect(icon.props.name).toEqual('md-checkbox');
    expect(icon.props.size).toEqual(24);
    expect(icon.props.color).toEqual(COLOR.PRIMARY);

    expect(text.props.children).toEqual(title);
  });
});
