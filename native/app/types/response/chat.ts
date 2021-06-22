export interface roomInfo {
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
