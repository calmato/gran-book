import express, { NextFunction, Request, Response } from 'express'
import { createChatRoom, getUser, listChatRoom, listUserWithUserIds } from '~/api'
import { GrpcError } from '~/types/exception'
import { ICreateChatRoomInput, IGetUserInput, IListChatRoomInput, IListUserByUserIdsInput } from '~/types/input'
import { IChatRoomListOutput, IChatRoomListOutputRoom, IChatRoomOutput, IUserHashOutput } from '~/types/output'
import { ICreateChatRoomRequest } from '~/types/request'
import {
  IChatRoomListResponse,
  IChatRoomListResponseMessage,
  IChatRoomListResponseRoom,
  IChatRoomListResponseUser,
  IChatRoomResponse,
  IChatRoomResponseUser,
} from '~/types/response'

const router = express.Router()

router.get(
  '/v1/users/:userId/chat',
  async (req: Request, res: Response<IChatRoomListResponse>, next: NextFunction): Promise<void> => {
    const { userId } = req.params
    const { limit, offset } = req.query as { [key: string]: string }

    const roomsInput: IListChatRoomInput = {
      userId: userId,
      limit: Number(limit) || 100,
      offset: Number(offset) || 0,
    }

    await listChatRoom(req, roomsInput)
      .then((roomsOutput: IChatRoomListOutput) => {
        let userIds: string[] = []
        roomsOutput.rooms.forEach((cr: IChatRoomListOutputRoom) => {
          userIds = userIds.concat(cr.userIds)
        })

        const userListInput: IListUserByUserIdsInput = {
          ids: Array.from(new Set(userIds)),
        }

        return listUserWithUserIds(req, userListInput)
          .then((usersOutput: IUserHashOutput) => {
            console.log('debug', usersOutput)
            return setChatRoomListResponse(roomsOutput, usersOutput)
          })
          .catch(() => {
            return setChatRoomListResponse(roomsOutput, {})
          })
      })
      .then((response: IChatRoomListResponse) => {
        res.status(200).json(response)
      })
      .catch((err: GrpcError) => next(err))
  }
)

router.post(
  '/v1/users/:userId/chat',
  async (req: Request, res: Response<IChatRoomResponse>, next: NextFunction): Promise<void> => {
    const { userId } = req.params
    const { users } = req.body as ICreateChatRoomRequest

    const userInput: IGetUserInput = {
      id: userId,
    }

    await getUser(req, userInput)
      .then(() => {
        const userListInput: IListUserByUserIdsInput = {
          ids: users,
        }

        return listUserWithUserIds(req, userListInput)
      })
      .then(async (usersOutput: IUserHashOutput) => {
        const roomInput: ICreateChatRoomInput = {
          userIds: users,
        }

        return createChatRoom(req, roomInput)
          .then((roomOutput: IChatRoomOutput) => {
            return setChatRoomResponse(roomOutput, usersOutput)
          })
          .catch((err: GrpcError) => {
            throw err
          })
      })
      .then((response: IChatRoomResponse) => {
        res.status(200).json(response)
      })
      .catch((err: GrpcError) => next(err))
  }
)

function setChatRoomResponse(roomOutput: IChatRoomOutput, usersOutput?: IUserHashOutput): IChatRoomResponse {
  const users = roomOutput.userIds.map(
    (userId: string): IChatRoomResponseUser => {
      if (!usersOutput || !usersOutput[userId]) {
        const user: IChatRoomListResponseUser = {
          id: userId,
          username: '',
          thumbnailUrl: '',
        }

        return user
      }

      const user: IChatRoomResponseUser = {
        id: userId,
        username: usersOutput[userId].username,
        thumbnailUrl: usersOutput[userId].thumbnailUrl,
      }

      return user
    }
  )

  const response: IChatRoomResponse = {
    id: roomOutput.id,
    users,
    createdAt: roomOutput.createdAt,
    updatedAt: roomOutput.updatedAt,
  }

  return response
}

function setChatRoomListResponse(
  roomsOutput: IChatRoomListOutput,
  usersOutput?: IUserHashOutput
): IChatRoomListResponse {
  const rooms = roomsOutput.rooms.map(
    (cr: IChatRoomListOutputRoom): IChatRoomListResponseRoom => {
      const users = cr.userIds.map(
        (userId: string): IChatRoomListResponseUser => {
          if (!usersOutput || !usersOutput[userId]) {
            const user: IChatRoomListResponseUser = {
              id: userId,
              username: '',
              thumbnailUrl: '',
            }

            return user
          }

          const user: IChatRoomListResponseUser = {
            id: userId,
            username: usersOutput[userId].username,
            thumbnailUrl: usersOutput[userId].thumbnailUrl,
          }

          return user
        }
      )

      const room: IChatRoomListResponseRoom = {
        id: cr.id,
        users,
        createdAt: cr.createdAt,
        updatedAt: cr.updatedAt,
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
    }
  )

  const response: IChatRoomListResponse = {
    rooms,
  }

  return response
}

export default router
