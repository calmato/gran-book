import 'jest';
import React from 'react';

import { shallow, configure } from 'enzyme';
import Adapter from 'enzyme-adapter-react-16.1';
import GoogleButton from '~/components/molecules/GoogleButton';

configure({ adapter: new Adapter() });

describe('<GoogleButton />', () => {
  it('has default props', () => {
    const wrapper = shallow(<GoogleButton />);
  });
});
