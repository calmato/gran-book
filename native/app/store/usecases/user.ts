import { AxiosResponse } from 'axios';
import { Dispatch } from 'redux';
import { internal } from '~/lib/axios';
import { User } from '~/store/models';
import { setUser } from '~/store/modules/user';
import { IOwnOtherProfileResponse } from '~/types/response';

export function getOwnProfileAsync(id: string) {
  return async (dispatch: Dispatch): Promise<void> => {
    console.log('http', 'getOwnProfile');
    return await internal
      .get(`/v1/users/${id}/profile`)
      .then(async (res: AxiosResponse<IOwnOtherProfileResponse>) => {
        const {
          id,
          username,
          thumbnailUrl,
          selfIntroduction,
          isFollow,
          isFollower,
          followCount,
          followerCount,
          reviewCount,
          rating,
          products,
        } = res.data as IOwnOtherProfileResponse;

        const values: User.UserValues = {
          id,
          username,
          thumbnailUrl,
          selfIntroduction,
          isFollow,
          isFollower,
          followCount,
          followerCount,
          reviewCount,
          rating,
          products,
        };

        dispatch(setUser(values));
      })
      .catch((err: Error) => {
        console.log(err);
        throw err;
      });
  };
}
