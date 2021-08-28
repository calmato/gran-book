import { AxiosResponse } from 'axios';
import { Dispatch } from 'redux';
import { internal } from '~/lib/axios';
import { User } from '~/store/models';
import { setUser } from '~/store/modules/user';
import { UserProfileV1Response } from '~/types/api/user_apiv1_response_pb';

export function getOwnProfileAsync(id: string) {
  return async (dispatch: Dispatch): Promise<void> => {
    return await internal
      .get(`/v1/users/${id}/profile`)
      .then(async (res: AxiosResponse<UserProfileV1Response.AsObject>) => {
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
          productsList,
        } = res.data;

        const products: User.product[] = productsList.map((val: UserProfileV1Response.Product.AsObject) => {
          const product: User.product = {
            id: val.id,
            name: val.name,
            thumbnailUrl: val.thumbnailUrl,
            authors: val.authorsList,
          };

          return product;
        });

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

export function getOtherProfileAsync(id: string) {
  return async (dispatch: Dispatch): Promise<void> => {
    return await internal
      .get(`/v1/users/${id}/profile`)
      .then(async (res: AxiosResponse<UserProfileV1Response.AsObject>) => {
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
          productsList,
        } = res.data;

        const products: User.product[] = productsList.map((val: UserProfileV1Response.Product.AsObject) => {
          const product: User.product = {
            id: val.id,
            name: val.name,
            thumbnailUrl: val.thumbnailUrl,
            authors: val.authorsList,
          };

          return product;
        });

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
        throw err;
      });
  };
}
