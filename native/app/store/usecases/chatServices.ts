import { AxiosResponse } from 'axios';
import { internal } from '~/lib/axios';
import { RoomInfoResponse } from '~/types/response/chat';

export async function getRoomInfoByUserId(userId: string) {
  const res = await internal
    .get(`/v1/users/${userId}/chat`)
    .then((res: AxiosResponse<RoomInfoResponse>) => {
      return res.data;
    });
  return res;
}
