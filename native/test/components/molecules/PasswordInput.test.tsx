import 'jest';
import React from 'react';

import { shallow, configure } from 'enzyme';
import Adapter from 'enzyme-adapter-react-16';
import PasswordInput from '~/components/molecules/PasswordInput';
import { colors, Input } from 'react-native-elements';

configure({ adapter: new Adapter() });

describe('<PasswordInput />', () => {
  it('has default props', () => {
    const placeholder = 'test';
    const value = 'test';
    const maxLength = 32;
    const color = colors.grey0;

    const wrapper = shallow(<PasswordInput
      placeholder={placeholder}
      value={value}
      onChangeText={() => console.log('test')}
    />);

    const input = wrapper.find(Input).get(0);
    expect(input.props.placeholder).toEqual(placeholder);
    expect(input.props.value).toEqual(value);
    expect(input.props.maxLength).toEqual(maxLength);
    expect(input.props.errorMessage).toEqual('');

    const leftIcon = input.props.leftIcon;
    expect(leftIcon.props.name).toEqual('lock');
    expect(leftIcon.props.size).toEqual(24);
    expect(leftIcon.props.color).toEqual(color);

    const touchableOpacity = input.props.rightIcon;
    expect(touchableOpacity.props.onPress()).toBeCalled;

    const rightIcon = touchableOpacity.props.children;
    expect(rightIcon.props.name).toEqual('md-eye');
    expect(rightIcon.props.size).toEqual(24);
    expect(rightIcon.props.color).toEqual(color);
  });
});
