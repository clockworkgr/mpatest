// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgDeletePosition } from "./types/mpa/tx";
import { MsgUpdatePosition } from "./types/mpa/tx";
import { MsgCreatePosition } from "./types/mpa/tx";
const types = [
    ["/chain.mpatest.mpa.MsgDeletePosition", MsgDeletePosition],
    ["/chain.mpatest.mpa.MsgUpdatePosition", MsgUpdatePosition],
    ["/chain.mpatest.mpa.MsgCreatePosition", MsgCreatePosition],
];
export const MissingWalletError = new Error("wallet is required");
const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw MissingWalletError;
    const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee, memo } = { fee: defaultFee, memo: "" }) => client.signAndBroadcast(address, msgs, fee, memo),
        msgDeletePosition: (data) => ({ typeUrl: "/chain.mpatest.mpa.MsgDeletePosition", value: data }),
        msgUpdatePosition: (data) => ({ typeUrl: "/chain.mpatest.mpa.MsgUpdatePosition", value: data }),
        msgCreatePosition: (data) => ({ typeUrl: "/chain.mpatest.mpa.MsgCreatePosition", value: data }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
