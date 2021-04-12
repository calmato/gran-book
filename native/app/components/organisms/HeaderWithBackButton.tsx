import React, { ReactElement } from 'react';
import { Header } from 'react-native-elements';
import HeaderText from '~/components/atoms/HeaderText';
import BackButton from '~/components/molecules/BackButton';
import { COLOR } from '~~/constants/theme';

interface Props {
  title: string;
  onPress: () => void | undefined;
}

const HeaderWithBackButton = function HeaderWithBackButton(props: Props): ReactElement {
  return (
    <Header
      leftComponent={
        <BackButton size={24} color={COLOR.TEXT_TITLE} onPress={() => props.onPress()} />
      }
      centerComponent={<HeaderText title={props.title} />}
    />
  );
};

export default HeaderWithBackButton;
