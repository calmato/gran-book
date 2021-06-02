import { ERROR_MESSAGE } from '../resources/messages_ja';

class CustomError extends Error {
  public errorMessageForUsers: string = '';
  constructor(...params) {
    super(...params);
    if(!this.message.replace(/[^0-9]/g, '')){
      this.generateErrorMessageForUsersWithErrorCode
    }
  }

  private generateErrorMessageForUsersWithErrorCode(){
    switch(this.message.replace(/[^0-9]/g, '')){
      case '400':
        this.errorMessageForUsers = ERROR_MESSAGE.BAD_REQUEST;
        break;
      case '401':
        this.errorMessageForUsers = ERROR_MESSAGE.UNAUTHORIZED;
        break;
      case '403':
      case '404':
      case '409':
        this.errorMessageForUsers = ERROR_MESSAGE.PROCESS_FAILED;
        break;
      case '500':
      case '501':
      case '503':
        this.errorMessageForUsers = ERROR_MESSAGE.SERVER_ERROR;
        break;
      case '504':
        this.errorMessageForUsers = ERROR_MESSAGE.TIMEOUT;
        break;
      default:
        this.errorMessageForUsers = ERROR_MESSAGE.UNEXPEXTED_ERROR;
        break;
    }
  }

  public getErrorMessageForUsers(): string{
    return this.errorMessageForUsers;
  }

  public setErrorMessageForUsers(message: string) {
    this.errorMessageForUsers = message;
    return;
  }
}

export default CustomError;
