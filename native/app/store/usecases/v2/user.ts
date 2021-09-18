import { AxiosResponse } from 'axios';
import { internal } from '~/lib/axios';
import { getAuthHeader } from '~/lib/axios/util';
import { User } from '~/store/models';
import { UserProfileV1Response } from '~/types/api/user_apiv1_response_pb';

/**
 *
 * @param userId
 * @param token
 * @returns
 */
export async function getOwnProfile(userId: string, token: string) {
  try {
    const { data }: AxiosResponse<UserProfileV1Response.AsObject> = await internal.get(
      `/v1/users/${userId}/profile`,
      getAuthHeader(token),
    );

    const products: User.product[] = data.productsList.map((product) => {
      return {
        id: product.id,
        name: product.name,
        thumbnailUrl: product.thumbnailUrl,
        authors: product.authorsList,
      };
    });

    const values: User.UserValues = {
      id: data.id,
      username: data.username,
      thumbnailUrl: data.thumbnailUrl,
      selfIntroduction: data.selfIntroduction,
      isFollow: data.isFollow,
      isFollower: data.isFollower,
      followCount: data.followCount,
      followerCount: data.followerCount,
      reviewCount: data.reviewCount,
      rating: data.rating,
      products,
    };

    return values;
  } catch (e) {
    return Promise.reject(e);
  }
}
