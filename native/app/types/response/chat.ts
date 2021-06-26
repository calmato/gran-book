export interface RoomInfoResponse {
  limit: number;
  offset: number;
  total: number;
  rooms: Array<RoomInfo>;
}

export interface RoomInfo {
  rooms: [
    {
      id: string,
      users: [
        {
          id: string,
          username: string,
          thumbnailUrl: string
        }
      ],
      latestMassage: {
        userId: string,
        text: string,
        image: string,
        createdAt: string
      },
      createdAt: string,
      updatedAt: string
    }
  ]
}
