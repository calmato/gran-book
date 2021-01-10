import { StackNavigationProp } from '@react-navigation/stack';
import { StyleSheet, View } from 'react-native';
import { ReactElement } from 'react';
import { RootStackParamList } from '~/types/navigation';
import { ListItem, Icon } from 'react-native-elements'
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
  },
  checkBox: {
    marginBottom: 24,
  },
});

type MyPageProp = StackNavigationProp<RootStackParamList, 'SignInSelect'>;

interface Props {
  navigaton: MyPageProp
}

const MyPage = function MyPage(props: Props): ReactElement {
  const navigaton = props.navigaton;

  return (
    <View>
      <HeaderWithBackButton
        title="マイページ"
        onPress={() => navigaton.goBack()}
      />
    </View>
  );
};

export default MyPage;
