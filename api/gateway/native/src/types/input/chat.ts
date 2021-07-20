export interface IListChatRoomInput {
  userId: string
  limit: number
  offset: number
}

export interface ICreateChatRoomInput {
  userIds: Array<string>
}
