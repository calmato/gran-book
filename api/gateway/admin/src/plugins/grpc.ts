import { credentials } from '@grpc/grpc-js'
import { BookServiceClient, IBookServiceClient } from '~/proto/book_apiv1_grpc_pb'
import {
  AdminServiceClient,
  AuthServiceClient,
  IAdminServiceClient,
  IAuthServiceClient,
} from '~/proto/user_apiv1_grpc_pb'

const userAPIURL: string = process.env.USER_API_URL || 'user_api:8080'
const bookAPIURL: string = process.env.BOOK_API_URL || 'book_api:8080'

const authClient: IAuthServiceClient = new AuthServiceClient(userAPIURL, credentials.createInsecure())
const adminClient: IAdminServiceClient = new AdminServiceClient(userAPIURL, credentials.createInsecure())
const bookClient: IBookServiceClient = new BookServiceClient(bookAPIURL, credentials.createInsecure())

export { authClient, adminClient, bookClient }
