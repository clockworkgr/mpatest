/* eslint-disable */
import * as Long from 'long'
import { util, configure, Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'chain.mpatest.mpa'

export interface Position {
  creator: string
  id: number
  collateral: number
  supply: number
}

const basePosition: object = { creator: '', id: 0, collateral: 0, supply: 0 }

export const Position = {
  encode(message: Position, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id)
    }
    if (message.collateral !== 0) {
      writer.uint32(24).int32(message.collateral)
    }
    if (message.supply !== 0) {
      writer.uint32(32).int32(message.supply)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): Position {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...basePosition } as Position
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.id = longToNumber(reader.uint64() as Long)
          break
        case 3:
          message.collateral = reader.int32()
          break
        case 4:
          message.supply = reader.int32()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): Position {
    const message = { ...basePosition } as Position
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id)
    } else {
      message.id = 0
    }
    if (object.collateral !== undefined && object.collateral !== null) {
      message.collateral = Number(object.collateral)
    } else {
      message.collateral = 0
    }
    if (object.supply !== undefined && object.supply !== null) {
      message.supply = Number(object.supply)
    } else {
      message.supply = 0
    }
    return message
  },

  toJSON(message: Position): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.id !== undefined && (obj.id = message.id)
    message.collateral !== undefined && (obj.collateral = message.collateral)
    message.supply !== undefined && (obj.supply = message.supply)
    return obj
  },

  fromPartial(object: DeepPartial<Position>): Position {
    const message = { ...basePosition } as Position
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id
    } else {
      message.id = 0
    }
    if (object.collateral !== undefined && object.collateral !== null) {
      message.collateral = object.collateral
    } else {
      message.collateral = 0
    }
    if (object.supply !== undefined && object.supply !== null) {
      message.supply = object.supply
    } else {
      message.supply = 0
    }
    return message
  }
}

declare var self: any | undefined
declare var window: any | undefined
var globalThis: any = (() => {
  if (typeof globalThis !== 'undefined') return globalThis
  if (typeof self !== 'undefined') return self
  if (typeof window !== 'undefined') return window
  if (typeof global !== 'undefined') return global
  throw 'Unable to locate global object'
})()

type Builtin = Date | Function | Uint8Array | string | number | undefined
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error('Value is larger than Number.MAX_SAFE_INTEGER')
  }
  return long.toNumber()
}

if (util.Long !== Long) {
  util.Long = Long as any
  configure()
}
