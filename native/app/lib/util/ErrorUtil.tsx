import { ERROR_MESSAGE } from '../resources/messages_ja';

export const generateErrorMessage = function generateErrorMessage(code:string): string {
  switch(code) {
  case 'There is no user record corresponding to this identifier. The user may have been deleted.':
    return ERROR_MESSAGE.USER_NOT_FOUND;
  case 'Email address is unapproved':
    return ERROR_MESSAGE.USER_UNAUTHORIZED;
  // case 400:
  //   return ERROR_MESSAGE.BAD_REQUEST;
  // case 401:
  //   return ERROR_MESSAGE.UNAUTHORIZED;
  // case 403:
  // case 404:
  // case 409:
  //   return ERROR_MESSAGE.PROCESS_FAILED;
  // case 500:
  // case 501:
  // case 503:
  //   return ERROR_MESSAGE.SERVER_ERROR;
  // case 504:
  //   return ERROR_MESSAGE.TIMEOUT;
  default:
    return ERROR_MESSAGE.UNEXPEXTED_ERROR;
  }
};
