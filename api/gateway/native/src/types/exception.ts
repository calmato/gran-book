abstract class CustomError extends Error {
  constructor(err?: string) {
    super(err)
  }
}

export class GrpcError extends CustomError {
  constructor(public status: number, public details?: Array<any>) {
    super(`gRPC Status: ${status}`)
  }
}

export class HttpError extends CustomError {
  constructor(public status: number, public message: string, public details?: Array<any>) {
    super(`HTTP Status: ${status}`)
  }
}
