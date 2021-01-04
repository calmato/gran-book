import { StackNavigationProp } from '@react-navigation/stack';
import React, { ReactElement } from 'react';
import { StyleSheet, View } from 'react-native';
import HeaderWithBackButton from '~/components/organisms/HeaderWithBackButton';
import { AuthStackParamList } from '~/types/navigation';

const styles = StyleSheet.create({
    container: {
        flex: 1,
        alignItems: 'center',
    }
});

type SignInProp = StackNavigationProp<AuthStackParamList, 'SignIn'>

interface Props {
    navigation: SignInProp,
}

const SignIn = function SignIn(props: Props): ReactElement{
    const navigation = props.navigation;
    return (
        <View style={styles.container}>
            <HeaderWithBackButton
                title='サインイン'
                onPress={() => navigation.goBack()}
            />
        </View>
    );
};

export default SignIn;