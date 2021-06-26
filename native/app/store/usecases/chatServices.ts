import { AxiosResponse } from 'axios';
import { internal } from '~/lib/axios';
import { RoomInfoResponse } from '~/types/response/chat';
import { IErrorResponse } from  '~/types/response/external/rakuten-books';


async function getRoomInfoByUserId(
  userId: string,
): Promise<AxiosResponse<any> | AxiosResponse<IErrorResponse>> {
  return internal
    .get('/v1/users/{userId}/chat')
    .then((res:AxiosResponse<RoomInfoResponse>) => {
      return res;
    })
    .catch((err: AxiosResponse<IErrorResponse>) => {
      return Promise.reject(err);
    });
}
