import { internal } from '~/lib/axios';

export async function getRoomInfoByUserId(userId: string) {
  const res = await internal
    .get(`/v1/users/${userId}/chat`);
  return res.data;
}
