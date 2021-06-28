export interface IChatRoomResponse {
  id: string
  users: Array<IChatRoomResponseUser>
  latestMessage?: IChatRoomResponseMessage
  createdAt: string
  updatedAt: string
}

export interface IChatRoomResponseUser {
  id: string
  username: string
  thumbnailUrl: string
}

export interface IChatRoomResponseMessage {
  userId: string
  text: string
  image: string
  createdAt: string
}

export interface IChatRoomListResponse {
  rooms: Array<IChatRoomListResponseRoom>
}

export interface IChatRoomListResponseRoom {
  id: string
  users: Array<IChatRoomListResponseUser>
  latestMassage?: IChatRoomListResponseMessage
  createdAt: string
  updatedAt: string
}

export interface IChatRoomListResponseUser {
  id: string
  username: string
  thumbnailUrl: string
}

export interface IChatRoomListResponseMessage {
  userId: string
  text: string
  image: string
  createdAt: string
}
