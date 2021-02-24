import AsyncStorage from '@react-native-community/async-storage';
import { Auth } from '~/store/models';
import { AUTH_KEY } from './keys';

export async function save(values: Auth.Model): Promise<void> {
  await AsyncStorage.setItem(AUTH_KEY, JSON.stringify(values));
}

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export async function retrieve(): Promise<any> {
  const serialized = await AsyncStorage.getItem(AUTH_KEY);
  if (!serialized) {
    return null;
  }

  return JSON.parse(serialized);
}

export async function clear(): Promise<void> {
  await AsyncStorage.removeItem(AUTH_KEY);
}
