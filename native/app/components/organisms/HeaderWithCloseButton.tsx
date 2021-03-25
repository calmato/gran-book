import React, { ReactElement } from 'react';
import { Header } from 'react-native-elements';
import HeaderText from '~/components/atoms/HeaderText';
import CloseButton from '~/components/molecules/CloseButton';

interface Props {
  title: string;
  onPress: () => void | undefined;
}

const HeaderWithCloseButton = function HeaderWithCloseButton(props: Props): ReactElement {
  return (
    <Header
      leftComponent={<CloseButton size={24} color="white" onPress={() => props.onPress()} />}
      centerComponent={<HeaderText title={props.title} />}
    />
  );
};

export default HeaderWithCloseButton;
