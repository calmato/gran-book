import express, { NextFunction, Request, Response } from 'express'
import { IChatRoomListOutput, IChatRoomListOutputRoom, IChatRoomOutput, IUserHashOutput } from '~/types/output'
import { IChatRoomListResponse, IChatRoomListResponseMessage, IChatRoomListResponseRoom, IChatRoomListResponseUser, IChatRoomResponse } from '~/types/response'

const router = express.Router()

router.get(
  '/v1/users/:userId/chat',
  async (req: Request, res: Response<IChatRoomListResponse>, next: NextFunction): Promise<void> => {}
)

router.post(
  '/v1/users/:userId/chat',
  async (req: Request, res: Response<IChatRoomListResponse>, next: NextFunction): Promise<void> => {}
)

function setChatRoomResponse(output: IChatRoomOutput): IChatRoomResponse {
  const response: IChatRoomResponse = {
    id: output.id,
    users: [],
    createdAt: output.createdAt,
    updatedAt: output.updatedAt,
  }

  return response
}

function setChatRoomListResponse(roomsOutput: IChatRoomListOutput, usersOutput: IUserHashOutput): IChatRoomListResponse {
  const rooms = roomsOutput.rooms.map((cr: IChatRoomListOutputRoom): IChatRoomListResponseRoom => {
    const users = cr.userIds.map((userId: string): IChatRoomListResponseUser => {
      const user: IChatRoomListResponseUser = {
        id: userId,
        username: usersOutput[userId]?.username || '',
        thumbnailUrl: usersOutput[userId]?.thumbnailUrl || '',
      }

      return user
    })

    const room: IChatRoomListResponseRoom = {
      id: cr.id,
      users,
    }

    if (cr.latestMessage) {
      const m: IChatRoomListResponseMessage = {
        userId: cr.latestMessage.userId,
        text: cr.latestMessage.text,
        image: cr.latestMessage.image,
        createdAt: cr.latestMessage.createdAt,
      }

      room.latestMassage = m
    }

    return room
  })

  const response: IChatRoomListResponse = {
    rooms,
  }

  return response
}

export default router
