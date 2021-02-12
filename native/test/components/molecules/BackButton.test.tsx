import 'jest';
import React from 'react';

import { shallow, configure } from 'enzyme';
import Adapter from 'enzyme-adapter-react-16.1';
import BackButton from '~/components/molecules/BackButton';
import BackIcon from '~/components/atoms/BackIcon';

configure({ adapter: new Adapter() });

describe('<BackButton />', () => {
  it('has default props', () => {
    const wrapper = shallow(<BackButton />);

    const backIcon = wrapper.find(BackIcon).get(0);

    expect(backIcon.size).toEqual(undefined);
    expect(backIcon.color).toEqual(undefined);
  });
});
