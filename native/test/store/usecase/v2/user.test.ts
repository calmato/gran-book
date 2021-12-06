import axios from 'axios';
import MockAdapter from 'axios-mock-adapter';
import { getOwnProfile } from '@/store/usecases/v2/user';
import internalInstance from '~/lib/axios/internal';
import { UserProfileV1Response } from '~/types/api/user_apiv1_response_pb';

/**
 * axiosのmock
 */
const mockAxios = new MockAdapter(internalInstance);
const API_VERSION = 'v1';

describe('getOwnProfile', () => {
  test('return positive response when correct userId and token', async () => {
    const userId = '';
    const token = '';

    const mockResponse: UserProfileV1Response.AsObject = {
      id: '',
      username: '',
      thumbnailUrl: '',
      selfIntroduction: '',
      isFollow: false,
      isFollower: false,
      followCount: 0,
      followerCount: 0,
      rating: 0,
      reviewCount: 0,
      productsList: [],
    };

    mockAxios.onGet(`/${API_VERSION}/users/${userId}/profile`).reply(200, mockResponse);

    const actual = await getOwnProfile(userId, token);

    expect(actual.id).toBe(userId);
  });

  test('return Promise.reject when api server is down', async () => {
    const userId = '';
    const token = '';
    mockAxios.onGet(`/${API_VERSION}/users/${userId}/profile`).networkError();
    try {
      await getOwnProfile(userId, token);
    } catch (e) {
      expect(axios.isAxiosError(e)).toBeTruthy();
      if (axios.isAxiosError(e)) {
        expect(e.message).toMatch('Network Error');
      }
    }
  });
});
