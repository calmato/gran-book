export interface RoomInfoResponse {
  limit: number;
  offset: number;
  total: number;
  info: RoomInfo[];
}

export interface RoomInfo {
  rooms: [
    {
      createdAt?: string;
      id: string;
      latestMessage: {
        createdAt: string;
        id: string;
        text: string;
        image: string;
        userId: string;
      };
      users: [
        {
          id: string;
          username: string;
          thumbnailUrl: string;
        },
      ];
      updatedAt?: string;
    },
  ];
}
