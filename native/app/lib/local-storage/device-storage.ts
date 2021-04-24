import AsyncStorage from '@react-native-community/async-storage';
import { DEVICE_KEY } from './keys';

export async function save(value: string): Promise<void> {
  await AsyncStorage.setItem(DEVICE_KEY, JSON.stringify(value));
}

export async function retrieve(): Promise<string> {
  const serialized = await AsyncStorage.getItem(DEVICE_KEY);
  if (!serialized) {
    return '';
  }

  return serialized;
}

export async function clear(): Promise<void> {
  await AsyncStorage.removeItem(DEVICE_KEY);
}
