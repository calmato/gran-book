import 'jest';

import { shallow, configure } from 'enzyme';
import Adapter from 'enzyme-adapter-react-16';
import React from 'react';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';

configure({ adapter: new Adapter() });

describe('<HeaderWithBackButton />', () => {
  it('has default props', () => {
    // TODO
    const title = 'test';
    const wrapper = shallow(<HeaderWithBackButton title={title} onPress={jest.fn()} />);
  });
});
