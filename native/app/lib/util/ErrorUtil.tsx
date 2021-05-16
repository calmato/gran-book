import { ERROR_MESSAGE } from '../resources/messages_ja';

export const generateErrorMessage = function generateErrorMessage(code: string): string {
  switch (code) {
    case 'There is no user record corresponding to this identifier. The user may have been deleted.':
      return ERROR_MESSAGE.USER_NOT_FOUND;
    case 'Email address is unapproved':
      return ERROR_MESSAGE.USER_UNAUTHORIZED;
    case 'Request failed with status code 400':
      return ERROR_MESSAGE.BAD_REQUEST;
    case 'Request failed with status code 401':
      return ERROR_MESSAGE.UNAUTHORIZED;
    case 'Request failed with status code 403':
    case 'Request failed with status code 404':
    case 'Request failed with status code 409':
      return ERROR_MESSAGE.PROCESS_FAILED;
    case 'Request failed with status code 500':
    case 'Request failed with status code 501':
    case 'Request failed with status code 503':
      return ERROR_MESSAGE.SERVER_ERROR;
    case 'Request failed with status code 504':
      return ERROR_MESSAGE.TIMEOUT;
    case 'Network Error':
      return ERROR_MESSAGE.NETWORK_ERROR;
    default:
      return ERROR_MESSAGE.UNEXPEXTED_ERROR;
  }
};
