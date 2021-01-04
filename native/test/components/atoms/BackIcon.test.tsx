import 'jest';

import React from 'react';
import * as ShallowRenderer from 'react-test-renderer/shallow';

import BackIcon from '@/components/atoms/BackIcon';
import { colors } from 'react-native-elements';

describe('<BackIcon />', () => {

  let render: ShallowRenderer.ShallowRenderer;

  beforeEach(() => {
    render = ShallowRenderer.createRenderer();
  });
  
  it('has default props', () => {
    render.render(<BackIcon />);
    const got = render.getRenderOutput();

    expect(got).toBeTruthy();
    expect(got.props.name).toEqual('ios-arrow-back');
    expect(got.props.size).toEqual(12);
    expect(got.props.color).toEqual(undefined);
  });

  it('has setting props', () => {
    const size = 24;
    const color = colors.primary;
    render.render(<BackIcon size={size} color={color} />);
    const got = render.getRenderOutput();

    expect(got).toBeTruthy();
    expect(got.props.size).toEqual(size);
    expect(got.props.color).toEqual(color);
  });
});
