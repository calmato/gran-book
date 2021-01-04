import React, { ReactElement } from 'react';
import { Header } from 'react-native-elements';
import HeaderText from '~/components/atoms/HeaderText';
import BackButton from '~/components/molecules/BackButton';


interface Props {
  title: string,
  onPress: () => void | undefined
}

const HeaderWithBackButton = function HeaderWithBackButton(props: Props): ReactElement {
  return (
    <Header
      leftComponent={
        <BackButton
          size={24}
          color="white"
          onPress={() => props.onPress()}
        />
      }
      centerComponent={<HeaderText title={props.title} />}
    />
  );
};

export default HeaderWithBackButton;
