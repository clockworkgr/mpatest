/* eslint-disable */
import { Position } from '../mpa/position';
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'chain.mpatest.mpa';
const baseGenesisState = {};
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        for (const v of message.positionList) {
            Position.encode(v, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.positionList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.positionList.push(Position.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseGenesisState };
        message.positionList = [];
        if (object.positionList !== undefined && object.positionList !== null) {
            for (const e of object.positionList) {
                message.positionList.push(Position.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.positionList) {
            obj.positionList = message.positionList.map((e) => e ? Position.toJSON(e) : undefined);
        }
        else {
            obj.positionList = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.positionList = [];
        if (object.positionList !== undefined && object.positionList !== null) {
            for (const e of object.positionList) {
                message.positionList.push(Position.fromPartial(e));
            }
        }
        return message;
    }
};
