export interface IChatRoomOutput {
  id: string
  userIds: Array<string>
  createdAt: string
  updatedAt: string
}

export interface IChatRoomListOutput {
  rooms: Array<IChatRoomListOutputRoom>
}

export interface IChatRoomListOutputRoom {
  id: string
  userIds: Array<string>
  createdAt: string
  updatedAt: string
  latestMessage?: IChatRoomListOutputMessage
}

export interface IChatRoomListOutputMessage {
  userId: string
  text: string
  image: string
  createdAt: string
}

export interface IChatMessageOutput {
  id: string
  userId: string
  text: string
  image: string
  createdAt: string
}
