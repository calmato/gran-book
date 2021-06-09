import 'jest';

import { shallow, configure } from 'enzyme';
import Adapter from 'enzyme-adapter-react-16';
import React from 'react';
import HeaderWithCloseButton from '~/components/organisms/HeaderWithCloseButton';

configure({ adapter: new Adapter() });

describe('<HeaderWithCloseButton />', () => {
  it('has default props', () => {
    // TODO
    const title = 'test';
    const wrapper = shallow(<HeaderWithCloseButton title={title} onPress={jest.fn()} />);
  });
});
