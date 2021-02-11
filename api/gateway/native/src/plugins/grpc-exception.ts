/* eslint-disable @typescript-eslint/explicit-module-boundary-types */
/* eslint-disable @typescript-eslint/no-unsafe-assignment */
/* eslint-disable @typescript-eslint/no-unsafe-call */
/* eslint-disable @typescript-eslint/no-unsafe-member-access */
/* eslint-disable @typescript-eslint/no-unsafe-return */
/* eslint-disable @typescript-eslint/unbound-method */
import { Metadata } from '@grpc/grpc-js'
import { Status } from "~/proto/status_pb"
import {   RetryInfo,
  DebugInfo,
  QuotaFailure,
  PreconditionFailure,
  BadRequest,
  RequestInfo,
  ResourceInfo,
  Help,
  LocalizedMessage,
} from '~/proto/error_details_pb'

export const googleDeserializeMap = {
  'google.rpc.RetryInfo': RetryInfo.deserializeBinary,
  'google.rpc.DebugInfo': DebugInfo.deserializeBinary,
  'google.rpc.QuotaFailure': QuotaFailure.deserializeBinary,
  'google.rpc.PreconditionFailure': PreconditionFailure.deserializeBinary,
  'google.rpc.BadRequest': BadRequest.deserializeBinary,
  'google.rpc.RequestInfo': RequestInfo.deserializeBinary,
  'google.rpc.ResourceInfo': ResourceInfo.deserializeBinary,
  'google.rpc.Help': Help.deserializeBinary,
  'google.rpc.LocalizedMessage': LocalizedMessage.deserializeBinary,
}

interface ServiceError {
  metadata?: Metadata
}

export function deserializeGrpcStatusDetails<T extends Record<string, (bytes: Uint8Array) => any>>(
  error: ServiceError,
  deserializeMap: T
): {
  status: Status;
  details: Array<ReturnType<T[keyof T]>>;
} | null {
  if (!error.metadata) {
    return null
  }

  const buffer = error.metadata.get("grpc-status-details-bin")[0]
  if (!buffer || typeof buffer === "string") {
    return null
  }

  const status: Status | undefined = Status.deserializeBinary(buffer)

  const details: Array<ReturnType<T[keyof T]>> = status
    .getDetailsList()
    .map((detail): any => {
      const deserialize: any = deserializeMap[detail.getTypeName()]
      if (deserialize) {
        return detail.unpack(deserialize, detail.getTypeName())
      }

      return null
    })
    .filter(notEmpty)

  return {
    status,
    details,
  }
}

export function deserializeGoogleGrpcStatusDetails(error: ServiceError) {
  return deserializeGrpcStatusDetails(error, googleDeserializeMap)
}

const notEmpty = <T>(value: T | null | undefined): value is T => {
  return value !== null && value !== undefined
}

export function getErrorDetails(details: Array<any>): Array<any> {
  // 配列の1つ目だけ取得 <- APIの設計で1つまでしか含まないようにしているため大丈夫なはず
  const detail = details[0]

  switch (true) {
    case detail instanceof BadRequest:
      return detail.toObject()?.fieldViolationsList
    default:
      return []
  }
}

export function getError(err: ServiceError): {status: number; message: string; details: Array<any>} {
  const grpcErrorDetails = deserializeGoogleGrpcStatusDetails(err);
  if (grpcErrorDetails) {
    const { status, details } = grpcErrorDetails
    const errorDetails = getErrorDetails(details)

    return {
      status: status.getCode(),
      message: status.getMessage(),
      details: errorDetails,
    }
  } else {
    return {
      status: 2,
      message: 'Unknown',
      details: [],
    }
  }
}
/* eslint-enable */