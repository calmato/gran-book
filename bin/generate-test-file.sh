#!/bin/bash

echo "generate components unit test files …"

test_init_code(){
  code=$(cat <<EOF
import 'jest';
import React from 'react';

import { shallow, configure } from 'enzyme';
import Adapter from 'enzyme-adapter-react-16.1';

configure({ adapter: new Adapter() });

describe('<$1 />', () => {
  it('has default props', () => {
    const wrapper = shallow();
  });
});
EOF
)
  echo "$code"
}

dir=($(ls $1))

for d in ${dir[@]}; do
  mkdir -p $2/${d}
  files=($(ls $1/${d}/))
  for file in ${files[@]}; do
    if [ "$file" != ".keep" ] && [ ! -f "$2/${d}/${file%%.*}.test.tsx" ]; then
      test_init_code ${file%%.*} >> "$2/${d}/${file%%.*}.test.tsx"
    fi
  done
done
