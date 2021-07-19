export interface RoomInfoResponse {
  limit: number;
  offset: number;
  total: number;
  info: RoomInfo[];
}

export interface RoomInfo {
  rooms: [
    {
      id: string,
      users: [
        {
          id: string,
          userName: string,
          thumbnailUrl: string
        }
      ],
      latestMessage: {
        userId: string,
        text?: string,
        image?: string,
        createdAt: string
      },
      createdAt?: string,
      updatedAt?: string
    }
  ]
}