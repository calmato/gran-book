import 'jest';

import { shallow, configure } from 'enzyme';
import Adapter from 'enzyme-adapter-react-16';
import React from 'react';
import { colors } from 'react-native-elements';
import CloseIcon from '~/components/atoms/CloseIcon';
import CloseButton from '~/components/molecules/CloseButton';

configure({ adapter: new Adapter() });

describe('<CloseButton />', () => {
  it('has default props', () => {
    const wrapper = shallow(<CloseButton />);

    const closeIcon = wrapper.find(CloseIcon).get(0);

    expect(closeIcon.props.size).toEqual(undefined);
    expect(closeIcon.props.color).toEqual(undefined);
  });

  it('has setting props', () => {
    const size = 24;

    const wrapper = shallow(
      <CloseButton size={size} color={colors.black} onPress={() => console.log('test')} />,
    );

    const closeIcon = wrapper.find(CloseIcon).get(0);

    expect(closeIcon.props.size).toEqual(size);
    expect(closeIcon.props.color).toEqual(colors.black);
  });
});
