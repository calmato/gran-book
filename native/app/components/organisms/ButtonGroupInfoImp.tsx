import React, { ReactElement } from 'react';
import { ButtonGroup } from 'react-native-elements';

interface Props {
  handleOnPressed: (selectedIndex) => void,
  selectedIndex: number,
}

const ButtonGroupInfoImp = function ButtonGroupInfoImp(props:Props): ReactElement {
  const buttons = ['情報','感想']
  return (
    <ButtonGroup
      buttons={buttons}
      onPress={(selectedIndex)=>props.handleOnPressed(selectedIndex)}
      selectedIndex={props.selectedIndex}
    />
  );
};

export default ButtonGroupInfoImp;
