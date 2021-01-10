// email形式ならtrueを返す
export const emailValidation = (value: string | undefined): boolean => {
  if (value) {
    return (/^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$/).test(value);
  } else {
    return true;
  }
};

// ６文字以上32文字以下ならtrueを返す
export const passwordValidation = (value: string): boolean => {
  return value.length >= 6 && value.length <= 32;
};
