import { AxiosResponse } from 'axios';
import { internal } from '~/lib/axios';
import { ChatRoomListV1Response } from '~/types/api/chat_apiv1_response_pb';

export async function getRoomInfoByUserId(userId: string) {
  const res = await internal
    .get(`/v1/users/${userId}/chat`)
    .then((res: AxiosResponse<ChatRoomListV1Response.AsObject>) => {
      return res.data;
    });
  return res;
}
