export interface IListChatRoomInput {
  userId: string
  limit: number
  offset: number
}

export interface ICreateChatRoomInput {
  userIds: Array<string>
}

export interface ICreateChatMessageInput {
  roomId: string
  text: string
}

export interface IUploadChatImageInput {
  roomId: string
  path: string
}
