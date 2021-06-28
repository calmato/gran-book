import { Request } from 'express'
import { chatClient } from '~/plugins/grpc'
import { getGrpcError } from '~/lib/grpc-exception'
import { getGrpcMetadata } from '~/lib/grpc-metadata'
import { ICreateChatRoomInput, IListChatRoomInput } from '~/types/input'
import {
  IChatRoomListOutput,
  IChatRoomListOutputMessage,
  IChatRoomListOutputRoom,
  IChatRoomOutput,
} from '~/types/output'
import {
  ChatRoomListResponse,
  ChatRoomResponse,
  CreateChatRoomRequest,
  ListChatRoomRequest,
} from '~/proto/user_apiv1_pb'

export function listChatRoom(req: Request<any>, input: IListChatRoomInput): Promise<IChatRoomListOutput> {
  const request = new ListChatRoomRequest()
  const metadata = getGrpcMetadata(req)

  request.setUserId(input.userId)
  request.setLimit(input.limit)
  request.setOffset(input.offset)

  return new Promise((resolve: (res: IChatRoomListOutput) => void, reject: (reason: Error) => void) => {
    chatClient.listRoom(request, metadata, (err: any, res: ChatRoomListResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      resolve(setChatRoomListOutput(res))
    })
  })
}

export function createChatRoom(req: Request<any>, input: ICreateChatRoomInput): Promise<IChatRoomOutput> {
  const request = new CreateChatRoomRequest()
  const metadata = getGrpcMetadata(req)

  request.setUserIdsList(input.userIds)

  return new Promise((resolve: (res: IChatRoomOutput) => void, reject: (reason: Error) => void) => {
    chatClient.createRoom(request, metadata, (err: any, res: ChatRoomResponse) => {
      if (err) {
        reject(getGrpcError(err))
        return
      }

      resolve(setChatRoomOutput(res))
    })
  })
}

function setChatRoomOutput(res: ChatRoomResponse): IChatRoomOutput {
  const output: IChatRoomOutput = {
    id: res.getId(),
    userIds: res.getUserIdsList(),
    createdAt: res.getCreatedAt(),
    updatedAt: res.getUpdatedAt(),
  }

  return output
}

function setChatRoomListOutput(res: ChatRoomListResponse): IChatRoomListOutput {
  const rooms = res.getRoomsList().map(
    (cr: ChatRoomListResponse.Room): IChatRoomListOutputRoom => {
      const room: IChatRoomListOutputRoom = {
        id: cr.getId(),
        userIds: cr.getUserIdsList(),
        createdAt: '', // TODO: 追加
        updatedAt: '', // TODO: 追加
      }

      const message: ChatRoomListResponse.Message | undefined = cr.getLatestmessage()
      if (message) {
        const m: IChatRoomListOutputMessage = {
          userId: message.getUserId(),
          text: message.getText(),
          image: message.getImage(),
          createdAt: message.getCreatedAt(),
        }

        room.latestMessage = m
      }

      return room
    }
  )

  const output: IChatRoomListOutput = {
    rooms,
  }

  return output
}
